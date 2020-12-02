package scraper

import (
	"github.com/krzysztofreczek/go-structurizr/pkg/model"
	"github.com/krzysztofreczek/go-structurizr/pkg/yaml"
	"github.com/pkg/errors"
	"reflect"
)

type Scraper struct {
	config    Configuration
	rules     []Rule
	structure model.Structure
}

func NewScraper(config Configuration) *Scraper {
	return &Scraper{
		config:    config,
		rules:     make([]Rule, 0),
		structure: model.NewStructure(),
	}
}

func NewScraperFromConfigFile(fileName string) (*Scraper, error) {
	configuration, err := yaml.LoadFromFile(fileName)
	if err != nil {
		return nil, errors.Wrapf(err,
			"could not load configuration from file `%s`", fileName)
	}

	config := toScraperConfig(configuration)
	rules, err := toScraperRules(configuration)
	if err != nil {
		return nil, errors.Wrapf(err,
			"could not load scraper rules from file `%s`", fileName)
	}

	return &Scraper{
		config:    config,
		rules:     rules,
		structure: model.NewStructure(),
	}, nil
}

func (s *Scraper) RegisterRule(r Rule) error {
	if r == nil {
		return errors.New("rule must not be nil")
	}
	s.rules = append(s.rules, r)
	return nil
}

func (s *Scraper) Scrap(i interface{}) model.Structure {
	v := reflect.ValueOf(i)
	s.scrap(v, "", 0)
	return s.structure
}
