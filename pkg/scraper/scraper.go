package scraper

import (
	"errors"
	"fmt"
	"hash/fnv"
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

	structure model.Structure
}

func NewScraper(config Configuration) *Scraper {
	return &Scraper{
		config:    config,
		rules:     make([]Rule, 0),
		sb:        strings.Builder{},
		structure: model.NewStructure(),
	}
}

func (s *Scraper) RegisterRule(r Rule) error {
	if r == nil {
		return errors.New("rule must not be nil")
	}
	s.rules = append(s.rules, r)
	return nil
}

func (s *Scraper) Scrap(i interface{}) model.Structure {
	v := reflect.ValueOf(i)
	s.scrap(v, rootElementName, "", 0)
	return s.structure
}

func (s *Scraper) scrap(
	v reflect.Value,
	name string,
	parentID string,
	level int,
) {
	if v.Kind() == reflect.Interface {
		v = v.Elem()
	}

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	v = normalize(v)

	if !s.isScrappable(v) {
		return
	}

	info, ok := s.getInfoFromInterface(v)
	if ok {
		c := s.addComponent(v, info, parentID)
		s.scrapAllFields(v, c.ID, level)
		return
	}

	info, ok = s.getInfoFromRules(v)
	if ok {
		c := s.addComponent(v, info, parentID)
		s.scrapAllFields(v, c.ID, level)
		return
	}

	s.scrapAllFields(v, parentID, level)
	return
}

func (s *Scraper) scrapAllFields(
	v reflect.Value,
	parentID string,
	level int,
) {
	for i := 0; i < v.NumField(); i++ {
		s.scrap(v.Field(i), v.Type().Field(i).Name, parentID, level+1)
	}
}

func (s *Scraper) addComponent(
	v reflect.Value,
	info model.Info,
	parentID string,
) model.Component {
	c := model.Component{
		ID:          componentID(v),
		Kind:        info.Kind,
		Name:        componentName(v),
		Description: info.Description,
		Technology:  info.Technology,
		Tags:        info.Tags,
	}
	s.structure.AddComponent(c, parentID)
	return c
}

func (s *Scraper) isScrappable(v reflect.Value) bool {
	vPkg := valuePackage(v)
	for _, pkg := range s.config.packages {
		if strings.HasPrefix(vPkg, pkg) {
			return true
		}
	}
	return false
}

func (s *Scraper) getInfoFromInterface(v reflect.Value) (model.Info, bool) {
	if !v.CanAddr() {
		return model.Info{}, false
	}

	// v.Addr() instead of v supports both value and pointer receiver
	info, ok := v.Addr().Interface().(model.HasInfo)
	if !ok {
		return model.Info{}, false
	}

	return info.Info(), true
}

func (s *Scraper) getInfoFromRules(v reflect.Value) (model.Info, bool) {
	vPkg := valuePackage(v)
	vType := v.Type().Name()
	for _, r := range s.rules {
		if !r.Applies(vPkg, vType) {
			continue
		}

		info, err := r.Apply()
		if err != nil {
			// TODO: log
			continue
		}

		return info, true
	}

	return model.Info{}, false
}

func normalize(v reflect.Value) reflect.Value {
	if !v.CanAddr() {
		return v
	}

	// supports unexported fields
	if !v.CanInterface() {
		v = reflect.NewAt(v.Type(), unsafe.Pointer(v.UnsafeAddr())).Elem()
	}

	return v
}

func componentID(v reflect.Value) string {
	id := fmt.Sprintf("%s/%s", valuePackage(v), v.Type().Name())
	h := fnv.New32a()
	// TODO: handle
	_, _ = h.Write([]byte(id))
	return fmt.Sprintf("%d", h.Sum32())
}

func componentName(v reflect.Value) string {
	pkg := strings.Split(valuePackage(v), "/")
	name := fmt.Sprintf("%s.%s", pkg[len(pkg)-1], v.Type().Name())
	return name
}

func valuePackage(v reflect.Value) string {
	return v.Type().PkgPath()
}
