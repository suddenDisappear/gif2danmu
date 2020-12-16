package transform

type Transformer interface {
	Transform() (*ColorMap, error)
}
