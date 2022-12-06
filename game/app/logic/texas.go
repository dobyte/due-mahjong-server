package logic

import (
	"context"
	"github.com/dobyte/due/cluster/node"
)

type texas struct {
	proxy node.Proxy
	ctx   context.Context
}

func NewTexas(proxy node.Proxy) *texas {
	return &texas{
		proxy: proxy,
		ctx:   context.Background(),
	}
}

func (l *texas) Init() {

}
