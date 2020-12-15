package transform

import "image"

const (
	DefaultFill = " "
	PixelSymbol = "■"
	NewLine     = "\n"
)

type Transformer interface {
	GetOrigin() *image.Image
	Transform() (*ColorMap, error)
}
