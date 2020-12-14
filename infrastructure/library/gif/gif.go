package gif

import (
	"gif2danmu/infrastructure/customize_error"
	"gif2danmu/infrastructure/library"
	"gif2danmu/infrastructure/resolver"
	"image/gif"
	"os"
)

type Image struct {
	origin *gif.GIF
}

func Open(file string) (library.Transformer, error) {
	// 打开文件
	f, err := os.OpenFile(file, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil, customize_error.New(err, "文件打开失败")
	}
	// 解码图片
	g, err := gif.DecodeAll(f)
	if err != nil {
		return nil, customize_error.New(err, "gif文件解码失败，请检查图片格式")
	}
	return &Image{origin: g}, nil
}

func (i *Image) Transform() (resolver.Resolver, error) {
	return nil, nil
}
