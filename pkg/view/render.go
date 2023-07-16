package view

import (
	"io"
	"strconv"
	"strings"

	"github.com/marianoceneri/go-structurizr/pkg/model"
)

const (
	defaultShape      = "rectangle"
	defaultShapeStyle = "DEFAULT"
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
	sb.WriteString(buildSkinParamGroup())

	for _, s := range v.componentStyles {
		sb.WriteString(buildSkinParamShape(s.id, s.backgroundColor, s.fontColor, s.borderColor, s.shape))
	}

	sb.WriteString(v.renderBody(s))
	sb.WriteString(buildUMLTail())

	return sb.String()
}

func (v view) renderBody(s model.Structure) string {
	ctx := v.newContext(s)

	v.renderRootComponents(ctx)

	for {
		ctx.level++
		rendered := v.renderNextBodyLayer(ctx)
		if rendered == 0 {
			break
		}
	}

	return ctx.sb.String()
}

type context struct {
	sb                strings.Builder
	s                 model.Structure
	excludedIDs       map[string]struct{}
	renderedIDs       map[string]struct{}
	renderedRelations map[string]struct{}
	level             int
}

func (v view) newContext(s model.Structure) *context {
	return &context{
		sb:                strings.Builder{},
		s:                 s,
		excludedIDs:       v.resolveExcludedComponentIDs(s),
		renderedIDs:       map[string]struct{}{},
		renderedRelations: map[string]struct{}{},
	}
}

func (v view) resolveExcludedComponentIDs(s model.Structure) map[string]struct{} {
	ids := map[string]struct{}{}
	for _, c := range s.Components {
		if !v.hasComponentTag(c.Tags...) {
			v.debug(c, "component will be excluded from the view")
			ids[c.ID] = struct{}{}
		}
	}
	return ids
}

func (v view) renderRootComponents(ctx *context) {
	for _, c := range ctx.s.Components {
		if !v.isRoot(c.Tags...) {
			continue
		}
		v.debug(c, "component has been recognised as root element")
		v.renderComponent(ctx, c, "")
	}
}

func (v view) renderNextBodyLayer(ctx *context) int {
	renderedPreviously := make(map[string]struct{})
	for id := range ctx.renderedIDs {
		renderedPreviously[id] = struct{}{}
	}

	for srcID := range renderedPreviously {
		srcRelations := ctx.s.Relations[srcID]

		for trgID := range srcRelations {
			c, exists := ctx.s.Components[trgID]
			if !exists {
				continue
			}

			v.renderComponent(ctx, c, srcID)
			v.renderRelation(ctx, srcID, trgID)
		}
	}

	componentsRendered := len(ctx.renderedIDs) - len(renderedPreviously)
	return componentsRendered
}

func (v view) renderComponent(ctx *context, c model.Component, parentID string) {
	_, excluded := ctx.excludedIDs[c.ID]
	if excluded {
		return
	}

	_, rendered := ctx.renderedIDs[c.ID]
	if rendered {
		return
	}

	shape := defaultShape
	shapeStyle := defaultShapeStyle
	if len(c.Tags) > 0 {
		shapeStyle = c.Tags[0]
		s, exists := v.componentStyles[shapeStyle]
		if exists {
			shape = s.shape
		}
	}

	group := groupID(parentID, shapeStyle, ctx.level)

	v.debug(c, "rendering component with shape '%s', shape style '%s', and group '%s'", shape, shapeStyle, group)

	ctx.sb.WriteString(buildComponent(c, shape, shapeStyle, group))
	ctx.renderedIDs[c.ID] = struct{}{}
}

func (v view) renderRelation(ctx *context, srcID string, trgID string) {
	_, rendered := ctx.renderedIDs[trgID]
	if !rendered {
		return
	}

	relationID := relationID(srcID, trgID)
	if _, rendered := ctx.renderedRelations[relationID]; rendered {
		return
	}

	v.debug(ctx.s.Components[srcID], "rendering relation to component of id '%s'", trgID)

	ctx.sb.WriteString(buildComponentConnection(srcID, trgID, v.lineColor))
	ctx.renderedRelations[relationID] = struct{}{}
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

func groupID(parentID string, style string, level int) string {
	return strings.Join([]string{parentID, strconv.Itoa(level), style}, "")
}

func relationID(srcID string, trgID string) string {
	return strings.Join([]string{srcID, trgID}, "")
}
