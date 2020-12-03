package view

import (
	"bytes"
	"image/color"
	"testing"

	"github.com/krzysztofreczek/go-structurizr/pkg/model"
	"github.com/krzysztofreczek/go-structurizr/pkg/yaml"
	"github.com/stretchr/testify/require"
)

func Test_toView(t *testing.T) {
	yamlConfiguration := yaml.Config{
		View: yaml.ConfigView{
			Title:     "TITLE_1",
			LineColor: "000000ff",
			Styles: []yaml.ConfigViewStyle{
				{
					ID:              "STYLE_1",
					BackgroundColor: "ffffffff",
					FontColor:       "000000ff",
					BorderColor:     "000000ff",
				},
			},
			ComponentTags:     []string{"TAG_1"},
			RootComponentTags: []string{"TAG_2"},
		},
	}

	actualView, err := toView(yamlConfiguration)
	require.NoError(t, err)

	expectedView := NewView().
		WithTitle("TITLE_1").
		WithLineColor(color.Black).
		WithComponentStyle(
			NewComponentStyle("STYLE_1").
				WithBackgroundColor(color.White).
				WithFontColor(color.Black).
				WithBorderColor(color.Black).
				Build(),
		).
		WithComponentTag("TAG_1").
		WithRootComponentTag("TAG_2").
		Build()

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

	actualOutput := bytes.Buffer{}
	err = actualView.RenderStructureTo(s, &actualOutput)
	require.NoError(t, err)
	require.NotEmpty(t, actualOutput)

	expectedOutput := bytes.Buffer{}
	err = expectedView.RenderStructureTo(s, &expectedOutput)
	require.NoError(t, err)
	require.NotEmpty(t, expectedOutput)

	require.Equal(t, expectedOutput, actualOutput)
}
