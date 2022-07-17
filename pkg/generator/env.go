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

type CtxKey string

const (
	CtxShowIDKey CtxKey = "showID"
)

type GenerationEnvironment struct {
	Config           config.Config
	MyRadioSession   *myradio.Session
	SetPhotoCallback func(ctx context.Context, path string) error
}
