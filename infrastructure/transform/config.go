package transform

import (
	"gif2danmu/infrastructure/flags"
	"sync"
)

var (
	conf Config
	once sync.Once
)

type Config struct {
	Fill        string
	PixelSymbol string
	PixelLimit  uint64
	OutputDir   string
}

// InitConfig 转换配置(仅第一次调用有效).
func InitConfig(flags flags.Flags) *Config {
	once.Do(func() {
		conf = Config{
			Fill:        flags.Fill,
			PixelSymbol: flags.PixelSymbol,
			PixelLimit:  flags.PixelLimit,
			OutputDir:   flags.OutputDir,
		}
	})
	return &conf
}

// GetConfig 获取转换配置.
func GetConfig() Config {
	return conf
}
