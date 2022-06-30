/**
URY Show Image Generator 2

Author: Michael Grace <michael.grace@ury.org.uk>
*/

package myradio

type MyRadioLoginEnvironment struct {
}

func (e *MyRadioLoginEnvironment) Close() {
	// TODO
}

func CreateMyRadioLoginEnvironment() (*MyRadioLoginEnvironment, error) {
	// TODO
	return &MyRadioLoginEnvironment{}, nil
}

func (e *MyRadioLoginEnvironment) SetShowPhoto(showID int, path string) {
	// TODO
}
