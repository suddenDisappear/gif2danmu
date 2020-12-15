package common

import (
	"fmt"
	"gif2danmu/infrastructure/customize_error"
	"gif2danmu/infrastructure/resolver"
	"gif2danmu/infrastructure/transform"
	"gif2danmu/infrastructure/util"
	"image"
	"image/color"
	"image/png"
	"os"
	"strconv"
	"strings"
)

const (
	defaultFill = " "
	pixelSymbol = "■"
	newLine     = "\n"
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

func (i *Image) Transform() (resolver.Resolver, error) {
	// 初始化空格数组
	bounds := i.origin.Bounds()
	init := util.New2Dimensions(bounds.Max.X, bounds.Max.Y, defaultFill)
	// 颜色转方块
	colorMap := make(map[string]*colorInfo)
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
				colorMap[key] = &colorInfo{contents: make([][]string, len(init))}
				for k, v := range init {
					colorMap[key].contents[k] = make([]string, len(v))
					copy(colorMap[key].contents[k], v)
				}
			}
			colorMap[key].contents[y][x] = pixelSymbol
			colorMap[key].pixelCount = colorMap[key].pixelCount + 1
		}
	}
	// 输出到resolver
	tmp, err := os.OpenFile("tmp.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer tmp.Close()
	recovery := image.NewRGBA(image.Rect(bounds.Min.X, bounds.Min.Y, bounds.Max.X, bounds.Max.Y))
	for c, info := range colorMap {
		// 忽视像素占比<?%颜色
		if info.IsIgnore(int64(bounds.Max.Y*bounds.Max.X) - ignorePixelCount) {
			continue
		}
		var txt = c + newLine
		for i := 0; i < len(info.contents); i++ {
			txt += strings.Join(info.contents[i], "") + newLine
			// 图像还原
			for j := 0; j < len(info.contents[i]); j++ {
				if info.contents[i][j] != pixelSymbol {
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
		txt = strings.TrimRight(txt, defaultFill)
		_, err := tmp.Write([]byte(strings.TrimRight(txt, defaultFill+newLine) + newLine))
		if err != nil {
			return nil, err
		}
	}
	// 保存还原后图像
	recoveryImage, err := os.OpenFile("tmp.png", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer recoveryImage.Close()
	err = png.Encode(recoveryImage, recovery)
	if err != nil {
		return nil, err
	}
	return nil, nil
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

type colorInfo struct {
	contents   [][]string
	pixelCount int64
}

func (c *colorInfo) IsIgnore(total int64) bool {
	return float64(c.pixelCount) < float64(total)*0.001
}
