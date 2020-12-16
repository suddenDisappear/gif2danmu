package transform

type ColorInfo struct {
	Contents   [][]string
	PixelCount int64
}

// ShouldIgnore 是否忽略当前颜色信息(根据颜色占比来判断).
func (c *ColorInfo) ShouldIgnore(total int64) bool {
	return float64(c.PixelCount) < float64(total)*conf.PixelCountThreshold
}
