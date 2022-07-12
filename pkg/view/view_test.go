package view_test

import (
	"bytes"
	"image/color"
	"testing"

	"github.com/krzysztofreczek/go-structurizr/pkg/model"
	"github.com/krzysztofreczek/go-structurizr/pkg/view"
	"github.com/stretchr/testify/require"
)

func TestNewView_empty(t *testing.T) {
	s := model.NewStructure()

	out := bytes.Buffer{}

	v := view.NewView().Build()
	err := v.RenderStructureTo(s, &out)
	require.NoError(t, err)

	outString := out.String()

	expectedContent := `This diagram has been generated with go-structurizr 
[https://github.com/krzysztofreczek/go-structurizr]

@startuml

title TITLE UNDEFINED

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

skinparam rectangle<<_GROUP>> {
  FontColor #ffffff
  BorderColor #ffffff
}

@enduml
`

	require.Equal(t, expectedContent, outString)
}

func TestNewView_with_title(t *testing.T) {
	s := model.NewStructure()

	out := bytes.Buffer{}

	v := view.NewView().
		WithTitle("TITLE").
		Build()
	err := v.RenderStructureTo(s, &out)
	require.NoError(t, err)

	outString := out.String()

	expectedContent := `title TITLE`
	require.Contains(t, outString, expectedContent)
}

func TestNewView_with_custom_style(t *testing.T) {
	s := model.NewStructure()

	out := bytes.Buffer{}

	style := view.NewComponentStyle("STYLE").
		WithBackgroundColor(color.White).
		WithFontColor(color.Black).
		WithBorderColor(color.White).
		Build()
	v := view.NewView().
		WithComponentStyle(style).
		Build()
	err := v.RenderStructureTo(s, &out)
	require.NoError(t, err)

	outString := out.String()

	expectedContent := `
skinparam rectangle<<STYLE>> {
  BackgroundColor #ffffff
  FontColor #000000
  BorderColor #ffffff
}`
	require.Contains(t, outString, expectedContent)
}

func TestNewView_with_component(t *testing.T) {
	s := model.NewStructure()
	s.Components = map[string]model.Component{
		"ID_1": {
			ID:          "ID_1",
			Kind:        "component",
			Name:        "test.Component",
			Description: "description",
			Technology:  "technology",
			Tags:        []string{"tag 1", "tag 2"},
		},
	}

	out := bytes.Buffer{}

	v := view.NewView().Build()
	err := v.RenderStructureTo(s, &out)
	require.NoError(t, err)

	outString := out.String()

	expectedContent := `
	rectangle "==test.Component\n<size:10>[component:technology]</size>\n\ndescription" <<tag 1>> as ID_1
`
	require.Contains(t, outString, expectedContent)
}

func TestNewView_with_relation(t *testing.T) {
	s := model.NewStructure()
	s.Components = map[string]model.Component{
		"ID_1": {
			ID: "ID_1",
		},
		"ID_2": {
			ID: "ID_2",
		},
		"ID_3": {
			ID: "ID_3",
		},
	}
	s.Relations = map[string]map[string]struct{}{
		"ID_1": {
			"ID_2": {},
			"ID_3": {},
		},
	}

	out := bytes.Buffer{}

	v := view.NewView().Build()
	err := v.RenderStructureTo(s, &out)
	require.NoError(t, err)

	outString := out.String()

	expectedContent := `
ID_1 .[#000000].> ID_2 : ""
`
	require.Contains(t, outString, expectedContent)

	expectedContent = `
ID_1 .[#000000].> ID_3 : ""
`
	require.Contains(t, outString, expectedContent)
}

func TestNewView_with_custom_line_color(t *testing.T) {
	s := model.NewStructure()
	s.Components = map[string]model.Component{
		"ID_1": {
			ID: "ID_1",
		},
		"ID_2": {
			ID: "ID_2",
		},
	}
	s.Relations = map[string]map[string]struct{}{
		"ID_1": {
			"ID_2": {},
		},
	}

	out := bytes.Buffer{}

	v := view.NewView().
		WithLineColor(color.White).
		Build()
	err := v.RenderStructureTo(s, &out)
	require.NoError(t, err)

	outString := out.String()

	expectedContent := `
ID_1 .[#ffffff].> ID_2 : ""
`
	require.Contains(t, outString, expectedContent)
}

func TestNewView_with_component_of_view_tag(t *testing.T) {
	s := model.NewStructure()
	s.Components = map[string]model.Component{
		"ID_1": {
			ID:          "ID_1",
			Kind:        "component",
			Name:        "test.Component",
			Description: "description",
			Technology:  "technology",
			Tags:        []string{"tag 1", "tag 2"},
		},
	}

	out := bytes.Buffer{}

	v := view.NewView().
		WithComponentTag("tag 1").
		Build()
	err := v.RenderStructureTo(s, &out)
	require.NoError(t, err)

	outString := out.String()

	expectedContent := `
	rectangle "==test.Component\n<size:10>[component:technology]</size>\n\ndescription" <<tag 1>> as ID_1
`
	require.Contains(t, outString, expectedContent)
}

func TestNewView_with_component_with_no_view_tag(t *testing.T) {
	s := model.NewStructure()
	s.Components = map[string]model.Component{
		"ID_1": {
			ID:          "ID_1",
			Kind:        "component",
			Name:        "test.Component",
			Description: "description",
			Technology:  "technology",
			Tags:        []string{},
		},
	}

	out := bytes.Buffer{}

	v := view.NewView().
		WithComponentTag("tag 1").
		Build()
	err := v.RenderStructureTo(s, &out)
	require.NoError(t, err)

	outString := out.String()

	expectedContent := `ID_1`
	require.NotContains(t, outString, expectedContent)
}

func TestNewView_with_two_joined_components_of_view_tag(t *testing.T) {
	s := model.NewStructure()
	s.Components = map[string]model.Component{
		"ID_1": {
			ID:          "ID_1",
			Kind:        "component",
			Name:        "test.Component",
			Description: "description",
			Technology:  "technology",
			Tags:        []string{"tag 1"},
		},
		"ID_2": {
			ID:          "ID_2",
			Kind:        "component",
			Name:        "test.Component",
			Description: "description",
			Technology:  "technology",
			Tags:        []string{"tag 2"},
		},
	}
	s.Relations = map[string]map[string]struct{}{
		"ID_1": {
			"ID_2": {},
		},
	}

	out := bytes.Buffer{}

	v := view.NewView().
		WithComponentTag("tag 1").
		WithComponentTag("tag 2").
		Build()
	err := v.RenderStructureTo(s, &out)
	require.NoError(t, err)

	outString := out.String()

	expectedContent := `
	rectangle "==test.Component\n<size:10>[component:technology]</size>\n\ndescription" <<tag 1>> as ID_1
`
	require.Contains(t, outString, expectedContent)

	expectedContent = `
	rectangle "==test.Component\n<size:10>[component:technology]</size>\n\ndescription" <<tag 2>> as ID_2
`
	require.Contains(t, outString, expectedContent)

	expectedContent = `
ID_1 .[#000000].> ID_2 : ""
`
	require.Contains(t, outString, expectedContent)
}

func TestNewView_with_two_joined_components_where_one_with_no_view_tag(t *testing.T) {
	s := model.NewStructure()
	s.Components = map[string]model.Component{
		"ID_1": {
			ID:          "ID_1",
			Kind:        "component",
			Name:        "test.Component",
			Description: "description",
			Technology:  "technology",
			Tags:        []string{"tag 1"},
		},
		"ID_2": {
			ID:          "ID_2",
			Kind:        "component",
			Name:        "test.Component",
			Description: "description",
			Technology:  "technology",
			Tags:        []string{},
		},
	}
	s.Relations = map[string]map[string]struct{}{
		"ID_1": {
			"ID_2": {},
		},
	}

	out := bytes.Buffer{}

	v := view.NewView().
		WithComponentTag("tag 1").
		WithComponentTag("tag 2").
		Build()
	err := v.RenderStructureTo(s, &out)
	require.NoError(t, err)

	outString := out.String()

	expectedContent := `
	rectangle "==test.Component\n<size:10>[component:technology]</size>\n\ndescription" <<tag 1>> as ID_1
`
	require.Contains(t, outString, expectedContent)

	expectedContent = `ID_2`
	require.NotContains(t, outString, expectedContent)
}

func TestNewView_with_component_of_custom_style_shape(t *testing.T) {
	s := model.NewStructure()
	s.Components = map[string]model.Component{
		"ID_1": {
			ID:          "ID_1",
			Kind:        "component",
			Name:        "test.Component",
			Description: "description",
			Technology:  "technology",
			Tags:        []string{"DB"},
		},
	}

	out := bytes.Buffer{}

	style := view.NewComponentStyle("DB").
		WithBackgroundColor(color.White).
		WithFontColor(color.Black).
		WithBorderColor(color.White).
		WithShape("database").
		Build()
	v := view.NewView().
		WithComponentStyle(style).
		Build()
	err := v.RenderStructureTo(s, &out)
	require.NoError(t, err)

	outString := out.String()

	expectedContent := `
skinparam database<<DB>> {
  BackgroundColor #ffffff
  FontColor #000000
  BorderColor #ffffff
}`
	require.Contains(t, outString, expectedContent)

	expectedContent = `
	database "==test.Component\n<size:10>[component:technology]</size>\n\ndescription" <<DB>> as ID_1
`
	require.Contains(t, outString, expectedContent)
}

func TestNewView_with_two_joined_components_of_view_root_tag(t *testing.T) {
	s := model.NewStructure()
	s.Components = map[string]model.Component{
		"ID_1": {
			ID:          "ID_1",
			Kind:        "component",
			Name:        "test.Component",
			Description: "description",
			Technology:  "technology",
			Tags:        []string{"ROOT"},
		},
		"ID_2": {
			ID:          "ID_2",
			Kind:        "component",
			Name:        "test.Component",
			Description: "description",
			Technology:  "technology",
			Tags:        []string{},
		},
	}
	s.Relations = map[string]map[string]struct{}{
		"ID_1": {
			"ID_2": {},
		},
	}

	out := bytes.Buffer{}

	v := view.NewView().
		WithRootComponentTag("ROOT").
		Build()
	err := v.RenderStructureTo(s, &out)
	require.NoError(t, err)

	outString := out.String()

	expectedContent := `
	rectangle "==test.Component\n<size:10>[component:technology]</size>\n\ndescription" <<ROOT>> as ID_1
`
	require.Contains(t, outString, expectedContent)

	expectedContent = `
	rectangle "==test.Component\n<size:10>[component:technology]</size>\n\ndescription" <<DEFAULT>> as ID_2
`
	require.Contains(t, outString, expectedContent)

	expectedContent = `
ID_1 .[#000000].> ID_2 : ""
`
	require.Contains(t, outString, expectedContent)
}

func TestNewView_with_two_joined_components_where_one_with_no_view_root_tag(t *testing.T) {
	s := model.NewStructure()
	s.Components = map[string]model.Component{
		"ID_1": {
			ID:          "ID_1",
			Kind:        "component",
			Name:        "test.Component",
			Description: "description",
			Technology:  "technology",
			Tags:        []string{""},
		},
		"ID_2": {
			ID:          "ID_2",
			Kind:        "component",
			Name:        "test.Component",
			Description: "description",
			Technology:  "technology",
			Tags:        []string{},
		},
	}
	s.Relations = map[string]map[string]struct{}{
		"ID_1": {
			"ID_2": {},
		},
	}

	out := bytes.Buffer{}

	v := view.NewView().
		WithRootComponentTag("ROOT").
		Build()
	err := v.RenderStructureTo(s, &out)
	require.NoError(t, err)

	outString := out.String()

	expectedContent := `ID_1`
	require.NotContains(t, outString, expectedContent)

	expectedContent = `ID_2`
	require.NotContains(t, outString, expectedContent)
}

func TestNewView_with_component_with_no_connection_to_root(t *testing.T) {
	s := model.NewStructure()
	s.Components = map[string]model.Component{
		"ID_1": {
			ID:          "ID_1",
			Kind:        "component",
			Name:        "test.Component",
			Description: "description",
			Technology:  "technology",
			Tags:        []string{"ROOT"},
		},
		"ID_2": {
			ID:          "ID_2",
			Kind:        "component",
			Name:        "test.Component",
			Description: "description",
			Technology:  "technology",
			Tags:        []string{},
		},
	}
	s.Relations = map[string]map[string]struct{}{}

	out := bytes.Buffer{}

	v := view.NewView().
		WithRootComponentTag("ROOT").
		Build()
	err := v.RenderStructureTo(s, &out)
	require.NoError(t, err)

	outString := out.String()

	expectedContent := `
	rectangle "==test.Component\n<size:10>[component:technology]</size>\n\ndescription" <<ROOT>> as ID_1
`
	require.Contains(t, outString, expectedContent)

	expectedContent = `ID_2`
	require.NotContains(t, outString, expectedContent)
}

func TestNewView_creates_grouping(t *testing.T) {
	s := model.NewStructure()
	s.Components = map[string]model.Component{
		"ID_1": {
			ID:   "ID_1",
			Tags: []string{"ROOT"},
		},
		"ID_2": {
			ID:   "ID_2",
			Tags: []string{"TAG_A"},
		},
		"ID_3": {
			ID:   "ID_3",
			Tags: []string{"TAG_B"},
		},
	}
	s.Relations = map[string]map[string]struct{}{
		"ID_1": {
			"ID_2": {},
			"ID_3": {},
		},
	}

	out := bytes.Buffer{}

	v := view.NewView().
		WithRootComponentTag("ROOT").
		Build()
	err := v.RenderStructureTo(s, &out)
	require.NoError(t, err)

	outString := out.String()

	expectedContent := `
rectangle 0ROOT <<_GROUP>> {
	rectangle "==\n<size:10>[]</size>\n\n" <<ROOT>> as ID_1
}`
	require.Contains(t, outString, expectedContent)

	expectedContent = `
rectangle ID_11TAG_A <<_GROUP>> {
	rectangle "==\n<size:10>[]</size>\n\n" <<TAG_A>> as ID_2
}`
	require.Contains(t, outString, expectedContent)

	expectedContent = `
rectangle ID_11TAG_B <<_GROUP>> {
	rectangle "==\n<size:10>[]</size>\n\n" <<TAG_B>> as ID_3
}`
	require.Contains(t, outString, expectedContent)
}
