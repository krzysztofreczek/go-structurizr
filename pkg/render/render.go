package render

import (
	"fmt"
	"strings"

	"github.com/krzysztofreczek/go-structurizr/pkg/model"
)

func PlantUML(s model.Structure) string {
	sb := strings.Builder{}

	for name, _ := range s.Components {
		s := fmt.Sprintf("component [%s] as %s\n", name, name)
		sb.WriteString(s)
	}

	for src, targets := range s.Relations {
		for trg, _ := range targets {
			s := fmt.Sprintf("%s -> %s\n", src, trg)
			sb.WriteString(s)
		}
	}

	return sb.String()
}

func GraphViz(s model.Structure) string {
	sb := strings.Builder{}

	for src, targets := range s.Relations {
		for trg, _ := range targets {
			s := fmt.Sprintf("%s -> %s;\n", src, trg)
			sb.WriteString(s)
		}
	}

	return sb.String()
}
