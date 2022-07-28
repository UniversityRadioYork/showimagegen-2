/**
URY Show Image Generator 2

Author: Michael Grace <michael.grace@ury.org.uk>
*/

package persistence

import (
	"fmt"
	"strings"
	"time"
)

func urlToPath(url string) string {
	splits := strings.Split(url, "/")
	return splits[len(splits)-1]
}

// List out (tab separated) the persisted entries for this showID.
func (e *Engine) List(showID int) {
	state, ok := e.state[showID]
	if !ok {
		fmt.Printf("Show ID %v Not Found!\n", showID)
		return
	}

	for _, entry := range state {
		fmt.Printf("%v\t| %v\t| %v\n", entry.Title, entry.Datetime.Format(time.ANSIC), urlToPath(entry.UploadedPhoto))
	}
}
