package scraper

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"unsafe"

	"github.com/krzysztofreczek/go-structurizr/pkg/model"
)

const (
	rootElementName = "ROOT"
)

type Scraper struct {
	config Configuration
	rules  []Rule

	sb strings.Builder

	components map[string]struct{}
	relations  map[string]map[string]struct{}
}

func NewScraper(config Configuration) *Scraper {
	return &Scraper{
		config:     config,
		rules:      make([]Rule, 0),
		sb:         strings.Builder{},
		components: make(map[string]struct{}),
		relations:  make(map[string]map[string]struct{}),
	}
}

func (s *Scraper) RegisterRule(r Rule) error {
	if r == nil {
		return errors.New("rule must not be nil")
	}
	s.rules = append(s.rules, r)
	return nil
}

func (s *Scraper) Scrap(i interface{}) string {
	v := reflect.ValueOf(i)
	s.scrap(v, rootElementName, "", 0)
	return s.sb.String()
}

func (s *Scraper) scrap(v reflect.Value, name string, parent string, level int) {
	if v.Kind() == reflect.Interface {
		v = v.Elem()
	}

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	scrapable := false
	for _, pkg := range s.config.packages {
		if strings.HasPrefix(v.Type().PkgPath(), pkg) {
			scrapable = true
		}
	}
	if !scrapable {
		return
	}

	str := fmt.Sprintln(strings.Repeat("  ", level) + v.Type().String() + ":" + name)
	s.sb.WriteString(str)

	componentName := strings.Replace(v.Type().String(), ".", "_", -1) + "_" + name
	s.components[componentName] = struct{}{}
	if parent != "" {
		_, ok := s.relations[parent]
		if !ok {
			s.relations[parent] = make(map[string]struct{})
		}
		s.relations[parent][componentName] = struct{}{}
	}

	if v.CanAddr() {
		// supports unexported fields
		if !v.CanInterface() && v.CanAddr() {
			v = reflect.NewAt(v.Type(), unsafe.Pointer(v.UnsafeAddr())).Elem()
		}

		// v.Addr() instead of v supports both value and pointer receiver
		info, ok := v.Addr().Interface().(model.HasInfo)
		if ok {
			str = fmt.Sprintln(strings.Repeat("  ", level) + info.Info().Name + ":" + info.Info().Kind)
			s.sb.WriteString(str)
		}
	}

	for i := 0; i < v.NumField(); i++ {
		s.scrap(v.Field(i), v.Type().Field(i).Name, componentName, level+1)
	}
}

func (s *Scraper) RenderPlantUML() string {
	sb := strings.Builder{}

	for name, _ := range s.components {
		s := fmt.Sprintf("component [%s] as %s\n", name, name)
		sb.WriteString(s)
	}

	for src, targets := range s.relations {
		for trg, _ := range targets {
			s := fmt.Sprintf("%s -> %s\n", src, trg)
			sb.WriteString(s)
		}
	}

	return sb.String()
}

func (s *Scraper) RenderGraphViz() string {
	sb := strings.Builder{}

	for src, targets := range s.relations {
		for trg, _ := range targets {
			s := fmt.Sprintf("%s -> %s;\n", src, trg)
			sb.WriteString(s)
		}
	}

	return sb.String()
}
