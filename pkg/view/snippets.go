package view

import (
	"fmt"
	"image/color"
	"strings"

	"github.com/krzysztofreczek/go-structurizr/pkg/model"
)

const (
	snippetUMLHead = `This diagram has been generated with go-structurizr 
[https://github.com/krzysztofreczek/go-structurizr]

@startuml
`
	snippetUMLTail = `
@enduml
`
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

scale 4096 width
`
	snippetSkinParamGroup = `
skinparam rectangle<<_GROUP>> {
  FontColor #ffffff
  BorderColor #ffffff
}

skinparam interface<<_INTERFACE>> {
  BackgroundColor #ffffff
  BorderColor #000000
}
`
	snippetSkinParamShape = `
skinparam {{shape}}<<{{shape_style}}>> {
  BackgroundColor {{background_color_hash}}
  FontColor {{font_color_hash}}
  BorderColor {{border_color_hash}}
}
`
	snippetComponent = `
rectangle {{group_name}} <<_GROUP>> {
	{{shape}} "=={{component_name}}\n<size:10>[{{component_kind}}{{component_technology}}]</size>\n\n{{component_desc}}" <<{{shape_style}}>> as {{component_id}}
}`
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
	paramGroupName            = "{{group_name}}"
	paramBackgroundColor      = "{{background_color_hash}}"
	paramFontColor            = "{{font_color_hash}}"
	paramBorderColor          = "{{border_color_hash}}"
	paramLineColor            = "{{line_color_hash}}"
	paramShape                = "{{shape}}"
	paramShapeStyle           = "{{shape_style}}"
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

func buildSkinParamGroup() string {
	return snippetSkinParamGroup
}

func buildSkinParamShape(
	name string,
	backgroundColor color.Color,
	fontColor color.Color,
	borderColor color.Color,
	shape string,
) string {
	s := snippetSkinParamShape
	s = strings.Replace(s, paramShapeStyle, name, -1)
	s = strings.Replace(s, paramBackgroundColor, toHex(backgroundColor), -1)
	s = strings.Replace(s, paramFontColor, toHex(fontColor), -1)
	s = strings.Replace(s, paramBorderColor, toHex(borderColor), -1)
	s = strings.Replace(s, paramShape, shape, -1)
	return s
}

func buildComponent(
	c model.Component,
	shape string,
	shapeStyle string,
	group string,
) string {
	s := snippetComponent
	s = strings.Replace(s, paramShape, shape, -1)
	s = strings.Replace(s, paramShapeStyle, shapeStyle, -1)
	s = strings.Replace(s, paramComponentID, c.ID, -1)
	s = strings.Replace(s, paramComponentName, c.Name, -1)
	s = strings.Replace(s, paramComponentKind, c.Kind, -1)
	s = strings.Replace(s, paramComponentDescription, c.Description, -1)
	s = strings.Replace(s, paramGroupName, group, -1)

	technology := c.Technology
	if technology != "" {
		technology = ":" + technology
	}
	s = strings.Replace(s, paramComponentTechnology, technology, -1)

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
