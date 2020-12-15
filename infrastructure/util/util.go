package util

import (
	"image"
	"image/png"
	"os"
)

// NewString2Dimensions return two dimensions array with default value
// x is column and y is row.
func New2Dimensions(x, y int, defaultValue ...string) [][]string {
	var s = make([][]string, y)
	for i := 0; i < y; i++ {
		s[i] = make([]string, x)
		if len(defaultValue) > 0 {
			for k := range s[i] {
				s[i][k] = defaultValue[0]
			}
		}
	}
	return s
}

func SavePngImage(i *image.Image, path string) error {
	if i == nil {
		return nil
	}
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()
	return png.Encode(f, *i)
}
