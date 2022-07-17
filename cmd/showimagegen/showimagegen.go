/**
URY Show Image Generator 2

Author: Michael Grace <michael.grace@ury.org.uk>
*/

package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	"github.com/UniversityRadioYork/myradio-go"
	myrsess "github.com/UniversityRadioYork/showimagegen-2/internal/myradio"
	"github.com/UniversityRadioYork/showimagegen-2/pkg/config"
	"github.com/UniversityRadioYork/showimagegen-2/pkg/generator"
	"github.com/UniversityRadioYork/showimagegen-2/pkg/logging"
)

func showImageGenerator() {
	config, err := config.NewConfigFromYAML()

	if err != nil {
		panic(err)
	}

	myr, err := myradio.NewSessionFromKeyFile()

	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	ctx = context.WithValue(ctx, myrsess.CtxKeyTimeoutSeconds, config.RequestTimeoutSeconds)
	myrLoginEnv, err := myrsess.CreateMyRadioLoginSession(ctx)

	if err != nil {
		panic(err)
	}

	defer myrLoginEnv.Close()

	env := generator.GenerationEnvironment{
		Config:           config,
		MyRadioSession:   myr,
		SetPhotoCallback: myrLoginEnv.SetShowPhoto,
	}

	shows, err := env.GetShowsToGenerateImageFor()
	if err != nil {
		panic(err)
	}

	for _, show := range shows {
		env.GenerateImageForShow(show)
	}

}

func daemon() {

	defer func() {
		if r := recover(); r != nil {
			logging.Error(fmt.Errorf("showImageGenerator Failed: %v", r))
		}
	}()

	showImageGenerator()

}

func main() {

	daemonMode := flag.Bool("daemon", false, "Usage TODO")
	flag.Parse()

	logging.Info(fmt.Sprintf("Show Image Generator | Running as Daemon: %v", *daemonMode))

	showImageGenerator()

	if *daemonMode {
		for {
			time.Sleep(time.Hour)
			daemon()
		}
	}
}
