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
		if !v.hasComponentTag(c.Tags...) {
			excludedComponentIds[c.ID] = struct{}{}
		}
	}

	componentsRendered := map[string]struct{}{}

	for _, c := range s.Components {
		_, excluded := excludedComponentIds[c.ID]
		if excluded {
			continue
		}

		if !v.isRoot(c.Tags...) {
			continue
		}

		v.renderComponent(&sb, c)
		componentsRendered[c.ID] = struct{}{}
	}

	for true {
		numRendered := len(componentsRendered)

		for src, to := range s.Relations {
			for trg := range to {
				if _, srcExcluded := excludedComponentIds[src]; srcExcluded {
					continue
				}

				if _, trgExcluded := excludedComponentIds[trg]; trgExcluded {
					continue
				}

				_, srcRendered := componentsRendered[src]
				if !srcRendered {
					continue
				}

				_, trgRendered := componentsRendered[trg]
				if !trgRendered {
					c, exists := s.Components[trg]
					if !exists {
						continue
					}

					v.renderComponent(&sb, c)
					componentsRendered[c.ID] = struct{}{}
				}

				sb.WriteString(buildComponentConnection(src, trg, v.lineColor))
			}
		}

		if numRendered == len(componentsRendered) {
			break
		}
	}

	sb.WriteString(buildUMLTail())

	return sb.String()
}

func (v view) renderComponent(sb *strings.Builder, c model.Component) {
	shape := defaultShape
	if len(c.Tags) > 0 {
		s, exists := v.componentStyles[c.Tags[0]]
		if exists {
			shape = s.shape
		}
	}

	sb.WriteString(buildComponent(c, shape))
}

func (v view) isRoot(tags ...string) bool {
	if len(v.rootComponentTags) == 0 {
		return true
	}

	for _, vt := range v.rootComponentTags {
		for _, t := range tags {
			if t == vt {
				return true
			}
		}
	}

	return false
}

func (v view) hasComponentTag(tags ...string) bool {
	if len(v.componentTags) == 0 {
		return true
	}

	for _, vt := range v.componentTags {
		for _, t := range tags {
			if t == vt {
				return true
			}
		}
	}

	return false
}
