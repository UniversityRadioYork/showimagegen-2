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

func (e *GenerationEnvironment) GenerateImageForShow(show myradio.ShowMeta) {
	ctx := context.Background()

	ctx = context.WithValue(ctx, im.CtxShowKey, show)
	ctx = context.WithValue(ctx, im.CtxBrandHandleKey, e.Config.Branding)

	// TODO: make random choice
	// TODO: goroutine it
	newImage, err := imageGenerators[0].Generate(ctx)

	if err != nil {
		// TODO
	}

	// TODO
	var cnl context.CancelFunc
	ctx, cnl = context.WithTimeout(ctx, 5*time.Second)
	defer cnl()
	if err := e.SetPhotoCallback(ctx, newImage); err != nil {
		// TODO
	}

}
