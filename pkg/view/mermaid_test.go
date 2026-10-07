package view_test

import (
	"bytes"
	"image/color"
	"testing"

	"github.com/krzysztofreczek/go-structurizr/pkg/model"
	"github.com/krzysztofreczek/go-structurizr/pkg/view"
	"github.com/stretchr/testify/require"
)

func newMermaidView() view.Builder {
	return view.NewView().WithDiagramType(view.DiagramMermaid)
}

func TestNewView_mermaid_empty(t *testing.T) {
	s := model.NewStructure()

	out := bytes.Buffer{}

	v := newMermaidView().Build()
	err := v.RenderStructureTo(s, &out)
	require.NoError(t, err)

	expectedContent := `---
title: "TITLE UNDEFINED"
---
%% This diagram has been generated with go-structurizr
%% https://github.com/krzysztofreczek/go-structurizr
%% title: TITLE UNDEFINED
flowchart TB
  linkStyle default stroke:#000000
`

	require.Equal(t, expectedContent, out.String())
}

func TestNewView_mermaid_with_title(t *testing.T) {
	s := model.NewStructure()

	out := bytes.Buffer{}

	v := newMermaidView().
		WithTitle("TITLE").
		Build()
	err := v.RenderStructureTo(s, &out)
	require.NoError(t, err)

	outString := out.String()
	require.Contains(t, outString, `title: "TITLE"`)
	require.Contains(t, outString, `%% title: TITLE`)
}

func TestNewView_mermaid_with_custom_style(t *testing.T) {
	s := model.NewStructure()

	out := bytes.Buffer{}

	style := view.NewComponentStyle("STYLE").
		WithBackgroundColor(color.White).
		WithFontColor(color.Black).
		WithBorderColor(color.White).
		Build()
	v := newMermaidView().
		WithComponentStyle(style).
		Build()
	err := v.RenderStructureTo(s, &out)
	require.NoError(t, err)

	require.Contains(t, out.String(), "  classDef STYLE fill:#ffffff,color:#000000,stroke:#ffffff\n")
}

func TestNewView_mermaid_with_component(t *testing.T) {
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

	v := newMermaidView().Build()
	err := v.RenderStructureTo(s, &out)
	require.NoError(t, err)

	outString := out.String()
	require.Contains(t, outString, `  ID_1["test.Component<br/>[component:technology]<br/><br/>description"]`)
	require.Contains(t, outString, "  class ID_1 tag_1\n")
}

func TestNewView_mermaid_with_relation(t *testing.T) {
	s := model.NewStructure()
	s.Components = map[string]model.Component{
		"ID_1": {ID: "ID_1"},
		"ID_2": {ID: "ID_2"},
		"ID_3": {ID: "ID_3"},
	}
	s.Relations = map[string]map[string]struct{}{
		"ID_1": {
			"ID_2": {},
			"ID_3": {},
		},
	}

	out := bytes.Buffer{}

	v := newMermaidView().Build()
	err := v.RenderStructureTo(s, &out)
	require.NoError(t, err)

	outString := out.String()
	require.Contains(t, outString, "  ID_1 -.-> ID_2\n")
	require.Contains(t, outString, "  ID_1 -.-> ID_3\n")
}

func TestNewView_mermaid_with_custom_line_color(t *testing.T) {
	s := model.NewStructure()
	s.Components = map[string]model.Component{
		"ID_1": {ID: "ID_1"},
		"ID_2": {ID: "ID_2"},
	}
	s.Relations = map[string]map[string]struct{}{
		"ID_1": {
			"ID_2": {},
		},
	}

	out := bytes.Buffer{}

	v := newMermaidView().
		WithLineColor(color.White).
		Build()
	err := v.RenderStructureTo(s, &out)
	require.NoError(t, err)

	require.Contains(t, out.String(), "  linkStyle default stroke:#ffffff\n")
}

func TestNewView_mermaid_with_component_of_view_tag(t *testing.T) {
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

	v := newMermaidView().
		WithComponentTag("tag 1").
		Build()
	err := v.RenderStructureTo(s, &out)
	require.NoError(t, err)

	require.Contains(t, out.String(), `  ID_1["test.Component<br/>[component:technology]<br/><br/>description"]`)
}

func TestNewView_mermaid_with_component_with_no_view_tag(t *testing.T) {
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

	v := newMermaidView().
		WithComponentTag("tag 1").
		Build()
	err := v.RenderStructureTo(s, &out)
	require.NoError(t, err)

	require.NotContains(t, out.String(), "ID_1")
}

func TestNewView_mermaid_with_two_joined_components_of_view_tag(t *testing.T) {
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

	v := newMermaidView().
		WithComponentTag("tag 1").
		WithComponentTag("tag 2").
		Build()
	err := v.RenderStructureTo(s, &out)
	require.NoError(t, err)

	outString := out.String()
	require.Contains(t, outString, "  class ID_1 tag_1\n")
	require.Contains(t, outString, "  class ID_2 tag_2\n")
	require.Contains(t, outString, "  ID_1 -.-> ID_2\n")
}

func TestNewView_mermaid_with_two_joined_components_where_one_with_no_view_tag(t *testing.T) {
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

	v := newMermaidView().
		WithComponentTag("tag 1").
		Build()
	err := v.RenderStructureTo(s, &out)
	require.NoError(t, err)

	outString := out.String()
	require.Contains(t, outString, "ID_1")
	require.NotContains(t, outString, "ID_2")
}

func TestNewView_mermaid_with_component_of_custom_style_shape(t *testing.T) {
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
	v := newMermaidView().
		WithComponentStyle(style).
		Build()
	err := v.RenderStructureTo(s, &out)
	require.NoError(t, err)

	outString := out.String()
	require.Contains(t, outString, "  classDef DB fill:#ffffff,color:#000000,stroke:#ffffff\n")
	require.Contains(t, outString, `  ID_1[("test.Component<br/>[component:technology]<br/><br/>description")]`)
	require.Contains(t, outString, "  class ID_1 DB\n")
}

func TestNewView_mermaid_with_component_shape(t *testing.T) {
	s := model.NewStructure()
	s.Components = map[string]model.Component{
		"ID_1": {
			ID:   "ID_1",
			Name: "queue",
			Tags: []string{"Q"},
		},
	}

	style := view.NewComponentStyle("Q").
		WithShape("queue").
		Build()
	out := bytes.Buffer{}
	v := newMermaidView().
		WithComponentStyle(style).
		Build()
	err := v.RenderStructureTo(s, &out)
	require.NoError(t, err)

	require.Contains(t, out.String(), `  ID_1[/"queue<br/>[]<br/><br/>"/]`)
}

func TestNewView_mermaid_with_two_joined_components_of_view_root_tag(t *testing.T) {
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

	v := newMermaidView().
		WithRootComponentTag("ROOT").
		Build()
	err := v.RenderStructureTo(s, &out)
	require.NoError(t, err)

	outString := out.String()
	require.Contains(t, outString, "  class ID_1 ROOT\n")
	require.Contains(t, outString, "  class ID_2 DEFAULT\n")
	require.Contains(t, outString, "  ID_1 -.-> ID_2\n")
}

func TestNewView_mermaid_with_two_joined_components_where_one_with_no_view_root_tag(t *testing.T) {
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

	v := newMermaidView().
		WithRootComponentTag("ROOT").
		Build()
	err := v.RenderStructureTo(s, &out)
	require.NoError(t, err)

	outString := out.String()
	require.NotContains(t, outString, "ID_1")
	require.NotContains(t, outString, "ID_2")
}

func TestNewView_mermaid_with_component_with_no_connection_to_root(t *testing.T) {
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

	v := newMermaidView().
		WithRootComponentTag("ROOT").
		Build()
	err := v.RenderStructureTo(s, &out)
	require.NoError(t, err)

	outString := out.String()
	require.Contains(t, outString, "ID_1")
	require.NotContains(t, outString, "ID_2")
}

func TestNewView_mermaid_sanitizes_numeric_ids(t *testing.T) {
	s := model.NewStructure()
	s.Components = map[string]model.Component{
		"123": {
			ID:   "123",
			Name: "hashed",
		},
		"456": {
			ID:   "456",
			Name: "other",
		},
	}
	s.Relations = map[string]map[string]struct{}{
		"123": {
			"456": {},
		},
	}

	out := bytes.Buffer{}

	v := newMermaidView().Build()
	err := v.RenderStructureTo(s, &out)
	require.NoError(t, err)

	outString := out.String()
	require.Contains(t, outString, `  n123["hashed<br/>[]<br/><br/>"]`)
	require.Contains(t, outString, `  n456["other<br/>[]<br/><br/>"]`)
	require.Contains(t, outString, "  n123 -.-> n456\n")
}

func TestNewView_mermaid_sanitizes_special_characters_in_labels(t *testing.T) {
	s := model.NewStructure()
	s.Components = map[string]model.Component{
		"ID_1": {
			ID:          "ID_1",
			Kind:        "component",
			Name:        `foo"bar` + "`baz",
			Description: "line1\nline2",
			Technology:  "tech",
		},
	}

	out := bytes.Buffer{}

	v := newMermaidView().Build()
	err := v.RenderStructureTo(s, &out)
	require.NoError(t, err)

	require.Contains(t, out.String(), `  ID_1["foo'bar'baz<br/>[component:tech]<br/><br/>line1<br/>line2"]`)
}

func TestNewView_mermaid_escapes_quotes_in_title(t *testing.T) {
	s := model.NewStructure()

	out := bytes.Buffer{}

	v := newMermaidView().
		WithTitle(`Title "quoted"`).
		Build()
	err := v.RenderStructureTo(s, &out)
	require.NoError(t, err)

	require.Contains(t, out.String(), `title: "Title \"quoted\""`)
}
