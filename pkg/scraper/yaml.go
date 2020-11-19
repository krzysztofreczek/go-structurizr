package scraper

import (
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
				func() model.Info {
					info := make([]string, len(r.Component.Tags)+2)
					info[0] = r.Component.Description
					info[1] = r.Component.Technology

					idx := 2
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
