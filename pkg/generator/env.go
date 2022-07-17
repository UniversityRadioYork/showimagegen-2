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

// CtxKey is a `string` type for keys in the context
type CtxKey string

const (
	// CtxShowIDKey is a key for use in a context that will contain the show ID
	CtxShowIDKey CtxKey = "showID"
)

// GenerationEnvironment contains all the attributes needed for generating
// the showimages
type GenerationEnvironment struct {
	Config           config.Config
	MyRadioSession   *myradio.Session
	SetPhotoCallback func(ctx context.Context, path string) error
}
