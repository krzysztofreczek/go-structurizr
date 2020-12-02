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

	outString := string(out.Bytes())

	expectedContent := `
@startuml
title 

skinparam {
  shadowing false
  arrowFontSize 10
  defaultTextAlignment center
  wrapWidth 200
  maxMessageSize 100
}
hide stereotype
top to bottom direction

@enduml`

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

	outString := string(out.Bytes())

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

	outString := string(out.Bytes())

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

	outString := string(out.Bytes())

	expectedContent := `
rectangle "==test.Component\n<size:10>[component:technology]</size>\n\ndescription" <<tag 1>> as ID_1
`
	require.Contains(t, outString, expectedContent)
}

func TestNewView_with_relation(t *testing.T) {
	s := model.NewStructure()
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

	outString := string(out.Bytes())

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

	outString := string(out.Bytes())

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
		WithTag("tag 1").
		Build()
	err := v.RenderStructureTo(s, &out)
	require.NoError(t, err)

	outString := string(out.Bytes())

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
		WithTag("tag 1").
		Build()
	err := v.RenderStructureTo(s, &out)
	require.NoError(t, err)

	outString := string(out.Bytes())

	expectedContent := `
rectangle "==test.Component\n<size:10>[component:technology]</size>\n\ndescription" <<tag 1>> as ID_1
`
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
		WithTag("tag 1").
		WithTag("tag 2").
		Build()
	err := v.RenderStructureTo(s, &out)
	require.NoError(t, err)

	outString := string(out.Bytes())

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
		WithTag("tag 1").
		WithTag("tag 2").
		Build()
	err := v.RenderStructureTo(s, &out)
	require.NoError(t, err)

	outString := string(out.Bytes())

	expectedContent := `
rectangle "==test.Component\n<size:10>[component:technology]</size>\n\ndescription" <<tag 1>> as ID_1
`
	require.Contains(t, outString, expectedContent)

	expectedContent = `
rectangle "==test.Component\n<size:10>[component:technology]</size>\n\ndescription" <<tag 1>> as ID_2
`
	require.NotContains(t, outString, expectedContent)

	expectedContent = `
ID_1 .[#000000].> ID_2 : ""
`
	require.NotContains(t, outString, expectedContent)
}
