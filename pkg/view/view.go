package view

import (
	"image/color"
	"io"

	"github.com/krzysztofreczek/go-structurizr/pkg/model"
	"github.com/krzysztofreczek/go-structurizr/pkg/yaml"
	"github.com/pkg/errors"
)

// View defines a generic view for rendering structures.
//
// RenderStructureTo renders the provided `model.Structure` to any `io.Writer`.
//
// It returns an error if the writer cannot be used.
type View interface {
	RenderStructureTo(s model.Structure, w io.Writer) error
}

type view struct {
	title             string
	rootComponentTags []string
	componentTags     []string
	componentStyles   map[string]ComponentStyle
	lineColor         color.Color
}

func newView(
	title string,
	rootComponentTags []string,
	componentTags []string,
	componentStyles map[string]ComponentStyle,
	lineColor color.Color,
) View {
	return view{
		title:             title,
		rootComponentTags: rootComponentTags,
		componentTags:     componentTags,
		componentStyles:   componentStyles,
		lineColor:         lineColor,
	}
}

// NewView returns a new, empty View builder.
func NewView() Builder {
	return &builder{
		view: view{
			title:             "TITLE UNDEFINED",
			rootComponentTags: make([]string, 0),
			componentTags:     make([]string, 0),
			componentStyles:   make(map[string]ComponentStyle),
			lineColor:         color.Black,
		},
	}
}

// NewViewFromConfigFile creates a new View instance using configuration
// loaded from the specified YAML file.
//
// It returns an error if the YAML file does not exist or contains invalid content.
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

// Builder simplifies the creation of a default View implementation.
//
// WithTitle sets the view's title.
// WithRootComponentTag adds a root tag to the view. If at least one root tag is defined,
// the view will include only those components directly or indirectly connected
// to a component with a root tag.
// WithComponentTag adds a tag to the view. If at least one tag is defined,
// the view will include only those components tagged with at least one of these tags.
// WithComponentStyle adds custom styles for components. Styles are applied to components
// tagged with the specified style ID.
// WithLineColor sets a custom line color.
//
// Build returns a default View implementation based on the provided configuration.
// Colors default to black or white if not specified.
type Builder interface {
	WithTitle(t string) Builder
	WithRootComponentTag(t string) Builder
	WithComponentTag(t string) Builder
	WithComponentStyle(s ComponentStyle) Builder
	WithLineColor(c color.Color) Builder

	Build() View
}

type builder struct {
	view
}

// WithTitle sets the title of the view.
func (b *builder) WithTitle(t string) Builder {
	b.title = t
	return b
}

// WithRootComponentTag adds a root tag to the view.
//
// If at least one root tag is defined, the view will include only those components
// that are directly or indirectly connected to a component with a root tag.
func (b *builder) WithRootComponentTag(t string) Builder {
	b.rootComponentTags = append(b.rootComponentTags, t)
	return b
}

// WithComponentTag adds a tag to the view.
//
// If at least one tag is defined, the view will include only those components
// that are tagged with at least one of these tags.
func (b *builder) WithComponentTag(t string) Builder {
	b.componentTags = append(b.componentTags, t)
	return b
}

// WithComponentStyle adds a custom style to the view.
//
// The style will be applied to components that are tagged with the specified
// component style ID.
func (b *builder) WithComponentStyle(s ComponentStyle) Builder {
	b.componentStyles[s.id] = s
	return b
}

// WithLineColor sets a custom color for the lines.
func (b *builder) WithLineColor(c color.Color) Builder {
	if c != nil {
		b.lineColor = c
	}
	return b
}

// Build returns a default View implementation based on the provided configuration.
//
// If not specified, all colors default to black or white.
func (b builder) Build() View {
	return newView(
		b.title,
		b.rootComponentTags,
		b.componentTags,
		b.componentStyles,
		b.lineColor,
	)
}

// ComponentStyle represents a custom style for the view that can be applied
// to scraped components.
//
// The style is applied to components that are tagged with the corresponding
// component style ID.
type ComponentStyle struct {
	id              string
	backgroundColor color.Color
	fontColor       color.Color
	borderColor     color.Color
	shape           string
}

func newComponentStyle(
	id string,
	backgroundColor color.Color,
	fontColor color.Color,
	borderColor color.Color,
	shape string,
) ComponentStyle {
	return ComponentStyle{
		id:              id,
		backgroundColor: backgroundColor,
		fontColor:       fontColor,
		borderColor:     borderColor,
		shape:           shape,
	}
}

func newDefaultComponentStyle(
	id string,
) ComponentStyle {
	return ComponentStyle{
		id:              id,
		backgroundColor: color.White,
		fontColor:       color.Black,
		borderColor:     color.Black,
		shape:           defaultShape,
	}
}

// ComponentStyleBuilder simplifies the creation of a default ComponentStyle
// implementation.
//
// WithBackgroundColor sets the background color.
// WithFontColor sets the font color.
// WithBorderColor sets the border color.
// WithShape sets the component shape, corresponding to PlantUML shapes
// (e.g., rectangle, component, database). If no shape is specified, it defaults to rectangle.
//
// Build returns a default ComponentStyle implementation based on the provided configuration.
type ComponentStyleBuilder interface {
	WithBackgroundColor(c color.Color) ComponentStyleBuilder
	WithFontColor(c color.Color) ComponentStyleBuilder
	WithBorderColor(c color.Color) ComponentStyleBuilder
	WithShape(s string) ComponentStyleBuilder

	Build() ComponentStyle
}

type componentStyleBuilder struct {
	ComponentStyle
}

// NewComponentStyle returns a ComponentStyleBuilder with the specified ID.
func NewComponentStyle(id string) ComponentStyleBuilder {
	return &componentStyleBuilder{
		ComponentStyle: newDefaultComponentStyle(id),
	}
}

// WithBackgroundColor sets the background color of the component style.
func (b *componentStyleBuilder) WithBackgroundColor(c color.Color) ComponentStyleBuilder {
	if c != nil {
		b.backgroundColor = c
	}
	return b
}

// WithFontColor sets the font color of the component style.
func (b *componentStyleBuilder) WithFontColor(c color.Color) ComponentStyleBuilder {
	if c != nil {
		b.fontColor = c
	}
	return b
}

// WithBorderColor sets the border color of the component style.
func (b *componentStyleBuilder) WithBorderColor(c color.Color) ComponentStyleBuilder {
	if c != nil {
		b.borderColor = c
	}
	return b
}

// WithShape sets the component shape according to PlantUML shapes
// (e.g., rectangle, component, database). If no shape is specified,
// it defaults to rectangle.
func (b *componentStyleBuilder) WithShape(s string) ComponentStyleBuilder {
	if s != "" {
		b.shape = s
	}
	return b
}

// Build returns a default ComponentStyle implementation based on
// the provided configuration.
func (b componentStyleBuilder) Build() ComponentStyle {
	return newComponentStyle(
		b.id,
		b.backgroundColor,
		b.fontColor,
		b.borderColor,
		b.shape,
	)
}
