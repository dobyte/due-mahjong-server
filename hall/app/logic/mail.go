package logic

import (
	"context"
	"due-mahjong-server/shared/code"
	"due-mahjong-server/shared/pb/common"
	pb "due-mahjong-server/shared/pb/mail"
	"due-mahjong-server/shared/route"
	"due-mahjong-server/shared/service"
	"github.com/dobyte/due/cluster/node"
	"github.com/dobyte/due/errors"
	"github.com/dobyte/due/log"
)

type mail struct {
	proxy   node.Proxy
	ctx     context.Context
	mailSvc *service.Mail
}

func NewMail(proxy node.Proxy) *mail {
	return &mail{
		proxy:   proxy,
		ctx:     context.Background(),
		mailSvc: service.NewMail(proxy),
	}
}

func (l *mail) Init() {
	// 拉取邮件列表
	l.proxy.AddRouteHandler(route.FetchMailList, false, l.fetchList)
	// 读取邮件
	l.proxy.AddRouteHandler(route.ReadMail, false, l.read)
	// 删除邮件
	l.proxy.AddRouteHandler(route.DeleteMail, false, l.delete)
}

// 拉取邮件列表
func (l *mail) fetchList(r node.Request) {

}

// 读取邮件
func (l *mail) read(r node.Request) {

}

// 删除邮件
func (l *mail) delete(r node.Request) {
	req := &pb.DeleteReq{}
	res := &pb.DeleteRes{}
	defer func() {
		if err := r.Response(res); err != nil {
			log.Errorf("delete mail response failed, err: %v", err)
		}
	}()

	if err := r.Parse(req); err != nil {
		log.Errorf("invalid delete mail message, err: %v", err)
		res.Code = common.Code_Abnormal
		return
	}

	err := l.mailSvc.Delete(req.MailID, false)
	if err != nil {
		switch errors.Code(err) {
		case code.NotFoundMail:
			res.Code = common.Code_NotFound
		case code.NoPermission:
			res.Code = common.Code_NoPermission
		default:
			res.Code = common.Code_Failed
		}
		log.Errorf("delete mail failed, err: %v", err)
		return
	}

	res.Code = common.Code_OK
}
