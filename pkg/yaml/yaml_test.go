package yaml_test

import (
	"strings"
	"testing"

	"github.com/krzysztofreczek/go-structurizr/pkg/yaml"
	"github.com/stretchr/testify/require"
)

const (
	testYAMLEmpty         = ``
	testYAMLConfiguration = `
configuration:
  pkgs: [PKG_1, PKG_2]
`

	testYAMLRules = `
rules:
  - pkg_regexps: [PKG_1, PKG_2]
    name_regexp: "^(\\w*)Client$"
    component:
      name: Client
      description: Client description
      technology: Client technology
      tags: [TAG_1, TAG_2]
  - pkg_regexps: [PKG_1, PKG_2]
    name_regexp: "^(\\w*)Repository$"
    component:
      name: Repository
      description: Repository description
      technology: Repository technology
      tags: [TAG_3]
`

	testYAMLViews = `
view:
  title: Title
  line_color: 000000ff
  styles:
    - id: STYLE_1
      background_color: ffffffff
      font_color: 000000ff
      border_color: 000000ff
  tags: [TAG_1, TAG_2]
`
)

func TestLoadFrom(t *testing.T) {
	var tests = []struct {
		name     string
		source   string
		expected yaml.Config
	}{
		{
			name:     "empty",
			source:   testYAMLEmpty,
			expected: yaml.Config{},
		},
		{
			name:   "configuration",
			source: testYAMLConfiguration,
			expected: yaml.Config{
				Configuration: yaml.ConfigConfiguration{
					Packages: []string{"PKG_1", "PKG_2"},
				},
			},
		},
		{
			name:   "rules",
			source: testYAMLRules,
			expected: yaml.Config{
				Rules: []yaml.ConfigRule{
					{
						PackageRegexps: []string{"PKG_1", "PKG_2"},
						NameRegexp:     `^(\w*)Client$`,
						Component: yaml.ConfigRuleComponent{
							Name:        "Client",
							Description: "Client description",
							Technology:  "Client technology",
							Tags:        []string{"TAG_1", "TAG_2"},
						},
					},
					{
						PackageRegexps: []string{"PKG_1", "PKG_2"},
						NameRegexp:     `^(\w*)Repository$`,
						Component: yaml.ConfigRuleComponent{
							Name:        "Repository",
							Description: "Repository description",
							Technology:  "Repository technology",
							Tags:        []string{"TAG_3"},
						},
					},
				},
			},
		},
		{
			name:   "view",
			source: testYAMLViews,
			expected: yaml.Config{
				View: yaml.ConfigView{
					Title:     "Title",
					LineColor: "000000ff",
					Styles: []yaml.ConfigViewStyle{
						{
							ID:              "STYLE_1",
							BackgroundColor: "ffffffff",
							FontColor:       "000000ff",
							BorderColor:     "000000ff",
						},
					},
					Tags: []string{"TAG_1", "TAG_2"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := strings.NewReader(tt.source)
			actual, err := yaml.LoadFrom(r)
			require.NoError(t, err)
			require.Equal(t, tt.expected, actual)
		})
	}
}
