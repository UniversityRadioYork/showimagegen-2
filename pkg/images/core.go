/**
URY Show Image Generator 2

Author: Michael Grace <michael.grace@ury.org.uk>
*/

package images

import "context"

// CtxKey is a `string` type for keys in the context.
type CtxKey string

const (
	// CtxShowKey is the context key for a MyRadioShowMeta object
	CtxShowKey CtxKey = "show"

	// CtxBrandHandleKey is the key for the string of branding to put on images
	CtxBrandHandleKey CtxKey = "brandHandle"
)

// ImageGenerator defines an interface all image generators must follow
type ImageGenerator interface {
	Generate(ctx context.Context) (string, error)
}
