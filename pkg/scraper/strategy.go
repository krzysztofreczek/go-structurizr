package scraper

import (
	"fmt"
	"reflect"
	"strings"
	"unsafe"

	"github.com/krzysztofreczek/go-structurizr/pkg/internal"
	"github.com/krzysztofreczek/go-structurizr/pkg/model"
)

var (
	maxRecursiveScrapes = 100
)

func (s *scraper) scrape(
	v reflect.Value,
	parentID string,
	level int,
) {
	if !v.IsValid() {
		s.debug(v, "value is not valid")
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
	s.debug(v, "interface scraping strategy applied: if the interface is not nil, the value will be scraped, otherwise scraper will try to resolve info data from the interface type")

	if !v.Elem().IsValid() {
		s.debug(v, "scraping the interface type")

		info, ok := s.getInfoFromRules(v)
		if ok {
			_ = s.addComponent(v, info, parentID)
		}

		return
	}

	s.debug(v, "scraping the value implementing the interface")
	v = v.Elem()

	s.scrape(v, parentID, level)
}

func (s *scraper) scrapePointerStrategy(
	v reflect.Value,
	parentID string,
	level int,
) {
	s.debug(v, "pointer scraping strategy applied: if the pointer is not nil, the value will be scraped")

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
	s.debug(v, "map scraping strategy applied: each of map elements will be scraped")

	iterator := v.MapRange()
	for {
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
	s.debug(v, "iterable scraping strategy applied: each of elements will be scraped")

	for i := 0; i < v.Len(); i++ {
		s.scrape(v.Index(i), parentID, level)
	}
}

func (s *scraper) scrapeFunc(
	v reflect.Value,
	parentID string,
	level int,
) {
	s.debug(v, "function scraping strategy applied: output types will be scraped")

	t := v.Type()
	for i := 0; i < t.NumOut(); i++ {
		v = reflect.New(t.Out(i))
		s.scrape(v, parentID, level)
	}
}

func (s *scraper) scrapeNoop(
	v reflect.Value,
	_ string,
	_ int,
) {
	s.debug(v, "value will not be scraped")
}

func (s *scraper) scrapeStruct(
	v reflect.Value,
	parentID string,
	level int,
) {
	s.debug(v, "struct scraping strategy applied: value and each of its properties will be scraped")

	if !s.isScrappable(v) {
		return
	}

	vID := componentID(v)
	if c, ok := s.typeCounters[vID]; ok && c > maxRecursiveScrapes {
		s.debug(v, "struct is being used recursively, skipping")
		return
	} else {
		s.typeCounters[vID]++
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
			s.debug(v, "value package '%s' is applicable for scraping", vPkg)
			return true
		}
	}

	s.debug(v, "value package '%s' IS NOT applicable for scraping", vPkg)
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

	i := info.Info()
	s.debug(v, "resolved info data %+v from .Info() method", i)

	return i, true
}

func (s *scraper) getInfoFromRules(v reflect.Value) (model.Info, bool) {
	vPkg := valuePackage(v)
	name := componentName(v)
	for _, r := range s.rules {
		if !r.Applies(vPkg, name) {
			continue
		}

		i := r.Apply(name)
		s.debug(v, "resolved info data %+v from one of the rules", i)

		return i, true
	}

	s.debug(v, "there was no rule applicable for this value")
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
