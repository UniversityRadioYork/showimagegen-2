/**
URY Show Image Generator 2

Author: Michael Grace <michael.grace@ury.org.uk>
*/

package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/UniversityRadioYork/myradio-go"
	myrsess "github.com/UniversityRadioYork/showimagegen-2/internal/myradio"
	"github.com/UniversityRadioYork/showimagegen-2/pkg/config"
	"github.com/UniversityRadioYork/showimagegen-2/pkg/generator"
	"github.com/UniversityRadioYork/showimagegen-2/pkg/persistence"
)

func usage() {
	fmt.Println(`usage: showimagecli {generate|recover} SHOW_ID

	generate: regenerate a show image for this show id and upload it
	recover: TBD. i kinda know what I want tho`)
	os.Exit(1)
}

func main() {
	// TODO --version flag
	args := os.Args[1:]

	if len(args) != 2 && len(args) != 3 {
		usage()
	}

	showID, err := strconv.Atoi(args[1])
	if err != nil {
		usage()
	}

	switch args[0] {
	case "generate":
		if len(args) != 2 {
			usage()
		}

		config, err := config.NewConfigFromYAML("config.yaml")
		if err != nil {
			// TODO
		}

		myr, err := myradio.NewSessionFromKeyFile()
		if err != nil {
			// TODO
		}

		myRadioLoginEnvironment, err := myrsess.CreateMyRadioLoginSession(config.MyRadioUsername, config.MyRadioPassword, config.RequestTimeoutSeconds, myr)
		if err != nil {
			// TODO
		}
		defer myRadioLoginEnvironment.Close()

		persistenceEngine, err := persistence.CreatePersistenceEngine()
		if err != nil {
			// TODO
		}

		env := generator.GenerationEnvironment{
			Config:           config,
			MyRadioSession:   myr,
			SetPhotoCallback: myRadioLoginEnvironment.SetShowPhoto,
			AddToPersistence: persistenceEngine.AddEntry,
		}

		show, err := env.MyRadioSession.GetShow(showID)
		if err != nil {
			// TODO
		}

		env.GenerateImageForShow(*show, config.Branding)
	case "recover":
		if len(args) == 2 {
			persistenceEngine, err := persistence.CreatePersistenceEngine()
			if err != nil {
				// TODO
			}

			persistenceEngine.List(showID)
		}

		// TODO
	default:
		usage()
	}
}
