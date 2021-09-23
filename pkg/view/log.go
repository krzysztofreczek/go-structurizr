package view

import (
	"fmt"
	"log"
	"os"

	"github.com/krzysztofreczek/go-structurizr/pkg/model"
)

func (v view) isDebugMode() bool {
	level := os.Getenv("LOG_LEVEL")
	return level == "DEBUG" || level == "debug"
}

func (v view) debug(c model.Component, format string, a ...interface{}) {
	if !v.isDebugMode() {
		return
	}

	m := fmt.Sprintf(format, a...)
	log.Printf("[%s][id: %s] %s\n", c.Name, c.ID, m)
}
