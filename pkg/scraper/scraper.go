package scraper

import (
	"fmt"
	"hash/fnv"
	"reflect"
	"strings"
	"unsafe"

	"github.com/krzysztofreczek/go-structurizr/pkg/model"
	"github.com/krzysztofreczek/go-structurizr/pkg/yaml"
	"github.com/pkg/errors"
)

type Scraper struct {
	config    Configuration
	rules     []Rule
	structure model.Structure
}

func NewScraper(config Configuration) *Scraper {
	return &Scraper{
		config:    config,
		rules:     make([]Rule, 0),
		structure: model.NewStructure(),
	}
}

func NewScraperFromConfigFile(fileName string) (*Scraper, error) {
	configuration, err := yaml.LoadFromFile(fileName)
	if err != nil {
		return nil, errors.Wrapf(err,
			"could not load configuration from file `%s`", fileName)
	}

	config := toScraperConfig(configuration)
	rules, err := toScraperRules(configuration)
	if err != nil {
		return nil, errors.Wrapf(err,
			"could not load scraper rules from from configuration file `%s`", fileName)
	}

	return &Scraper{
		config:    config,
		rules:     rules,
		structure: model.NewStructure(),
	}, nil
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
	s.scrap(v, "", 0)
	return s.structure
}

func (s *Scraper) scrap(
	v reflect.Value,
	parentID string,
	level int,
) {
	if v.Kind() == reflect.Interface {
		v = v.Elem()
	}

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() == reflect.Map {
		iterator := v.MapRange()
		for true {
			if !iterator.Next() {
				break
			}
			s.scrap(iterator.Value(), parentID, level)
		}
		return
	}

	if v.Kind() == reflect.Slice || v.Kind() == reflect.Array {
		for i := 0; i < v.Len(); i++ {
			s.scrap(v.Index(i), parentID, level)
		}
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
		s.scrap(v.Field(i), parentID, level+1)
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

		return r.Apply(), true
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
