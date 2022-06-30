/**
URY Show Image Generator 2

Author: Michael Grace <michael.grace@ury.org.uk>
*/

package main

import (
	"flag"
	"time"

	"github.com/UniversityRadioYork/myradio-go"
	myrle "github.com/UniversityRadioYork/showimagegen-2/internal/myradio"
	"github.com/UniversityRadioYork/showimagegen-2/pkg/config"
	"github.com/UniversityRadioYork/showimagegen-2/pkg/generator"
)

func showImageGenerator() {
	config, err := config.NewConfigFromYAML()

	if err != nil {
		// TODO
	}

	myr, err := myradio.NewSessionFromKeyFile()

	if err != nil {
		// TODO
	}

	myrLoginEnv, err := myrle.CreateMyRadioLoginEnvironment()

	if err != nil {
		// TODO
	}

	defer myrLoginEnv.Close()

	env := generator.GenerationEnvironment{
		Config:                  config,
		MyRadioSession:          myr,
		MyRadioLoginEnvironment: myrLoginEnv,
	}

	for _, show := range env.GetShowsToGenerateImageFor() {
		env.GenerateImageForShow(show)
	}

}

func main() {

	daemonMode := flag.Bool("daemon", false, "Usage TODO")
	flag.Parse()

	showImageGenerator()

	if *daemonMode {
		for {
			time.Sleep(time.Hour)
			showImageGenerator()
		}
	}
}
