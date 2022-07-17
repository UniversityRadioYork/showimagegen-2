/**
URY Show Image Generator 2

Author: Michael Grace <michael.grace@ury.org.uk>
*/

package images

import "context"

type CtxKey string

const (
	CtxShowKey        CtxKey = "show"
	CtxBrandHandleKey CtxKey = "brandHandle"
)

type ImageInfo struct {
	Title       string
	BrandHandle string
	ShowSubtype string
}

type ImageGenerator interface {
	Generate(ctx context.Context) (string, error)
}
