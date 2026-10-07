package view

import (
	"fmt"
	"image/color"
	"strings"
	"unicode"

	"github.com/krzysztofreczek/go-structurizr/pkg/model"
)

type mermaidWriter struct {
	sb        strings.Builder
	lineColor color.Color
}

func newMermaidWriter(lineColor color.Color) *mermaidWriter {
	return &mermaidWriter{lineColor: lineColor}
}

func (w *mermaidWriter) writePreamble(title string) {
	escaped := escapeMermaidFrontmatterTitle(title)
	s := snippetMermaidHead
	s = strings.ReplaceAll(s, paramTitle, escaped)
	w.sb.WriteString(s)
}

func (w *mermaidWriter) writeDefaultStyles() {}

func (w *mermaidWriter) writeStyle(s ComponentStyle) {
	out := snippetMermaidClassDef
	out = strings.ReplaceAll(out, paramShapeStyle, mermaidID(s.id))
	out = strings.ReplaceAll(out, paramBackgroundColor, toHex(s.backgroundColor))
	out = strings.ReplaceAll(out, paramFontColor, toHex(s.fontColor))
	out = strings.ReplaceAll(out, paramBorderColor, toHex(s.borderColor))
	w.sb.WriteString(out)
}

func (w *mermaidWriter) writeComponent(c model.Component, shape, shapeStyle, _ string) {
	id := mermaidID(c.ID)
	label := mermaidLabel(c)
	node := mermaidNode(id, label, shape)

	s := snippetMermaidComponent
	s = strings.ReplaceAll(s, paramNode, node)
	s = strings.ReplaceAll(s, paramComponentID, id)
	s = strings.ReplaceAll(s, paramShapeStyle, mermaidID(shapeStyle))
	w.sb.WriteString(s)
}

func (w *mermaidWriter) writeRelation(fromID, toID string, _ color.Color) {
	s := snippetMermaidConnection
	s = strings.ReplaceAll(s, paramComponentIDFrom, mermaidID(fromID))
	s = strings.ReplaceAll(s, paramComponentIDTo, mermaidID(toID))
	w.sb.WriteString(s)
}

func (w *mermaidWriter) writeEpilogue() {
	s := snippetMermaidLinkStyle
	s = strings.ReplaceAll(s, paramLineColor, toHex(w.lineColor))
	w.sb.WriteString(s)
}

func (w *mermaidWriter) output() string {
	return w.sb.String()
}

func mermaidLabel(c model.Component) string {
	kindTech := escapeMermaidLabel(c.Kind)
	technology := escapeMermaidLabel(c.Technology)
	if technology != "" {
		kindTech += ":" + technology
	}

	return strings.Join([]string{
		escapeMermaidLabel(c.Name),
		"[" + kindTech + "]",
		"",
		escapeMermaidLabel(c.Description),
	}, "<br/>")
}

func mermaidNode(id, label, shape string) string {
	switch shape {
	case "database":
		return fmt.Sprintf(`%s[("%s")]`, id, label)
	case "component":
		return fmt.Sprintf(`%s[["%s"]]`, id, label)
	case "queue":
		return fmt.Sprintf(`%s[/"%s"/]`, id, label)
	default:
		return fmt.Sprintf(`%s["%s"]`, id, label)
	}
}

func mermaidID(id string) string {
	var b strings.Builder
	b.Grow(len(id) + 1)

	first := true
	for _, r := range id {
		switch {
		case r == '_' || unicode.IsLetter(r):
			b.WriteRune(r)
		case unicode.IsDigit(r):
			if first {
				b.WriteByte('n')
			}
			b.WriteRune(r)
		default:
			if first {
				b.WriteByte('n')
			}
			b.WriteByte('_')
		}
		first = false
	}

	if b.Len() == 0 {
		return "n"
	}
	return b.String()
}

func escapeMermaidLabel(s string) string {
	s = strings.ReplaceAll(s, `"`, "'")
	s = strings.ReplaceAll(s, "`", "'")
	s = strings.ReplaceAll(s, "\r\n", "<br/>")
	s = strings.ReplaceAll(s, "\n", "<br/>")
	s = strings.ReplaceAll(s, "\r", "")
	return s
}

func escapeMermaidFrontmatterTitle(s string) string {
	s = strings.ReplaceAll(s, `\`, `\\`)
	s = strings.ReplaceAll(s, `"`, `\"`)
	s = strings.ReplaceAll(s, "\n", " ")
	s = strings.ReplaceAll(s, "\r", "")
	return s
}
