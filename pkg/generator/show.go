/**
URY Show Image Generator 2

Author: Michael Grace <michael.grace@ury.org.uk>
*/

package generator

import (
	"github.com/UniversityRadioYork/myradio-go"
	"github.com/UniversityRadioYork/showimagegen-2/pkg/images"
)

var imageGenerators []images.ImageGenerator = []images.ImageGenerator{
	images.LegacyImageGenerator{},
}

func (e *GenerationEnvironment) GenerateImageForShow(show myradio.Show) {
	imageInfo := images.ImageInfo{
		Title:       show.Title,
		ShowSubtype: "TODO",
		BrandHandle: e.Config.Branding,
	}

	// TODO: make random choice
	// TODO: goroutine it
	newImage, err := imageGenerators[0].Generate(imageInfo)

	if err != nil {
		// TODO
	}

	// TODO
	e.MyRadioLoginEnvironment.SetShowPhoto(int(show.Id), newImage)

}
