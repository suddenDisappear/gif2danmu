package library

import "gif2danmu/infrastructure/resolver"

type Transformer interface {
	Transform() (resolver.Resolver, error)
}
