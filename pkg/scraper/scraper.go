package scraper

import (
	"reflect"

	"github.com/krzysztofreczek/go-structurizr/pkg/model"
	"github.com/krzysztofreczek/go-structurizr/pkg/yaml"
	"github.com/pkg/errors"
)

// Scraper represents the default responsibilities of a scraper.
//
// Scrape processes the given structure according to the internal configuration
// and registered rules. It returns an open `model.Structure` containing recognized
// components and the relationships between them.
//
// RegisterRule registers a `Rule` with the scraper. It will return an error
// if the provided rule is nil.
type Scraper interface {
	Scrape(i interface{}) model.Structure
	RegisterRule(r Rule) error
}

type scraper struct {
	config       Configuration
	rules        []Rule
	structure    model.Structure
	typeCounters map[string]int
}

// NewScraper creates a new Scraper instance using the provided Configuration.
func NewScraper(config Configuration) Scraper {
	return &scraper{
		config:       config,
		rules:        make([]Rule, 0),
		structure:    model.NewStructure(),
		typeCounters: make(map[string]int),
	}
}

// NewScraperFromConfigFile creates a new Scraper instance using Configuration
// loaded from the specified YAML configuration file.
//
// It returns an error if the YAML file does not exist or contains invalid content.
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
		config:       config,
		rules:        rules,
		structure:    model.NewStructure(),
		typeCounters: make(map[string]int),
	}, nil
}

// RegisterRule adds the specified Rule to the scraper.
//
// It returns an error if the provided rule is nil.
func (s *scraper) RegisterRule(r Rule) error {
	if r == nil {
		return errors.New("rule must not be nil")
	}
	s.rules = append(s.rules, r)
	return nil
}

// Scrape processes the given structure according to the internal configuration
// and registered rules.
//
// It returns an open `model.Structure` containing recognized components
// and their relationships.
func (s *scraper) Scrape(i interface{}) model.Structure {
	v := reflect.ValueOf(i)
	s.scrape(v, "", 0)
	return s.structure
}
