/**
URY Show Image Generator 2

Author: Michael Grace <michael.grace@ury.org.uk>
*/

package images

import (
	"fmt"
	"image"
	"image/draw"
	"log"
	"math/rand"
	"time"

	"image/png"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

// LegacyImageGenerator is used to generate images following the
// same style as the old image generator.
type LegacyImageGenerator struct {
}

func getBackgroundFilepaths() ([]string, error) {
	var backgrounds []string
	if err := filepath.WalkDir("assets/bw_backgrounds", func(path string, stat fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !stat.IsDir() {
			backgrounds = append(backgrounds, path)
		}
		return nil
	}); err != nil {
		return backgrounds, err
	}

	return backgrounds, nil
}

func addBranding(baseImage *image.RGBA, branding string) *image.RGBA {
	// TODO multiple lines
	fontFile, err := os.ReadFile("assets/fonts/Raleway-LightItalic.ttf")
	if err != nil {
		// TODO
	}

	f, err := freetype.ParseFont(fontFile)
	if err != nil {
		// TODO
	}

	ctx := freetype.NewContext()
	ctx.SetFont(f)
	ctx.SetFontSize(48)
	ctx.SetClip(baseImage.Bounds())
	ctx.SetDst(baseImage)
	ctx.SetSrc(image.White)

	advance := font.MeasureString(
		truetype.NewFace(f, &truetype.Options{
			Size: 48,
		}),
		branding,
	)

	if _, err := ctx.DrawString(
		branding,
		freetype.Pt(
			(baseImage.Bounds().Dx()-advance.Round())/2,
			baseImage.Bounds().Dy()*2/3,
		),
	); err != nil {
		// TODO
	}

	return baseImage
}

func addShowTitleText(baseImage *image.RGBA, title string) *image.RGBA {
	// TODO - splitting into multiple lines
	fontFile, err := os.ReadFile("assets/fonts/Raleway-Bold.ttf")
	if err != nil {
		// TODO
	}

	f, err := freetype.ParseFont(fontFile)
	if err != nil {
		// TODO
	}

	ctx := freetype.NewContext()
	ctx.SetFont(f)
	ctx.SetClip(baseImage.Bounds())
	ctx.SetFontSize(72)
	ctx.SetDst(baseImage)
	ctx.SetSrc(image.White)

	advance := font.MeasureString(
		truetype.NewFace(f, &truetype.Options{
			Size: 72,
		}),
		title,
	)

	if _, err := ctx.DrawString(
		title,
		freetype.Pt(
			(baseImage.Bounds().Dx()-advance.Round())/2,
			baseImage.Bounds().Dy()/2,
		),
	); err != nil {
		// TODO
	}

	return baseImage
}

// Generate takes show info and creates the image,
// returning the path to it.
func (g LegacyImageGenerator) Generate(data ShowImageData) (string, error) {
	log.Printf("%v | using legacy image generator", data.Show.ShowID)

	// Get a background image
	backgrounds, err := getBackgroundFilepaths()
	if err != nil {
		return "", err
	}

	rand.Seed(time.Now().Unix())
	backgroundPath := backgrounds[rand.Intn(len(backgrounds))]
	background, err := os.Open(backgroundPath)
	if err != nil {
		return "", err
	}
	defer background.Close()

	backgroundPng, err := png.Decode(background)
	if err != nil {
		return "", err
	}

	// Overlay the subtype colour.
	subtypeColourBarPath := fmt.Sprintf("assets/subtype_colour_bars/%s.png", data.Show.Subtype.Class)
	subtypeColourBar, err := os.Open(subtypeColourBarPath)
	if err != nil {
		return "", err
	}
	defer subtypeColourBar.Close()

	colourBarPng, err := png.Decode(subtypeColourBar)
	if err != nil {
		return "", err
	}

	showImage := image.NewRGBA(backgroundPng.Bounds())
	draw.Draw(showImage, backgroundPng.Bounds(), backgroundPng, image.Point{}, draw.Src)
	draw.Draw(showImage, backgroundPng.Bounds(), colourBarPng, image.Point{}, draw.Over)

	// Add the text
	showImage = addShowTitleText(showImage, data.Show.Title)
	showImage = addBranding(showImage, data.Branding)

	// Save the image file
	imageFile := fmt.Sprintf("out/%v.%v.png", data.Show.ShowID, time.Now().Unix())
	outFile, err := os.Create(imageFile)
	if err != nil {
		return "", err
	}
	defer outFile.Close()

	if err := png.Encode(outFile, showImage); err != nil {
		return "", err
	}

	return imageFile, nil
}
