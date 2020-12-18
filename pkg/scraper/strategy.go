package scraper

import (
	"fmt"
	"reflect"
	"strings"
	"unsafe"

	"github.com/krzysztofreczek/go-structurizr/pkg/internal"
	"github.com/krzysztofreczek/go-structurizr/pkg/model"
)

func (s *scraper) scrape(
	v reflect.Value,
	parentID string,
	level int,
) {
	// TODO: test invalid type
	if !v.IsValid() {
		return
	}

	// TODO: test circular dependencies
	if v.Type().String() == "model.HasInfo" {

	} else if _, scraped := s.scrapedTypes[v.Type()]; !scraped {
		s.scrapedTypes[v.Type()] = struct{}{}
	} else {
		return
	}

	strategy := s.resolveScrapingStrategy(v)
	strategy(v, parentID, level)
}

type scrapingStrategy func(
	v reflect.Value,
	parentID string,
	level int,
)

func (s *scraper) resolveScrapingStrategy(v reflect.Value) scrapingStrategy {
	switch v.Kind() {
	case reflect.Interface:
		return s.scrapeInterfaceStrategy
	case reflect.Ptr:
		return s.scrapePointerStrategy
	case reflect.Map:
		return s.scrapeMapStrategy
	case reflect.Slice, reflect.Array:
		return s.scrapeIterableStrategy
	// TODO: test functions
	case reflect.Func:
		return s.scrapeFunc
	case reflect.Struct:
		return s.scrapeStruct
	default:
		return s.scrapeNoop
	}
}

func (s *scraper) scrapeInterfaceStrategy(
	v reflect.Value,
	parentID string,
	level int,
) {
	v = v.Elem()
	s.scrape(v, parentID, level)
}

func (s *scraper) scrapePointerStrategy(
	v reflect.Value,
	parentID string,
	level int,
) {
	// TODO: Test nil pointer
	if v.Elem().IsValid() {
		v = v.Elem()
	} else {
		v = reflect.New(v.Type().Elem()).Elem()
	}
	s.scrape(v, parentID, level)
}

func (s *scraper) scrapeMapStrategy(
	v reflect.Value,
	parentID string,
	level int,
) {
	iterator := v.MapRange()
	for true {
		if !iterator.Next() {
			break
		}
		s.scrape(iterator.Value(), parentID, level)
	}
}

func (s *scraper) scrapeIterableStrategy(
	v reflect.Value,
	parentID string,
	level int,
) {
	for i := 0; i < v.Len(); i++ {
		s.scrape(v.Index(i), parentID, level)
	}
}

func (s *scraper) scrapeFunc(
	v reflect.Value,
	parentID string,
	level int,
) {
	if !v.CanAddr() {
		return
	}

	t := v.Type()
	for i := 0; i < t.NumOut(); i++ {
		v = reflect.NewAt(t.Out(i), unsafe.Pointer(v.UnsafeAddr())).Elem()
		s.scrape(v, parentID, level)
	}
}

func (s *scraper) scrapeNoop(
	_ reflect.Value,
	_ string,
	_ int,
) {
}

func (s *scraper) scrapeStruct(
	v reflect.Value,
	parentID string,
	level int,
) {
	if !s.isScrappable(v) {
		return
	}

	var c model.Component

	info, ok := s.getInfoFromInterface(v)
	if ok {
		c = s.addComponent(v, info, parentID)
	}

	info, ok = s.getInfoFromRules(v)
	if ok {
		c = s.addComponent(v, info, parentID)
	}

	if c.ID != "" {
		parentID = c.ID
	}

	s.scrapeValueFields(v, parentID, level)
	return
}

func (s *scraper) scrapeValueFields(
	v reflect.Value,
	parentID string,
	level int,
) {
	for i := 0; i < v.NumField(); i++ {
		s.scrape(v.Field(i), parentID, level+1)
	}
}

func (s *scraper) addComponent(
	v reflect.Value,
	info model.Info,
	parentID string,
) model.Component {
	c := model.Component{
		ID:          componentID(v),
		Kind:        info.Kind,
		Name:        info.Name,
		Description: info.Description,
		Technology:  info.Technology,
		Tags:        info.Tags,
	}
	s.structure.AddComponent(c, parentID)
	return c
}

func (s *scraper) isScrappable(v reflect.Value) bool {
	vPkg := valuePackage(v)
	for _, pkg := range s.config.Packages {
		if strings.HasPrefix(vPkg, pkg) {
			return true
		}
	}
	return false
}

func (s *scraper) getInfoFromInterface(v reflect.Value) (model.Info, bool) {
	var info model.HasInfo
	var ok bool

	if v.CanAddr() {
		// it allows accessing new pointer by the interface
		v = reflect.NewAt(v.Type(), unsafe.Pointer(v.UnsafeAddr())).Elem()
		// v.Addr() instead of v supports both value and pointer receiver
		info, ok = v.Addr().Interface().(model.HasInfo)
	} else if v.CanInterface() {
		info, ok = v.Interface().(model.HasInfo)
	} else {
		// it allows accessing new instance by the interface
		v = reflect.New(v.Type())
		info, ok = v.Interface().(model.HasInfo)
	}

	if !ok || info == nil {
		return model.Info{}, false
	}

	return info.Info(), true
}

func (s *scraper) getInfoFromRules(v reflect.Value) (model.Info, bool) {
	vPkg := valuePackage(v)
	name := componentName(v)
	for _, r := range s.rules {
		if !r.Applies(vPkg, name) {
			continue
		}
		return r.Apply(name), true
	}

	return model.Info{}, false
}

func componentID(v reflect.Value) string {
	id := fmt.Sprintf("%s.%s", valuePackage(v), v.Type().Name())
	return internal.Hash(id)
}

func componentName(v reflect.Value) string {
	pkg := strings.Split(valuePackage(v), "/")
	name := fmt.Sprintf("%s.%s", pkg[len(pkg)-1], v.Type().Name())
	return name
}

func valuePackage(v reflect.Value) string {
	return v.Type().PkgPath()
}
