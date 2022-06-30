/**
URY Show Image Generator 2

Author: Michael Grace <michael.grace@ury.org.uk>
*/

package generator

import (
	"github.com/UniversityRadioYork/myradio-go"
	myrle "github.com/UniversityRadioYork/showimagegen-2/internal/myradio"
	"github.com/UniversityRadioYork/showimagegen-2/pkg/config"
)

type GenerationEnvironment struct {
	Config                  config.Config
	MyRadioSession          *myradio.Session
	MyRadioLoginEnvironment *myrle.MyRadioLoginEnvironment
}
