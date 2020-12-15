package main

import "gif2danmu/infrastructure/transform/gif_image"

func main() {
	// TODO:替换成接收命令行参数
	g, err := gif_image.Open("D:\\download\\image\\test_resize.gif")
	if err != nil {
		panic(err)
	}
	// TODO:处理实际输入输出
	_, _ = g.Transform()
}
