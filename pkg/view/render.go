package view

import (
	"image/color"
	"io"
	"strconv"
	"strings"

	"github.com/krzysztofreczek/go-structurizr/pkg/model"
	"github.com/pkg/errors"
)

const (
	defaultShape      = "rectangle"
	defaultShapeStyle = "DEFAULT"
)

type diagramWriter interface {
	writePreamble(title string)
	writeDefaultStyles()
	writeStyle(s ComponentStyle)
	writeComponent(c model.Component, shape, shapeStyle, group string)
	writeRelation(fromID, toID string, lineColor color.Color)
	writeEpilogue()
	output() string
}

// RenderStructureTo renders the provided `model.Structure` to any `io.Writer`
// using the configured diagram type. The default type is PlantUML.
//
// It returns an error if the writer cannot be used or the diagram type is unsupported.
func (v view) RenderStructureTo(s model.Structure, w io.Writer) error {
	dw, err := v.newDiagramWriter()
	if err != nil {
		return err
	}

	out := v.renderWith(s, dw)
	_, err = w.Write([]byte(out))
	return err
}

func (v view) newDiagramWriter() (diagramWriter, error) {
	switch v.diagramType {
	case DiagramMermaid:
		return newMermaidWriter(v.lineColor), nil
	case DiagramPlantUML, "":
		return newPlantUMLWriter(), nil
	default:
		return nil, errors.Errorf("unsupported diagram type `%s`", v.diagramType)
	}
}

func (v view) renderWith(s model.Structure, w diagramWriter) string {
	w.writePreamble(v.title)
	w.writeDefaultStyles()

	for _, style := range v.componentStyles {
		w.writeStyle(style)
	}

	v.walk(s, w)
	w.writeEpilogue()

	return w.output()
}

func (v view) walk(s model.Structure, w diagramWriter) {
	ctx := v.newContext(s, w)

	v.renderRootComponents(ctx)

	for {
		ctx.level++
		rendered := v.renderNextBodyLayer(ctx)
		if rendered == 0 {
			break
		}
	}
}

type context struct {
	w                 diagramWriter
	s                 model.Structure
	excludedIDs       map[string]struct{}
	renderedIDs       map[string]struct{}
	renderedRelations map[string]struct{}
	level             int
}

func (v view) newContext(s model.Structure, w diagramWriter) *context {
	return &context{
		w:                 w,
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

	ctx.w.writeComponent(c, shape, shapeStyle, group)
	ctx.renderedIDs[c.ID] = struct{}{}
}

func (v view) renderRelation(ctx *context, srcID string, trgID string) {
	_, rendered := ctx.renderedIDs[trgID]
	if !rendered {
		return
	}

	relID := relationID(srcID, trgID)
	if _, rendered := ctx.renderedRelations[relID]; rendered {
		return
	}

	v.debug(ctx.s.Components[srcID], "rendering relation to component of id '%s'", trgID)

	ctx.w.writeRelation(srcID, trgID, v.lineColor)
	ctx.renderedRelations[relID] = struct{}{}
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

type plantumlWriter struct {
	sb strings.Builder
}

func newPlantUMLWriter() *plantumlWriter {
	return &plantumlWriter{}
}

func (w *plantumlWriter) writePreamble(title string) {
	w.sb.WriteString(buildUMLHead())
	w.sb.WriteString(buildUMLTitle(title))
}

func (w *plantumlWriter) writeDefaultStyles() {
	w.sb.WriteString(buildSkinParamDefault())
	w.sb.WriteString(buildSkinParamGroup())
}

func (w *plantumlWriter) writeStyle(s ComponentStyle) {
	w.sb.WriteString(buildSkinParamShape(s.id, s.backgroundColor, s.fontColor, s.borderColor, s.shape))
}

func (w *plantumlWriter) writeComponent(c model.Component, shape, shapeStyle, group string) {
	w.sb.WriteString(buildComponent(c, shape, shapeStyle, group))
}

func (w *plantumlWriter) writeRelation(fromID, toID string, lineColor color.Color) {
	w.sb.WriteString(buildComponentConnection(fromID, toID, lineColor))
}

func (w *plantumlWriter) writeEpilogue() {
	w.sb.WriteString(buildUMLTail())
}

func (w *plantumlWriter) output() string {
	return w.sb.String()
}
