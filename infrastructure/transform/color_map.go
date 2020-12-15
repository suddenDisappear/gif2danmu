package transform

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

type ColorMap map[string]*ColorInfo

func (c ColorMap) Save(path string) error {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()
	// 总有效像素数
	var total int64
	for _, info := range c {
		total += info.PixelCount
	}
	for c, info := range c {
		if info.ShouldIgnore(total) {
			continue
		}
		var txt = c + NewLine
		for i := 0; i < len(info.Contents); i++ {
			txt += strings.TrimRight(strings.Join(info.Contents[i], ""), DefaultFill) + NewLine
		}
		if utf8.RuneCountInString(txt) > 5000 {
			return errors.New(fmt.Sprintf("最大像素数不能超过5000，当前帧像素:%d", utf8.RuneCountInString(txt)))
		}
		_, err := f.Write([]byte(strings.TrimRight(txt, DefaultFill+NewLine) + NewLine))
		if err != nil {
			return err
		}
	}
	return nil
}
