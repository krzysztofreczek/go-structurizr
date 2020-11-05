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

func StructurizrDSL(s model.Structure) string {
	sb := strings.Builder{}

	for _, c := range s.Components {
		tags := ""
		for _, t := range c.Tags {
			tags += fmt.Sprintf(` "%s"`, t)
		}
		s := fmt.Sprintf(`%s = %s "%s" "%s" "%s" %s`, c.ID, c.Kind, c.Name, c.Description, c.Technology, tags)
		sb.WriteString(s)
		sb.WriteString("\n")
	}

	for from, tos := range s.Relations {
		for to := range tos {
			s := fmt.Sprintf("%s -> %s\n", from, to)
			sb.WriteString(s)
		}
	}

	return sb.String()
}
