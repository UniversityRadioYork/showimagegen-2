/**
URY Show Image Generator 2

Author: Michael Grace <michael.grace@ury.org.uk>
*/

package generator

import (
	"log"

	"github.com/UniversityRadioYork/myradio-go"
)

// GetShowsToGenerateImageFor will get all the shows in the current term from MyRadio and
// work out which ones need a show image generating for, returning only those ones.
func (e *GenerationEnvironment) GetShowsToGenerateImageFor() ([]myradio.ShowMeta, error) {
	// TODO this is temporary
	s, err := e.MyRadioSession.GetShow(13933)
	if err != nil {
		return nil, err
	}

	log.Printf("getting seasons in current term")
	seasons, err := e.MyRadioSession.GetAllSeasonsInLatestTerm()
	if err != nil {
		return nil, err
	}

	var shows []myradio.ShowMeta = []myradio.ShowMeta{*s}

	for _, season := range seasons {
		if season.ShowMeta.Photo == "" {
			log.Printf("%s (ID %v) needs an image generating for\n", season.ShowMeta.Title, season.ShowID)
			shows = append(shows, season.ShowMeta)
		}
	}

	log.Printf("found %v shows that need an image", len(shows))

	return shows, nil
}
