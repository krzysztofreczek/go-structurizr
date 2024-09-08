package yaml

import (
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

// Config represents a YAML configuration structure.
type Config struct {
	Configuration ConfigConfiguration `yaml:"configuration"`
	Rules         []ConfigRule        `yaml:"rules"`
	View          ConfigView          `yaml:"view"`
}

// ConfigConfiguration represents a YAML configuration structure.
type ConfigConfiguration struct {
	Packages []string `yaml:"pkgs"`
}

// ConfigRule represents a YAML configuration structure for rules.
type ConfigRule struct {
	PackageRegexps []string            `yaml:"pkg_regexps"`
	NameRegexp     string              `yaml:"name_regexp"`
	Component      ConfigRuleComponent `yaml:"component"`
}

// ConfigRuleComponent represents a YAML configuration structure for rule components.
type ConfigRuleComponent struct {
	Name        string   `yaml:"name"`
	Description string   `yaml:"description"`
	Technology  string   `yaml:"technology"`
	Tags        []string `yaml:"tags"`
}

// ConfigView represents a YAML configuration structure for views.
type ConfigView struct {
	Title             string            `yaml:"title"`
	LineColor         string            `yaml:"line_color"`
	Styles            []ConfigViewStyle `yaml:"styles"`
	ComponentTags     []string          `yaml:"component_tags"`
	RootComponentTags []string          `yaml:"root_component_tags"`
}

// ConfigViewStyle represents a YAML configuration structure for view styles.
type ConfigViewStyle struct {
	ID              string `yaml:"id"`
	BackgroundColor string `yaml:"background_color"`
	FontColor       string `yaml:"font_color"`
	BorderColor     string `yaml:"border_color"`
	Shape           string `yaml:"shape"`
}

// LoadFromFile loads a Config from a YAML file.
//
// It returns an error if the file does not exist or cannot be decoded.
func LoadFromFile(fileName string) (Config, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return Config{}, err
	}
	defer func() {
		_ = f.Close()
	}()

	return LoadFrom(f)
}

// LoadFrom loads a Config from YAML content read from an io.Reader.
//
// It returns an error if the content from the io.Reader cannot be decoded.
func LoadFrom(source io.Reader) (Config, error) {
	var cfg Config
	decoder := yaml.NewDecoder(source)
	err := decoder.Decode(&cfg)
	if err != nil && err == io.EOF {
		return Config{}, nil
	}
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}
