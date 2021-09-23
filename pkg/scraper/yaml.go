package scraper

import (
	"fmt"
	"strings"

	"github.com/krzysztofreczek/go-structurizr/pkg/model"
	"github.com/krzysztofreczek/go-structurizr/pkg/yaml"
)

func toScraperConfig(c yaml.Config) Configuration {
	return NewConfiguration(c.Configuration.Packages...)
}

func toScraperRules(c yaml.Config) ([]Rule, error) {
	rules := make([]Rule, len(c.Rules))
	for i, r := range c.Rules {
		r := r
		rule, err := NewRule().
			WithNameRegexp(r.NameRegexp).
			WithPkgRegexps(r.PackageRegexps...).
			WithApplyFunc(
				func(name string, groups ...string) model.Info {
					info := make([]string, len(r.Component.Tags)+3)

					if r.Component.Name != "" {
						name = r.Component.Name
					}

					for i, g := range groups {
						placeholder := fmt.Sprintf("{%d}", i)
						name = strings.Replace(name, placeholder, g, -1)
					}
					info[0] = name

					info[1] = r.Component.Description
					info[2] = r.Component.Technology

					idx := 3
					for _, t := range r.Component.Tags {
						info[idx] = t
						idx++
					}

					return model.ComponentInfo(info...)
				},
			).Build()
		if err != nil {
			return nil, err
		}
		rules[i] = rule
	}
	return rules, nil
}
