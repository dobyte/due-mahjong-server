package logic

import (
	"context"
	"due-mahjong-server/shared/route"
	"github.com/dobyte/due/cluster/node"
)

type mail struct {
	proxy node.Proxy
	ctx   context.Context
}

func NewMail(proxy node.Proxy) *mail {
	return &mail{
		proxy: proxy,
		ctx:   context.Background(),
	}
}

func (l *mail) Init() {
	// 拉取邮件列表
	l.proxy.AddRouteHandler(route.FetchMailList, false, l.fetchList)
	// 读取邮件
	l.proxy.AddRouteHandler(route.ReadMail, false, l.read)
}

// 拉取邮件列表
func (l *mail) fetchList(r node.Request) {

}

// 读取邮件
func (l *mail) read(r node.Request) {

}
