/**
URY Show Image Generator 2

Author: Michael Grace <michael.grace@ury.org.uk>
*/

package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/UniversityRadioYork/myradio-go"
	myrsess "github.com/UniversityRadioYork/showimagegen-2/internal/myradio"
	"github.com/UniversityRadioYork/showimagegen-2/pkg/config"
	"github.com/UniversityRadioYork/showimagegen-2/pkg/generator"
	"github.com/UniversityRadioYork/showimagegen-2/pkg/logging"
	"github.com/UniversityRadioYork/showimagegen-2/pkg/persistence"
)

func showImageGenerator() {
	config, err := config.NewConfigFromYAML("config.yaml")

	if err != nil {
		panic(err)
	}

	myr, err := myradio.NewSessionFromKeyFile()

	if err != nil {
		panic(err)
	}

	myrLoginSess, err := myrsess.CreateMyRadioLoginSession(config.MyRadioUsername, config.MyRadioPassword, config.RequestTimeoutSeconds, myr)

	if err != nil {
		panic(err)
	}

	defer myrLoginSess.Close()

	persistenceEngine, err := persistence.CreatePersistenceEngine()
	if err != nil {
		panic(err)
	}

	defer persistenceEngine.Close()

	go persistenceEngine.Daemon()

	env := generator.GenerationEnvironment{
		Config:           config,
		MyRadioSession:   myr,
		SetPhotoCallback: myrLoginSess.SetShowPhoto,
		AddToPersistence: persistenceEngine.AddEntry,
	}

	shows, err := env.GetShowsToGenerateImageFor()
	if err != nil {
		panic(err)
	}

	for _, show := range shows {
		env.GenerateImageForShow(show, config.Branding)
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
	versionFlag := flag.Bool("version", false, "display version")
	flag.Parse()

	if *versionFlag {
		commit, err := config.GetCommitHash()
		if err != nil {
			panic(err)
		}

		fmt.Printf(`URY Show Image Generator 2 - Electric Photoloo
	
Commit: %v
`, commit)
		return
	}

	logging.Info(fmt.Sprintf("Show Image Generator | Running as Daemon: %v", *daemonMode))

	showImageGenerator()

	if *daemonMode {
		for {
			time.Sleep(time.Hour)
			daemon()
		}
	}
}
