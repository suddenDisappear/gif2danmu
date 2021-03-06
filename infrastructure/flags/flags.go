package flags

import (
	"errors"
	"flag"
	"unicode/utf8"
)

const (
	defaultFill        = " "
	defaultPixelSymbol = "■"
	defaultOutputDir   = "output"
	acFunPixelLimit    = 5000
)

type Flags struct {
	File                string
	OutputDir           string
	Fill                string
	PixelSymbol         string
	PixelLimit          uint64
	PixelCountThreshold float64
	Debug               bool
}

func (f *Flags) Validate() error {
	if utf8.RuneCount([]byte(f.Fill)) > 1 {
		return errors.New("像素占位符不可超过1个utf8字符")
	}
	if utf8.RuneCount([]byte(f.PixelSymbol)) != 1 {
		return errors.New("像素标记仅可为1个utf8字符")
	}
	return nil
}

// RegisterFlags 初始化命令行待解析参数.
func RegisterFlags() *Flags {
	var flags Flags
	flag.StringVar(&flags.File, "file", "", "文件地址")
	flag.StringVar(&flags.OutputDir, "output", defaultOutputDir, "输出文件夹")
	flag.StringVar(&flags.Fill, "fill", defaultFill, "填空元素，用于占位，不可为#")
	flag.StringVar(&flags.PixelSymbol, "pixel_symbol", defaultPixelSymbol, "像素标记，不可为#")
	flag.Uint64Var(&flags.PixelLimit, "pixel_limit", acFunPixelLimit, "转化后单个颜色对应像素个数限制")
	flag.BoolVar(&flags.Debug, "debug", false, "是否开启debug模式，开启将打印错误堆栈信息")
	flag.Float64Var(&flags.PixelCountThreshold, "pixel_count_threshold", 0, "过滤掉像素在图像帧占比小于当前值的颜色，会丢失图像细节")
	return &flags
}
