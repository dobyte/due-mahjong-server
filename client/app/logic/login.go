package logic

import (
	"due-mahjong-server/client/app/store"
	commonpb "due-mahjong-server/shared/pb/common"
	loginpb "due-mahjong-server/shared/pb/login"
	"due-mahjong-server/shared/route"
	"github.com/dobyte/due/cluster/client"
	"github.com/dobyte/due/log"
)

type login struct {
	proxy client.Proxy
}

func NewLogin(proxy client.Proxy) *login {
	return &login{proxy: proxy}
}

func (l *login) Init() {
	// 用户注册
	l.proxy.AddRouteHandler(route.Register, l.register)
	// 用户登录
	l.proxy.AddRouteHandler(route.Login, l.login)
}

// 用户注册
func (l *login) register(r client.Request) {

}

// 用户登录
func (l *login) login(r client.Request) {
	res := &loginpb.LoginRes{}

	err := r.Parse(res)
	if err != nil {
		log.Errorf("invalid login message, err: %v", err)
		return
	}

	if res.Code != commonpb.Code_OK {
		log.Warnf("%+v", res)
		return
	}

	store.Token = res.Token

	log.Info("login success")

	err = l.proxy.Push(0, route.QuickStart, &commonpb.EmptyReq{})
	if err != nil {
		log.Errorf("push message failed: %v", err)
	}
}
