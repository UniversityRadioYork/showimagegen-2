/**
URY Show Image Generator 2

Author: Michael Grace <michael.grace@ury.org.uk>
*/

package config

import (
	"fmt"
	"runtime/debug"
)

// GetCommitHash will return the commit hash of the code in the
// build. If it fails, ensure the package was built rather than
// the file, i.e.
// go build ./cmd/showimagegen
// rather than
// go build cmd/showimagegen/showimagegen.go
func GetCommitHash() (string, error) {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return "", fmt.Errorf("can't get build info")
	}

	for _, setting := range info.Settings {
		if setting.Key == "vcs.revision" {
			return setting.Value, nil
		}
	}

	return "", fmt.Errorf("can't find vcs.revision")

}
