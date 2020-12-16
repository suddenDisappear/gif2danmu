package util

import (
	"image"
	"image/png"
	"os"
	"strconv"
)

// New2Dimensions return two dimensions array with default value
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

// SavePngImage save png image to specific path.
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

// SaveFile save string contents to specific path.
func SaveFile(contents string, path string) error {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(contents)
	return err
}

// IntSliceToString concat int to string with separator.
func IntSliceToString(origin []int, separator string) string {
	res := ""
	for _, v := range origin {
		res += strconv.Itoa(v) + separator
	}
	return res
}
