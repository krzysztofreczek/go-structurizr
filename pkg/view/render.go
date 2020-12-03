package view

import (
	"io"
	"strings"

	"github.com/krzysztofreczek/go-structurizr/pkg/model"
)

// RenderStructureTo renders provided model.Structure into any io.Writer.
// RenderStructureTo will return an error in case the writer
// cannot be used.
func (v view) RenderStructureTo(s model.Structure, w io.Writer) error {
	out := v.render(s)
	_, err := w.Write([]byte(out))
	return err
}

func (v view) render(s model.Structure) string {
	sb := strings.Builder{}

	sb.WriteString(buildUMLHead())
	sb.WriteString(buildUMLTitle(v.title))
	sb.WriteString(buildSkinParamDefault())

	for _, s := range v.componentStyles {
		sb.WriteString(buildSkinParamShape(s.id, s.backgroundColor, s.fontColor, s.borderColor, s.shape))
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

		shape := defaultShape
		if len(c.Tags) > 0 {
			s, exists := v.componentStyles[c.Tags[0]]
			if exists {
				shape = s.shape
			}
		}

		sb.WriteString(buildComponent(c, shape))
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

func (v view) hasTag(tags ...string) bool {
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
