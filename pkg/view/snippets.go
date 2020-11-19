package view

import (
	"fmt"
	"image/color"
	"strings"

	"github.com/krzysztofreczek/go-structurizr/pkg/model"
)

const (
	snippetUMLHead = `
@startuml`
	snippetUMLTail = `
@enduml`
	snippetUMLTitle = `
title {{title}}
`
	snippetSkinParamDefault = `
skinparam {
  shadowing false
  arrowFontSize 10
  defaultTextAlignment center
  wrapWidth 200
  maxMessageSize 100
}
hide stereotype
top to bottom direction
`
	snippetSkinParamRectangle = `
skinparam rectangle<<{{rec_name}}>> {
  BackgroundColor {{background_color_hash}}
  FontColor {{font_color_hash}}
  BorderColor {{border_color_hash}}
}
`
	snippetComponent = `
rectangle "=={{component_name}}\n<size:10>[{{component_kind}}{{component_technology}}]</size>\n\n{{component_desc}}" <<{{rec_name}}>> as {{component_id}}`
	snippetComponentConnection = `
{{component_id_from}} .[{{line_color_hash}}].> {{component_id_to}} : ""`

	paramComponentID          = "{{component_id}}"
	paramComponentIDFrom      = "{{component_id_from}}"
	paramComponentIDTo        = "{{component_id_to}}"
	paramComponentName        = "{{component_name}}"
	paramComponentKind        = "{{component_kind}}"
	paramComponentTechnology  = "{{component_technology}}"
	paramComponentDescription = "{{component_desc}}"
	paramTitle                = "{{title}}"
	paramRectangleName        = "{{rec_name}}"
	paramBackgroundColor      = "{{background_color_hash}}"
	paramFontColor            = "{{font_color_hash}}"
	paramBorderColor          = "{{border_color_hash}}"
	paramLineColor            = "{{line_color_hash}}"
)

func buildUMLHead() string {
	return snippetUMLHead
}

func buildUMLTail() string {
	return snippetUMLTail
}

func buildUMLTitle(
	title string,
) string {
	s := snippetUMLTitle
	s = strings.Replace(s, paramTitle, title, -1)
	return s
}

func buildSkinParamDefault() string {
	return snippetSkinParamDefault
}

func buildSkinParamRectangle(
	name string,
	backgroundColor color.Color,
	fontColor color.Color,
	borderColor color.Color,
) string {
	s := snippetSkinParamRectangle
	s = strings.Replace(s, paramRectangleName, name, -1)
	s = strings.Replace(s, paramBackgroundColor, toHex(backgroundColor), -1)
	s = strings.Replace(s, paramFontColor, toHex(fontColor), -1)
	s = strings.Replace(s, paramBorderColor, toHex(borderColor), -1)
	return s
}

func buildComponent(
	c model.Component,
) string {
	s := snippetComponent
	s = strings.Replace(s, paramComponentID, c.ID, -1)
	s = strings.Replace(s, paramComponentName, c.Name, -1)
	s = strings.Replace(s, paramComponentKind, c.Kind, -1)
	s = strings.Replace(s, paramComponentDescription, c.Description, -1)

	technology := c.Technology
	if technology != "" {
		technology = ":" + technology
	}
	s = strings.Replace(s, paramComponentTechnology, technology, -1)

	if len(c.Tags) > 0 {
		s = strings.Replace(s, paramRectangleName, c.Tags[0], -1)
	}

	return s
}

func buildComponentConnection(
	fromID string,
	toID string,
	lineColor color.Color,
) string {
	s := snippetComponentConnection
	s = strings.Replace(s, paramComponentIDFrom, fromID, -1)
	s = strings.Replace(s, paramComponentIDTo, toID, -1)
	s = strings.Replace(s, paramLineColor, toHex(lineColor), -1)
	return s
}

func toHex(c color.Color) string {
	rgba := color.RGBAModel.Convert(c).(color.RGBA)
	return fmt.Sprintf("#%.2x%.2x%.2x", rgba.R, rgba.G, rgba.B)
}
