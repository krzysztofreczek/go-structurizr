package scraper

import (
	"reflect"

	"github.com/krzysztofreczek/go-structurizr/pkg/model"
	"github.com/krzysztofreczek/go-structurizr/pkg/yaml"
	"github.com/pkg/errors"
)

// Scraper represents default scraper responsibilities.
//
// Scrape reflects given structure in accordance with internal configuration
// and registered rules.
// Scrape returns an open model.Structure with recognised components
// and relations between those.
//
// RegisterRule registers given Rule in the scraper.
// RegisterRule will return an error in case the given rule is nil.
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

// NewScraper instantiates a default Scraper implementation
// with provided Configuration.
func NewScraper(config Configuration) Scraper {
	return &scraper{
		config:       config,
		rules:        make([]Rule, 0),
		structure:    model.NewStructure(),
		typeCounters: make(map[string]int),
	}
}

// NewScraperFromConfigFile instantiates a default Scraper implementation
// with Configuration loaded from provided YAML configuration file.
// NewScraperFromConfigFile will return an error in case the YAML configuration
// file does not exist or contains invalid content.
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

// RegisterRule registers given Rule in the scraper.
// RegisterRule will return an error in case the given rule is nil.
func (s *scraper) RegisterRule(r Rule) error {
	if r == nil {
		return errors.New("rule must not be nil")
	}
	s.rules = append(s.rules, r)
	return nil
}

// Scrape reflects given structure in accordance with internal configuration
// and registered rules.
// Scrape returns an open model.Structure with recognised components
// and relations between those.
func (s *scraper) Scrape(i interface{}) model.Structure {
	v := reflect.ValueOf(i)
	s.scrape(v, "", 0)
	return s.structure
}
