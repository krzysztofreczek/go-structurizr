package scraper

import (
	"fmt"
	"log"
	"reflect"
)

func (s *scraper) debug(v reflect.Value, format string, a ...interface{}) {
	if !s.config.LogDebug {
		return
	}

	m := fmt.Sprintf(format, a...)
	log.Printf("[%s][id: %s] %s\n", componentName(v), componentID(v), m)
}
