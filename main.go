package main

import (
	"flag"
	"fmt"
	"gif2danmu/infrastructure/customize_error"
	"gif2danmu/infrastructure/flags"
	"gif2danmu/infrastructure/transform"
	"gif2danmu/infrastructure/transform/gif_image"
)

func main() {
	// 解析命令行参数
	fl := flags.RegisterFlags()
	flag.Parse()
	// 验证数据
	err := fl.Validate()
	if err != nil {
		fmt.Println(err)
		return
	}
	// 初始化
	transform.InitConfig(*fl)
	customize_error.SetDebug(fl.Debug)
	// 转换图片
	g, err := gif_image.Open(fl.File)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = g.Transform()
	if err != nil {
		fmt.Println(err)
		return
	}
}
