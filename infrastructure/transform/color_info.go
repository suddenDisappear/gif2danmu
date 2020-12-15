package transform

type ColorInfo struct {
	Contents   [][]string
	PixelCount int64
}

func (c *ColorInfo) ShouldIgnore(total int64) bool {
	return float64(c.PixelCount) < float64(total)*0.01
}
