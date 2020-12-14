package common

import (
	"fmt"
	"gif2danmu/infrastructure/customize_error"
	"gif2danmu/infrastructure/library"
	"gif2danmu/infrastructure/resolver"
	"image"
	"os"
	"strconv"
)

type Image struct {
	origin   image.Image
	internal bool
}

func Open(file string) (library.Transformer, error) {
	// 打开文件
	f, err := os.OpenFile(file, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil, customize_error.New(err, "文件打开失败")
	}
	// 解码图片
	i, _, err := image.Decode(f)
	if err != nil {
		return nil, customize_error.New(err, "gif文件解码失败，请检查图片格式")
	}
	return &Image{origin: i}, nil
}

func openInternal(image image.Image) (library.Transformer, error) {
	return &Image{origin: image, internal: true}, nil
}

func (i *Image) Transform() (resolver.Resolver, error) {
	// 初始化空格数组
	b := i.origin.Bounds()
	var init = make([][]string, b.Max.Y)
	for y := 0; y < b.Max.Y; y++ {
		var t = make([]string, b.Max.X)
		for x := 0; x < b.Max.X; x++ {
			t[x] = " "
		}
		init[y] = t
	}
	// 颜色图标
	pixelMap := make(map[string][][]string)
	for y := 0; y < b.Max.Y; y++ {
		for x := 0; x < b.Max.X; x++ {
			r, g, b, a := b.At(x, y).RGBA()
			key := fmt.Sprintf("#%s%s%s%s", strconv.FormatUint(uint64(r), 16), strconv.FormatUint(uint64(g), 16), strconv.FormatUint(uint64(b), 16), strconv.FormatUint(uint64(a), 16))
			if _, ok := pixelMap[key]; !ok {
				pixelMap[key] = init
			}
			pixelMap[key][y][x] = "■"
		}
	}
	// TODO:使用不同resolver处理
	return nil, nil
}
