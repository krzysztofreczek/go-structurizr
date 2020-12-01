package scraper

import (
	"testing"

	"github.com/krzysztofreczek/go-structurizr/pkg/model"
	"github.com/krzysztofreczek/go-structurizr/pkg/yaml"
	"github.com/stretchr/testify/require"
)

func Test_toScraperConfig(t *testing.T) {
	yamlConfiguration := yaml.Config{
		Configuration: yaml.ConfigConfiguration{
			Packages: []string{"PKG_1", "PKG_2"},
		},
	}

	c := toScraperConfig(yamlConfiguration)
	require.Equal(t, yamlConfiguration.Configuration.Packages, c.packages)
}

func Test_toScraperRules(t *testing.T) {
	yamlConfiguration := yaml.Config{
		Rules: []yaml.ConfigRule{
			{
				PackageRegexps: []string{"PKG_1", "PKG_2"},
				NameRegexp:     `^test.TestClient$`,
				Component: yaml.ConfigRuleComponent{
					Name:        "Client",
					Description: "Client description",
					Technology:  "Client technology",
					Tags:        []string{"TAG_1", "TAG_2"},
				},
			},
		},
	}

	rules, err := toScraperRules(yamlConfiguration)
	require.NoError(t, err)
	require.Len(t, rules, len(yamlConfiguration.Rules))

	r := rules[0]

	expectedRule, err := NewRule().
		WithPkgRegexps("PKG_1", "PKG_2").
		WithNameRegexp(`^test.TestClient$`).
		WithApplyFunc(func(name string, groups ...string) model.Info {
			return model.ComponentInfo(
				"Client",
				"Client description",
				"Client technology",
				"TAG_1",
				"TAG_2",
			)
		}).
		Build()
	require.NoError(t, err)

	require.Equal(t, expectedRule.Applies("", "test.TestClient"), r.Applies("", "test.TestClient"))
	require.Equal(t, expectedRule.Applies("PKG_1", ""), r.Applies("PKG_1", ""))
	require.Equal(t, expectedRule.Applies("PKG_1", "test.TestClient"), r.Applies("PKG_1", "test.TestClient"))
	require.Equal(t, expectedRule.Apply("test.TestClient"), r.Apply("test.TestClient"))
}
