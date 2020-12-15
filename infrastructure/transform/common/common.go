package common

import (
	"fmt"
	"gif2danmu/infrastructure/customize_error"
	"gif2danmu/infrastructure/transform"
	"gif2danmu/infrastructure/util"
	"image"
	"image/color"
	"os"
	"strconv"
)

type Image struct {
	origin   image.Image
	internal bool
}

func Open(file string) (transform.Transformer, error) {
	// 打开文件
	f, err := os.OpenFile(file, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil, customize_error.New(err, "文件打开失败")
	}
	// 解码图片
	i, _, err := image.Decode(f)
	if err != nil {
		return nil, customize_error.New(err, "图片文件解码失败，请检查图片格式")
	}
	return &Image{origin: i}, nil
}

func OpenInternal(image image.Image) (transform.Transformer, error) {
	return &Image{origin: image, internal: true}, nil
}

func Recovery(colorMap transform.ColorMap) *Image {
	// 初始化画布
	var recovery = new(image.RGBA)
	for _, v := range colorMap {
		if len(v.Contents) == 0 || len(v.Contents[0]) == 0 {
			return nil
		}
		recovery = image.NewRGBA(image.Rect(0, 0, len(v.Contents[0]), len(v.Contents)))
		break
	}
	// 图像还原
	for c, info := range colorMap {
		for i := 0; i < len(info.Contents); i++ {
			for j := 0; j < len(info.Contents[i]); j++ {
				if info.Contents[i][j] != transform.PixelSymbol {
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
	return &Image{origin: recovery}
}

func (i *Image) Transform() (*transform.ColorMap, error) {
	// 初始化空格数组
	bounds := i.origin.Bounds()
	init := util.New2Dimensions(bounds.Max.X, bounds.Max.Y, transform.DefaultFill)
	// 颜色转方块
	colorMap := make(transform.ColorMap)
	var ignorePixelCount int64 = 0
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			// 忽略特定像素
			r, g, b, a := i.origin.At(x, y).RGBA()
			if i.shouldIgnore(r, g, b, a) {
				ignorePixelCount++
				continue
			}
			// 根据rgba分组
			key := fmt.Sprintf("#%s%s%s%s", strconv.FormatUint(uint64(r>>8), 16), strconv.FormatUint(uint64(g>>8), 16), strconv.FormatUint(uint64(b>>8), 16), strconv.FormatUint(uint64(a>>8), 16))
			if _, ok := colorMap[key]; !ok {
				colorMap[key] = &transform.ColorInfo{Contents: make([][]string, len(init))}
				for k, v := range init {
					colorMap[key].Contents[k] = make([]string, len(v))
					copy(colorMap[key].Contents[k], v)
				}
			}
			colorMap[key].Contents[y][x] = transform.PixelSymbol
			colorMap[key].PixelCount = colorMap[key].PixelCount + 1
		}
	}
	return &colorMap, nil
}

func (i *Image) GetOrigin() *image.Image {
	return &i.origin
}

func (i *Image) shouldIgnore(r, g, b, a uint32) bool {
	var max uint32 = 1<<16 - 1
	if r == max && g == max && b == max {
		return true
	}
	if a == 0 {
		return true
	}
	return false
}
