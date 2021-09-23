package scraper

import (
	"fmt"
	"log"
	"os"
	"reflect"
)

func (s *scraper) isDebugMode() bool {
	level := os.Getenv("LOG_LEVEL")
	return level == "DEBUG" || level == "debug"
}

func (s *scraper) debug(v reflect.Value, format string, a ...interface{}) {
	if !s.isDebugMode() {
		return
	}

	m := fmt.Sprintf(format, a...)
	log.Printf("[%s][id: %s] %s\n", componentName(v), componentID(v), m)
}
