package view

import (
	"io"
	"strings"

	"github.com/krzysztofreczek/go-structurizr/pkg/model"
)

func (v View) render(s model.Structure) string {
	sb := strings.Builder{}

	sb.WriteString(buildUMLHead())
	sb.WriteString(buildUMLTitle(v.title))
	sb.WriteString(buildSkinParamDefault())

	for _, s := range v.componentStyles {
		sb.WriteString(buildSkinParamRectangle(s.id, s.backgroundColor, s.fontColor, s.borderColor))
	}

	excludedComponentIds := map[string]struct{}{}
	for _, c := range s.Components {
		if !v.hasTag(c.Tags...) {
			excludedComponentIds[c.ID] = struct{}{}
		}
	}

	for _, c := range s.Components {
		_, excluded := excludedComponentIds[c.ID]
		if excluded {
			continue
		}
		sb.WriteString(buildComponent(c))
	}

	for src, to := range s.Relations {
		for trg, _ := range to {
			_, srcExcluded := excludedComponentIds[src]
			if srcExcluded {
				continue
			}
			_, trgExcluded := excludedComponentIds[trg]
			if trgExcluded {
				continue
			}
			sb.WriteString(buildComponentConnection(src, trg, v.lineColor))
		}
	}

	sb.WriteString(buildUMLTail())

	return sb.String()
}

func (v View) RenderTo(s model.Structure, w io.Writer) error {
	out := v.render(s)
	_, err := w.Write([]byte(out))
	return err
}

func (v View) hasTag(tags ...string) bool {
	if len(v.tags) == 0 {
		return true
	}

	for _, vt := range v.tags {
		for _, t := range tags {
			if t == vt {
				return true
			}
		}
	}

	return false
}
