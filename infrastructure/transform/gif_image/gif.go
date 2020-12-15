package gif_image

import (
	"fmt"
	"gif2danmu/infrastructure/customize_error"
	"gif2danmu/infrastructure/transform"
	"gif2danmu/infrastructure/transform/common"
	"gif2danmu/infrastructure/util"
	"image/gif"
	"os"
	"path/filepath"
	"strconv"
)

type Image struct {
	origin   *gif.GIF
	taskName string
}

func Open(path string) (*Image, error) {
	// 打开文件
	f, err := os.OpenFile(path, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil, customize_error.New(err, "文件打开失败")
	}
	// 解码图片
	g, err := gif.DecodeAll(f)
	if err != nil {
		return nil, customize_error.New(err, "gif文件解码失败，请检查图片格式")
	}
	return &Image{origin: g, taskName: filepath.Base(path)}, nil
}

func (i *Image) Transform() (*transform.ColorMap, error) {
	dir := "output" + string(filepath.Separator) + i.taskName
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return nil, customize_error.New(err, fmt.Sprintf("初始化文件夹%s失败", dir))
	}
	// 保存每帧间隔信息
	err = util.SaveFile(util.IntSliceToString(i.origin.Delay), dir+string(filepath.Separator)+"delay.txt")
	if err != nil {
		return nil, customize_error.New(err, "保存帧延时信息失败")
	}
	for index, frame := range i.origin.Image {
		c, err := common.OpenInternal(frame)
		if err != nil {
			return nil, customize_error.New(err, fmt.Sprintf("打开第%d帧失败", index))
		}
		colorMap, err := c.Transform()
		if err != nil {
			return nil, customize_error.New(err, fmt.Sprintf("转换第%d帧失败", index))
		}
		// 保存txt文件
		indexStr := strconv.Itoa(index)
		err = colorMap.Save(dir + string(filepath.Separator) + indexStr + ".txt")
		if err != nil {
			return nil, err
		}
		// 保存还原图片
		frameImage := common.Recovery(*colorMap)
		err = util.SavePngImage(frameImage.GetOrigin(), dir+string(filepath.Separator)+indexStr+".png")
		if err != nil {
			return nil, err
		}
	}
	return nil, nil
}
