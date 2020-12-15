package main

import (
	"flag"
	"gif2danmu/infrastructure/transform/gif_image"
)

func main() {
	flags := RegisterFlags()
	flag.Parse()
	// 转换图片
	g, err := gif_image.Open(flags.File)
	if err != nil {
		panic(err)
	}
	_, _ = g.Transform()
}

type Flags struct {
	File string
}

func RegisterFlags() *Flags {
	var flags Flags
	flag.StringVar(&flags.File, "file", "", "文件地址")
	return &flags
}
