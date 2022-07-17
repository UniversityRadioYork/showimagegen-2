/**
URY Show Image Generator 2

Author: Michael Grace <michael.grace@ury.org.uk>
*/

package images

import (
	"context"
	"fmt"
	"image"
	"image/color"
	"image/draw"

	"image/png"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/UniversityRadioYork/myradio-go"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

type LegacyImageGenerator struct {
}

func (g LegacyImageGenerator) Generate(ctx context.Context) (string, error) {
	show, ok := ctx.Value(CtxShowKey).(myradio.ShowMeta)
	if !ok {
		return "", fmt.Errorf("%s in ctx is not myradio ShowMeta", CtxShowKey)
	}

	var backgrounds []string
	filepath.WalkDir("assets/bw_backgrounds", func(path string, stat fs.DirEntry, err error) error {
		if err != nil {
			// TODO
			panic(err)
		}
		if !stat.IsDir() {
			backgrounds = append(backgrounds, path)
		}
		return nil
	})

	// TODO Pick Random
	backgroundPath := backgrounds[0]
	background, err := os.Open(backgroundPath)
	if err != nil {
		// TODO
		panic(err)
	}
	defer background.Close()

	backgroundPng, err := png.Decode(background)

	if err != nil {
		// TODO
		panic(err)
	}

	// TODO Pick Appropriately
	subtypeColourBarPath := "assets/subtype_colour_bars/Flagship.png"
	subtypeColourBar, err := os.Open(subtypeColourBarPath)
	if err != nil {
		// TODO
		panic(err)
	}
	defer subtypeColourBar.Close()

	colourBarPng, err := png.Decode(subtypeColourBar)
	if err != nil {
		// TODO
		panic(err)
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

	textDrawer.DrawString(show.Title)

	// TODO Show ID
	outFile, err := os.Create("test.png")
	if err != nil {
		// TODO
		panic(err)
	}
	defer outFile.Close()

	png.Encode(outFile, showImage)

	return "TODO", nil
}
