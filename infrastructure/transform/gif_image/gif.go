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

// Open 加载指定文件.
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

// Transform 转换图片并输出到指定文件夹.
func (i *Image) Transform() error {
	// 初始化文件夹
	separator := string(filepath.Separator)
	dir := transform.GetConfig().OutputDir + separator + i.taskName
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return customize_error.New(err, fmt.Sprintf("初始化文件夹%s失败", dir))
	}
	// 保存每帧图片转换字符和还原png图像
	for index, frame := range i.origin.Image {
		c, err := common.OpenInternal(frame)
		if err != nil {
			return customize_error.New(err, fmt.Sprintf("打开第%d帧失败", index))
		}
		colorMap, err := c.Transform()
		if err != nil {
			return customize_error.New(err, fmt.Sprintf("转换第%d帧失败", index))
		}
		// 保存转换后文本信息
		indexStr := strconv.Itoa(index)
		err = colorMap.Save(dir + separator + indexStr + ".txt")
		if err != nil {
			return customize_error.New(err, fmt.Sprintf("保存第%d帧文本信息失败", index))
		}
		// 保存还原后图片
		restoredImage := colorMap.Recover()
		if restoredImage == nil {
			return fmt.Errorf("恢复第%d帧图像信息错误", index)
		}
		err = util.SavePngImage(&restoredImage, dir+separator+indexStr+".png")
		if err != nil {
			return customize_error.New(err, fmt.Sprintf("保存第%d帧png图像失败", index))
		}
	}
	// 保存帧间隔信息
	err = util.SaveFile(util.IntSliceToString(i.origin.Delay, "\n"), dir+string(filepath.Separator)+"delay.txt")
	if err != nil {
		return customize_error.New(err, "保存帧延时信息失败")
	}
	return nil
}
