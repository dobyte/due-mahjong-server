package logic

import (
	pb "due-mahjong-server/shared/pb/mail"
	"due-mahjong-server/shared/route"
	"github.com/dobyte/due/cluster/client"
	"github.com/dobyte/due/log"
)

type mail struct {
	proxy client.Proxy
}

func NewMail(proxy client.Proxy) *mail {
	return &mail{proxy: proxy}
}

func (l *mail) Init() {
	// 新邮件
	l.proxy.AddRouteHandler(route.NewMail, l.newMail)
}

// 新邮件
func (l *mail) newMail(r client.Request) {
	res := &pb.Mail{}

	err := r.Parse(res)
	if err != nil {
		log.Errorf("invalid new mail message, err: %v", err)
		return
	}

	log.Infof("receive a new mail, %+v", res)
}
