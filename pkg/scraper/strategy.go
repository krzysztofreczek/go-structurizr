package scraper

import (
	"fmt"
	"reflect"
	"strings"
	"unsafe"

	"github.com/krzysztofreczek/go-structurizr/pkg/internal"
	"github.com/krzysztofreczek/go-structurizr/pkg/model"
)

func (s *Scraper) scrap(
	v reflect.Value,
	parentID string,
	level int,
) {
	strategy := s.resolveScrapingStrategy(v)
	strategy(v, parentID, level)
}

type scrapingStrategy func(
	v reflect.Value,
	parentID string,
	level int,
)

func (s *Scraper) resolveScrapingStrategy(v reflect.Value) scrapingStrategy {
	switch v.Kind() {
	case reflect.Interface:
		return s.scrapeInterfaceStrategy
	case reflect.Ptr:
		return s.scrapePointerStrategy
	case reflect.Map:
		return s.scrapeMapStrategy
	case reflect.Slice, reflect.Array:
		return s.scrapeIterableStrategy
	}

	return s.scrapeValue
}

func (s *Scraper) scrapeInterfaceStrategy(
	v reflect.Value,
	parentID string,
	level int,
) {
	v = v.Elem()
	s.scrap(v, parentID, level)
}

func (s *Scraper) scrapePointerStrategy(
	v reflect.Value,
	parentID string,
	level int,
) {
	v = v.Elem()
	s.scrap(v, parentID, level)
}

func (s *Scraper) scrapeMapStrategy(
	v reflect.Value,
	parentID string,
	level int,
) {
	iterator := v.MapRange()
	for true {
		if !iterator.Next() {
			break
		}
		s.scrap(iterator.Value(), parentID, level)
	}
}

func (s *Scraper) scrapeIterableStrategy(
	v reflect.Value,
	parentID string,
	level int,
) {
	for i := 0; i < v.Len(); i++ {
		s.scrap(v.Index(i), parentID, level)
	}
}

func (s *Scraper) scrapeValue(
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

func (s *Scraper) scrapeValueFields(
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
		Name:        info.Name,
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

func (s *Scraper) getInfoFromRules(v reflect.Value) (model.Info, bool) {
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
