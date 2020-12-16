package transform

import "image"

type Transformer interface {
	GetOrigin() *image.Image
	Transform() (*ColorMap, error)
}
