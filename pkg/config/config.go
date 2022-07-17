/**
URY Show Image Generator 2

Author: Michael Grace <michael.grace@ury.org.uk>
*/

package config

// Config is a collection of all options that can be set in configuration.
type Config struct {
	Branding              string
	RequestTimeoutSeconds int
}

// NewConfigFromYAML will read a YAML file and fill in and return a Config struct.
func NewConfigFromYAML() (Config, error) {
	// TODO
	return Config{
		RequestTimeoutSeconds: 5,
	}, nil
}
