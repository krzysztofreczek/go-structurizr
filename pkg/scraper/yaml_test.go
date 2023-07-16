package scraper

import (
	"fmt"
	"testing"

	"github.com/marianoceneri/go-structurizr/pkg/model"
	"github.com/marianoceneri/go-structurizr/pkg/yaml"
	"github.com/stretchr/testify/require"
)

func Test_toScraperConfig(t *testing.T) {
	yamlConfiguration := yaml.Config{
		Configuration: yaml.ConfigConfiguration{
			Packages: []string{"PKG_1", "PKG_2"},
		},
	}

	c := toScraperConfig(yamlConfiguration)
	require.Equal(t, yamlConfiguration.Configuration.Packages, c.Packages)
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

	yamlRule := yamlConfiguration.Rules[0]

	rules, err := toScraperRules(yamlConfiguration)
	require.NoError(t, err)
	require.Len(t, rules, len(yamlConfiguration.Rules))

	r := rules[0]

	expectedRule, err := NewRule().
		WithPkgRegexps(yamlRule.PackageRegexps...).
		WithNameRegexp(yamlRule.NameRegexp).
		WithApplyFunc(func(name string, groups ...string) model.Info {
			return model.ComponentInfo(
				yamlRule.Component.Name,
				yamlRule.Component.Description,
				yamlRule.Component.Technology,
				yamlRule.Component.Tags[0],
				yamlRule.Component.Tags[1],
			)
		}).
		Build()
	require.NoError(t, err)

	pkg := "PKG_1"
	name := "test.TestClient"
	require.Equal(t, expectedRule.Applies("", name), r.Applies("", name))
	require.Equal(t, expectedRule.Applies(pkg, ""), r.Applies(pkg, ""))
	require.Equal(t, expectedRule.Applies(pkg, name), r.Applies(pkg, name))
	require.Equal(t, expectedRule.Apply(name), r.Apply(name))
}

func Test_toScraperRules_with_name_aliases(t *testing.T) {
	yamlConfiguration := yaml.Config{
		Rules: []yaml.ConfigRule{
			{
				PackageRegexps: []string{"PKG_1", "PKG_2"},
				NameRegexp:     `^test.(\w*)Client$`,
				Component: yaml.ConfigRuleComponent{
					Name:        "test.Client{0}",
					Description: "Client description",
					Technology:  "Client technology",
					Tags:        []string{"TAG_1", "TAG_2"},
				},
			},
		},
	}

	yamlRule := yamlConfiguration.Rules[0]

	rules, err := toScraperRules(yamlConfiguration)
	require.NoError(t, err)
	require.Len(t, rules, len(yamlConfiguration.Rules))

	r := rules[0]

	expectedRule, err := NewRule().
		WithPkgRegexps(yamlRule.PackageRegexps...).
		WithNameRegexp(yamlRule.NameRegexp).
		WithApplyFunc(func(name string, groups ...string) model.Info {
			n := fmt.Sprintf("test.Client%s", groups[0])
			return model.ComponentInfo(
				n,
				yamlRule.Component.Description,
				yamlRule.Component.Technology,
				yamlRule.Component.Tags[0],
				yamlRule.Component.Tags[1],
			)
		}).
		Build()
	require.NoError(t, err)

	pkg := "PKG_1"

	name := "test.TestClient"
	require.Equal(t, expectedRule.Applies(pkg, name), r.Applies(pkg, name))
	require.Equal(t, expectedRule.Apply(name), r.Apply(name))

	name = "test.MockClient"
	require.Equal(t, expectedRule.Applies(pkg, name), r.Applies(pkg, name))
	require.Equal(t, expectedRule.Apply(name), r.Apply(name))
}
