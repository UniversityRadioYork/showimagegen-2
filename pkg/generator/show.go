/**
URY Show Image Generator 2

Author: Michael Grace <michael.grace@ury.org.uk>
*/

package generator

import (
	"context"
	"log"
	"math/rand"
	"time"

	"github.com/UniversityRadioYork/myradio-go"
	im "github.com/UniversityRadioYork/showimagegen-2/pkg/images"
)

var imageGenerators []im.ImageGenerator = []im.ImageGenerator{
	im.LegacyImageGenerator{},
}

// GenerateImageForShow will call an appropriate ImageGenerator to make an image
// then call the SetPhotoCallback function to associate that image to the show.
func (e *GenerationEnvironment) GenerateImageForShow(show myradio.ShowMeta, branding string) error {
	log.Printf("%v | creating show image for %s", show.ShowID, show.Title)

	rand.Seed(time.Now().Unix())
	newImage, err := imageGenerators[rand.Intn(len(imageGenerators))].Generate(im.ShowImageData{
		Show:     show,
		Branding: branding,
	})

	if err != nil {
		log.Printf("%v | error making image: %v", show.ShowID, err)
		return err
	}

	ctx, cnl := context.WithTimeout(context.Background(), time.Duration(e.Config.RequestTimeoutSeconds)*time.Second)
	defer cnl()

	log.Printf("%v | setting show photo to newly created photo %s", show.ShowID, newImage)
	uploadedPhoto, err := e.SetPhotoCallback(ctx, show.ShowID, newImage)
	if err != nil {
		log.Printf("%v | error setting show photo: %v", show.ShowID, err)
		return err
	}

	log.Printf("%v | adding the new photo (%s:%s) to the showimagegen persistence", show.ShowID, newImage, uploadedPhoto)
	e.AddToPersistence(show.ShowID, show.Title, newImage, uploadedPhoto)

	return nil
}
