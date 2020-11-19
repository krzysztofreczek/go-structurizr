package view

import (
	"image/color"

	"github.com/krzysztofreczek/go-structurizr/pkg/yaml"
	"github.com/pkg/errors"
)

type View struct {
	title           string
	componentStyles map[string]ComponentStyle
	lineColor       color.Color
}

func newView(
	title string,
	componentStyles map[string]ComponentStyle,
	lineColor color.Color,
) View {
	return View{
		title:           title,
		componentStyles: componentStyles,
		lineColor:       lineColor,
	}
}

type Builder struct {
	View
}

func NewView() *Builder {
	return &Builder{
		View: View{
			title:           "",
			componentStyles: make(map[string]ComponentStyle),
			lineColor:       color.Black,
		},
	}
}

func NewViewFromConfigFile(fileName string) (View, error) {
	configuration, err := yaml.LoadFromFile(fileName)
	if err != nil {
		return View{}, errors.Wrapf(err,
			"could not load configuration from file `%s`", fileName)
	}

	v, err := toView(configuration)
	if err != nil {
		return View{}, errors.Wrapf(err,
			"could not load view from file `%s`", fileName)
	}

	return v, nil
}

func (b *Builder) WithTitle(t string) *Builder {
	b.title = t
	return b
}

func (b *Builder) WithComponentStyle(s ComponentStyle) *Builder {
	b.componentStyles[s.id] = s
	return b
}

func (b *Builder) WithLineColor(c color.Color) *Builder {
	if c != nil {
		b.lineColor = c
	}
	return b
}

func (b Builder) Build() View {
	return newView(
		b.title,
		b.componentStyles,
		b.lineColor,
	)
}

type ComponentStyle struct {
	id              string
	backgroundColor color.Color
	fontColor       color.Color
	borderColor     color.Color
}

func newComponentStyle(
	id string,
	backgroundColor color.Color,
	fontColor color.Color,
	borderColor color.Color,
) ComponentStyle {
	return ComponentStyle{
		id:              id,
		backgroundColor: backgroundColor,
		fontColor:       fontColor,
		borderColor:     borderColor,
	}
}

type ComponentStyleBuilder struct {
	ComponentStyle
}

func NewComponentStyle(id string) *ComponentStyleBuilder {
	return &ComponentStyleBuilder{
		ComponentStyle: ComponentStyle{
			id:              id,
			backgroundColor: color.White,
			fontColor:       color.Black,
			borderColor:     color.Black,
		},
	}
}

func (b *ComponentStyleBuilder) WithBackgroundColor(c color.Color) *ComponentStyleBuilder {
	if c != nil {
		b.backgroundColor = c
	}
	return b
}

func (b *ComponentStyleBuilder) WithFontColor(c color.Color) *ComponentStyleBuilder {
	if c != nil {
		b.fontColor = c
	}
	return b
}

func (b *ComponentStyleBuilder) WithBorderColor(c color.Color) *ComponentStyleBuilder {
	if c != nil {
		b.borderColor = c
	}
	return b
}

func (b ComponentStyleBuilder) Build() ComponentStyle {
	return newComponentStyle(
		b.id,
		b.backgroundColor,
		b.fontColor,
		b.borderColor,
	)
}
