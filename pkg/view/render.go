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

	for _, c := range s.Components {
		sb.WriteString(buildComponent(c))
	}

	for src, to := range s.Relations {
		for trg, _ := range to {
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
