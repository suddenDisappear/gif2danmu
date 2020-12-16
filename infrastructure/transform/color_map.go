package transform

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

const newLine = "\n"

type ColorMap map[string]*ColorInfo

// Save 保存信息到文本文件.
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
		var txt = c + newLine
		for i := 0; i < len(info.Contents); i++ {
			txt += strings.TrimRight(strings.Join(info.Contents[i], ""), conf.Fill) + newLine
		}
		if utf8.RuneCountInString(txt) > int(conf.PixelLimit) {
			return errors.New(fmt.Sprintf("最大像素数不能超过%d，当前帧像素:%d", conf.PixelLimit, utf8.RuneCountInString(txt)))
		}
		_, err := f.Write([]byte(strings.TrimRight(txt, conf.Fill+newLine) + newLine))
		if err != nil {
			return err
		}
	}
	return nil
}

// Recovery 从信息中恢复图像.
func (c ColorMap) Recover() image.Image {
	// 初始化画布
	var recovery = new(image.RGBA)
	for _, v := range c {
		if len(v.Contents) == 0 || len(v.Contents[0]) == 0 {
			return nil
		}
		recovery = image.NewRGBA(image.Rect(0, 0, len(v.Contents[0]), len(v.Contents)))
		break
	}
	// 总有效像素数
	var total int64
	for _, info := range c {
		total += info.PixelCount
	}
	// 图像还原
	for c, info := range c {
		if info.ShouldIgnore(total) {
			continue
		}
		for i := 0; i < len(info.Contents); i++ {
			for j := 0; j < len(info.Contents[i]); j++ {
				if info.Contents[i][j] != conf.PixelSymbol {
					continue
				}
				// 颜色还原:rgba
				r, _ := strconv.ParseUint(c[1:3], 16, 8)
				g, _ := strconv.ParseUint(c[3:5], 16, 8)
				b, _ := strconv.ParseUint(c[5:7], 16, 8)
				a, _ := strconv.ParseUint(c[7:], 16, 8)
				recovery.Set(j, i, color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: uint8(a)})
			}
		}
	}
	return recovery
}
