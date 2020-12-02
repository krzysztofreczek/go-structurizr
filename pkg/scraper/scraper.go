package scraper

import (
	"reflect"

	"github.com/krzysztofreczek/go-structurizr/pkg/model"
	"github.com/krzysztofreczek/go-structurizr/pkg/yaml"
	"github.com/pkg/errors"
)

// Scraper
type Scraper interface {
	Scrap(i interface{}) model.Structure
	RegisterRule(r Rule) error
}

type scraper struct {
	config    Configuration
	rules     []Rule
	structure model.Structure
}

// NewScraper instantiates a default Scraper implementation with provided Configuration.
func NewScraper(config Configuration) Scraper {
	return &scraper{
		config:    config,
		rules:     make([]Rule, 0),
		structure: model.NewStructure(),
	}
}

// NewScraperFromConfigFile
func NewScraperFromConfigFile(fileName string) (Scraper, error) {
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

	return &scraper{
		config:    config,
		rules:     rules,
		structure: model.NewStructure(),
	}, nil
}

func (s *scraper) RegisterRule(r Rule) error {
	if r == nil {
		return errors.New("rule must not be nil")
	}
	s.rules = append(s.rules, r)
	return nil
}

func (s *scraper) Scrap(i interface{}) model.Structure {
	v := reflect.ValueOf(i)
	s.scrap(v, "", 0)
	return s.structure
}
