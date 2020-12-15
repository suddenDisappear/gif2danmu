package gif_image

import (
	"fmt"
	"gif2danmu/infrastructure/customize_error"
	"gif2danmu/infrastructure/resolver"
	"gif2danmu/infrastructure/transform"
	"gif2danmu/infrastructure/transform/common"
	"image/gif"
	"os"
)

type Image struct {
	origin *gif.GIF
}

func Open(file string) (transform.Transformer, error) {
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
	// 图片像素数限制
	if g.Config.Height*g.Config.Width > 150000 {
		return nil, customize_error.New(err, "图片过大，请调整大小后重试")
	}
	return &Image{origin: g}, nil
}

func (i *Image) Transform() (resolver.Resolver, error) {
	fmt.Printf("%v", i.origin.Delay)
	for index, frame := range i.origin.Image {
		c, err := common.OpenInternal(frame)
		if err != nil {
			return nil, customize_error.New(err, fmt.Sprintf("打开第%d帧失败", index))
		}
		// TODO:处理每帧并输出
		_, err = c.Transform()
		if err != nil {
			return nil, customize_error.New(err, fmt.Sprintf("转换第%d帧失败", index))
		}
	}
	return nil, nil
}
