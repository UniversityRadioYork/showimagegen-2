/**
URY Show Image Generator 2

Author: Michael Grace <michael.grace@ury.org.uk>
*/

package images

import (
	"github.com/UniversityRadioYork/myradio-go"
)

// ShowImageData represents what is neccesary to make any show image
type ShowImageData struct {
	Show     myradio.ShowMeta
	Branding []string
}

// ImageGenerator defines an interface all image generators must follow
type ImageGenerator interface {
	Generate(data ShowImageData) (string, error)
}
