/**
URY Show Image Generator 2

Author: Michael Grace <michael.grace@ury.org.uk>
*/

package images

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"log"
	"time"

	"image/png"
	"io/fs"
	"os"
	"path/filepath"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

// LegacyImageGenerator is used to generate images following the
// same style as the old image generator.
type LegacyImageGenerator struct {
}

// Generate takes the context containing the show info and creates the image,
// returning the path to it.
func (g LegacyImageGenerator) Generate(data ShowImageData) (string, error) {
	log.Printf("%v | using legacy image generator", data.Show.ShowID)

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
		return "", err
	}

	// TODO Pick Random
	backgroundPath := backgrounds[0]
	background, err := os.Open(backgroundPath)
	if err != nil {
		return "", err
	}
	defer background.Close()

	backgroundPng, err := png.Decode(background)

	if err != nil {
		return "", err
	}

	// TODO Pick Appropriately
	subtypeColourBarPath := "assets/subtype_colour_bars/primetime.png"
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

	// TODO - font, font size, actual location
	// splitting lines, branding
	textDrawer := font.Drawer{
		Dst:  showImage,
		Src:  image.NewUniform(color.RGBA{255, 255, 255, 255}),
		Face: basicfont.Face7x13,
		Dot: fixed.Point26_6{
			X: fixed.I(backgroundPng.Bounds().Dx() / 2),
			Y: fixed.I(backgroundPng.Bounds().Dy() / 2),
		},
	}

	textDrawer.DrawString(data.Show.Title)

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
