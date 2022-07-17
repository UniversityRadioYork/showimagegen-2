/**
URY Show Image Generator 2

Author: Michael Grace <michael.grace@ury.org.uk>
*/

package main

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/UniversityRadioYork/myradio-go"
	myrsess "github.com/UniversityRadioYork/showimagegen-2/internal/myradio"
	"github.com/UniversityRadioYork/showimagegen-2/pkg/config"
	"github.com/UniversityRadioYork/showimagegen-2/pkg/generator"
)

func usage() {
	fmt.Println(`usage: showimagecli {generate|recover} SHOW_ID

	generate: regenerate a show image for this show id and upload it
	recover: TBD. i kinda know what I want tho`)
	os.Exit(1)
}

func main() {
	args := os.Args[1:]

	if len(args) != 2 {
		usage()
	}

	showID, err := strconv.Atoi(args[1])
	if err != nil {
		usage()
	}

	switch args[0] {
	case "generate":
		config, err := config.NewConfigFromYAML()
		if err != nil {
			// TODO
		}

		myr, err := myradio.NewSessionFromKeyFile()
		if err != nil {
			// TODO
		}

		myRadioLoginEnvironment, err := myrsess.CreateMyRadioLoginSession(context.TODO())
		if err != nil {
			// TODO
		}
		defer myRadioLoginEnvironment.Close()

		env := generator.GenerationEnvironment{
			Config:           config,
			MyRadioSession:   myr,
			SetPhotoCallback: myRadioLoginEnvironment.SetShowPhoto,
		}

		show, err := env.MyRadioSession.GetShow(showID)
		if err != nil {
			// TODO
		}

		env.GenerateImageForShow(*show)
	case "recover":
		// TODO
	default:
		usage()
	}
}
