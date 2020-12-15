package transform

import (
	"os"
	"strings"
)

type ColorMap map[string]*ColorInfo

func (c ColorMap) Save(path string) error {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()
	for c, info := range c {
		var txt = c + NewLine
		for i := 0; i < len(info.Contents); i++ {
			txt += strings.Join(info.Contents[i], "") + NewLine
		}
		_, err := f.Write([]byte(strings.TrimRight(txt, DefaultFill+NewLine) + NewLine))
		if err != nil {
			return err
		}
	}
	return nil
}
