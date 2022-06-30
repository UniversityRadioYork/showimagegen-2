/**
URY Show Image Generator 2

Author: Michael Grace <michael.grace@ury.org.uk>
*/

package images

type ImageInfo struct {
	Title       string
	BrandHandle string
	ShowSubtype string
}

type ImageGenerator interface {
	Generate(show ImageInfo) string
}
