package yaml

import (
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

// Config is an open YAML configuration structure.
type Config struct {
	Configuration ConfigConfiguration `yaml:"configuration"`
	Rules         []ConfigRule        `yaml:"rules"`
	View          ConfigView          `yaml:"view"`
}

// ConfigConfiguration is an open YAML configuration structure.
type ConfigConfiguration struct {
	Packages []string `yaml:"pkgs"`
}

// ConfigRule is an open YAML configuration structure.
type ConfigRule struct {
	PackageRegexps []string            `yaml:"pkg_regexps"`
	NameRegexp     string              `yaml:"name_regexp"`
	Component      ConfigRuleComponent `yaml:"component"`
}

// ConfigRuleComponent is an open YAML configuration structure.
type ConfigRuleComponent struct {
	Name        string   `yaml:"name"`
	Description string   `yaml:"description"`
	Technology  string   `yaml:"technology"`
	Tags        []string `yaml:"tags"`
}

// ConfigView is an open YAML configuration structure.
type ConfigView struct {
	Title             string            `yaml:"title"`
	LineColor         string            `yaml:"line_color"`
	Styles            []ConfigViewStyle `yaml:"styles"`
	ComponentTags     []string          `yaml:"component_tags"`
	RootComponentTags []string          `yaml:"root_component_tags"`
}

// ConfigViewStyle is an open YAML configuration structure.
type ConfigViewStyle struct {
	ID              string `yaml:"id"`
	BackgroundColor string `yaml:"background_color"`
	FontColor       string `yaml:"font_color"`
	BorderColor     string `yaml:"border_color"`
	Shape           string `yaml:"shape"`
}

// LoadFromFile loads Config from YAML file.
// LoadFromFile will return an error in case file does not exists
// or cannot be decoded.
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

// LoadFrom loads Config from YAML content read from io.Reader.
// LoadFrom will return an error in case io.Reader content cannot be decoded.
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
