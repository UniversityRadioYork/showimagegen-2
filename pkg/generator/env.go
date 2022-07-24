/**
URY Show Image Generator 2

Author: Michael Grace <michael.grace@ury.org.uk>
*/

package generator

import (
	"context"

	"github.com/UniversityRadioYork/myradio-go"
	"github.com/UniversityRadioYork/showimagegen-2/pkg/config"
)

// GenerationEnvironment contains all the attributes needed for generating
// the showimages
type GenerationEnvironment struct {
	Config           *config.Config
	MyRadioSession   *myradio.Session
	SetPhotoCallback func(ctx context.Context, showID int, path string) error
	AddToPersistence func(showID int, title, filepath string)
}
