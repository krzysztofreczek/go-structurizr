package view

import (
	"image/color"
	"io"

	"github.com/krzysztofreczek/go-structurizr/pkg/model"
	"github.com/krzysztofreczek/go-structurizr/pkg/yaml"
	"github.com/pkg/errors"
)

// View defines generic view.
//
// RenderStructureTo renders provided model.Structure into any io.Writer.
// RenderStructureTo will return an error in case the writer
// cannot be used.
type View interface {
	RenderStructureTo(s model.Structure, w io.Writer) error
}

type view struct {
	title           string
	tags            []string
	componentStyles map[string]ComponentStyle
	lineColor       color.Color
}

func newView(
	title string,
	tags []string,
	componentStyles map[string]ComponentStyle,
	lineColor color.Color,
) View {
	return view{
		title:           title,
		tags:            tags,
		componentStyles: componentStyles,
		lineColor:       lineColor,
	}
}

// NewView returns an empty Builder.
func NewView() Builder {
	return &builder{
		view: view{
			title:           "",
			tags:            make([]string, 0),
			componentStyles: make(map[string]ComponentStyle),
			lineColor:       color.Black,
		},
	}
}

// NewViewFromConfigFile instantiates a default View implementation
// with configuration loaded from provided YAML configuration file.
// NewViewFromConfigFile will return an error in case the YAML configuration
// file does not exist or contains invalid content.
func NewViewFromConfigFile(fileName string) (View, error) {
	configuration, err := yaml.LoadFromFile(fileName)
	if err != nil {
		return view{}, errors.Wrapf(err,
			"could not load configuration from file `%s`", fileName)
	}

	v, err := toView(configuration)
	if err != nil {
		return view{}, errors.Wrapf(err,
			"could not load view from file `%s`", fileName)
	}

	return v, nil
}

// Builder simplifies instantiation of default View implementation.
//
// WithTitle sets view title.
// WithTag adds tag to the view.
// If at least one tag is defines, view will contain only those components
// which are tagged with at least one of those tags.
// WithComponentStyle adds custom component style to the view.
// ComponentStyle will be applied to the components tagged
// with component style ID.
// WithLineColor sets custom line color.
//
// Build returns default View implementation constructed from
// the provided configuration.
// If not specified all colors are defaulted to either black or white.
type Builder interface {
	WithTitle(t string) Builder
	WithTag(t string) Builder
	WithComponentStyle(s ComponentStyle) Builder
	WithLineColor(c color.Color) Builder

	Build() View
}

type builder struct {
	view
}

// WithTitle sets view title.
func (b *builder) WithTitle(t string) Builder {
	b.title = t
	return b
}

// WithTag adds tag to the view.
// If at least one tag is defines, view will contain only those components
// which are tagged with at least one of those tags.
func (b *builder) WithTag(t string) Builder {
	b.tags = append(b.tags, t)
	return b
}

// WithComponentStyle adds custom component style to the view.
// ComponentStyle will be applied to the components tagged
// with component style ID.
func (b *builder) WithComponentStyle(s ComponentStyle) Builder {
	b.componentStyles[s.id] = s
	return b
}

// WithLineColor sets custom line color.
func (b *builder) WithLineColor(c color.Color) Builder {
	if c != nil {
		b.lineColor = c
	}
	return b
}

// Build returns default View implementation constructed from
// the provided configuration.
// If not specified all colors are defaulted to either black or white.
func (b builder) Build() View {
	return newView(
		b.title,
		b.tags,
		b.componentStyles,
		b.lineColor,
	)
}

// ComponentStyle is a structure that represents custom view style
// that can be applied to scraped components.
// ComponentStyle is applied to the components tagged with component style ID.
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

// ComponentStyleBuilder simplifies instantiation of default ComponentStyle
// implementation.
//
// WithBackgroundColor sets background color.
// WithFontColor sets font color.
// WithBorderColor sets border color
//
// Build returns default ComponentStyle implementation constructed from
// the provided configuration.
type ComponentStyleBuilder interface {
	WithBackgroundColor(c color.Color) ComponentStyleBuilder
	WithFontColor(c color.Color) ComponentStyleBuilder
	WithBorderColor(c color.Color) ComponentStyleBuilder

	Build() ComponentStyle
}

type componentStyleBuilder struct {
	ComponentStyle
}

// NewView returns ComponentStyleBuilder with provided id.
func NewComponentStyle(id string) ComponentStyleBuilder {
	return &componentStyleBuilder{
		ComponentStyle: ComponentStyle{
			id:              id,
			backgroundColor: color.White,
			fontColor:       color.Black,
			borderColor:     color.Black,
		},
	}
}

// WithBackgroundColor sets background color.
func (b *componentStyleBuilder) WithBackgroundColor(c color.Color) ComponentStyleBuilder {
	if c != nil {
		b.backgroundColor = c
	}
	return b
}

// WithFontColor sets font color.
func (b *componentStyleBuilder) WithFontColor(c color.Color) ComponentStyleBuilder {
	if c != nil {
		b.fontColor = c
	}
	return b
}

// WithBorderColor sets border color
func (b *componentStyleBuilder) WithBorderColor(c color.Color) ComponentStyleBuilder {
	if c != nil {
		b.borderColor = c
	}
	return b
}

// Build returns default ComponentStyle implementation constructed from
// the provided configuration.
func (b componentStyleBuilder) Build() ComponentStyle {
	return newComponentStyle(
		b.id,
		b.backgroundColor,
		b.fontColor,
		b.borderColor,
	)
}
