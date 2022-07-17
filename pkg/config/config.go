/**
URY Show Image Generator 2

Author: Michael Grace <michael.grace@ury.org.uk>
*/

package config

type Config struct {
	Branding              string
	RequestTimeoutSeconds int
}

func NewConfigFromYAML() (Config, error) {
	// TODO
	return Config{
		RequestTimeoutSeconds: 5,
	}, nil
}
