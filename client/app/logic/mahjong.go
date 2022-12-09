package logic

import (
	"due-mahjong-server/shared/pb/common"
	mahjongpb "due-mahjong-server/shared/pb/mahjong"
	"due-mahjong-server/shared/route"
	"github.com/dobyte/due/cluster/client"
	"github.com/dobyte/due/log"
)

type mahjong struct {
	proxy client.Proxy
}

func NewMahjong(proxy client.Proxy) *mahjong {
	return &mahjong{proxy: proxy}
}

func (l *mahjong) Init() {
	// 快速开始回执
	l.proxy.AddRouteHandler(route.QuickStart, l.quickStartAck)
	// 游戏信息通知
	l.proxy.AddRouteHandler(route.GameInfoNotify, l.gameInfoNotify)
}

func (l *mahjong) quickStartAck(r client.Request) {
	res := &mahjongpb.QuickStartRes{}

	err := r.Parse(res)
	if err != nil {
		log.Errorf("invalid quick start ack message, err: %v", err)
		return
	}

	if res.Code != common.Code_OK {
		log.Warnf("%v", res.Code)
		return
	}

	log.Info("quick start success")
	log.Infof("%+v", res)
}

func (l *mahjong) gameInfoNotify(r client.Request) {
	ntf := &mahjongpb.GameInfoNotify{}

	err := r.Parse(ntf)
	if err != nil {
		log.Errorf("invalid game info notify message, err: %v", err)
		return
	}

	log.Info("receive game info")
	log.Infof("%+v", ntf)
}
