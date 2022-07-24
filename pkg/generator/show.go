/**
URY Show Image Generator 2

Author: Michael Grace <michael.grace@ury.org.uk>
*/

package generator

import (
	"context"
	"time"

	"github.com/UniversityRadioYork/myradio-go"
	im "github.com/UniversityRadioYork/showimagegen-2/pkg/images"
)

var imageGenerators []im.ImageGenerator = []im.ImageGenerator{
	im.LegacyImageGenerator{},
}

// GenerateImageForShow will call an appropriate ImageGenerator to make an image
// then call the SetPhotoCallback function to associate that image to the show.
func (e *GenerationEnvironment) GenerateImageForShow(show myradio.ShowMeta, branding string) {
	// TODO: make random choice
	// TODO: goroutine it
	newImage, err := imageGenerators[0].Generate(im.ShowImageData{
		Show:     show,
		Branding: branding,
	})

	if err != nil {
		// TODO
	}

	// TODO
	ctx, cnl := context.WithTimeout(context.Background(), 5*time.Second)
	defer cnl()
	if err := e.SetPhotoCallback(ctx, show.ShowID, newImage); err != nil {
		// TODO
	}

}
