/**
URY Show Image Generator 2

Author: Michael Grace <michael.grace@ury.org.uk>
*/

package generator

import "github.com/UniversityRadioYork/myradio-go"

func (e *GenerationEnvironment) GetShowsToGenerateImageFor() ([]myradio.ShowMeta, error) {
	seasons, err := e.MyRadioSession.GetAllSeasonsInLatestTerm()
	if err != nil {
		// TODO
	}

	var shows []myradio.ShowMeta

	for _, season := range seasons {
		if season.ShowMeta.Photo == "" {
			shows = append(shows, season.ShowMeta)
		}
	}

	return shows, nil
}
