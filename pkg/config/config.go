/**
URY Show Image Generator 2

Author: Michael Grace <michael.grace@ury.org.uk>
*/

package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

// Config is a collection of all options that can be set in configuration.
type Config struct {
	Branding              string
	RequestTimeoutSeconds int
	MyRadioUsername       string
	MyRadioPassword       string
}

// NewConfigFromYAML will read a YAML file and fill in and return a Config struct.
func NewConfigFromYAML(path string) (*Config, error) {
	var config Config

	configFile, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	if err = yaml.Unmarshal(configFile, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
