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
	l.proxy.AddRouteHandler(route.ReadMail, false, l.readMail)
	// 一键读取所有邮件
	l.proxy.AddRouteHandler(route.ReadAllMail, false, l.readAllMail)
	// 删除邮件
	l.proxy.AddRouteHandler(route.DeleteMail, false, l.deleteMail)
	// 删除所有邮件
	l.proxy.AddRouteHandler(route.DeleteAllMail, false, l.deleteAllMail)
}

// 拉取邮件列表
func (l *mail) fetchList(r node.Request) {

}

// 读取邮件
func (l *mail) readMail(r node.Request) {
	req := &pb.ReadMailReq{}
	res := &pb.ReadMailRes{}
	defer func() {
		if err := r.Response(res); err != nil {
			log.Errorf("read mail response failed, err: %v", err)
		}
	}()

	if err := r.Parse(req); err != nil {
		log.Errorf("invalid read mail message, err: %v", err)
		res.Code = common.Code_Abnormal
		return
	}

	err := l.mailSvc.ReadMail(req.MailID, r.UID())
	if err != nil {
		switch errors.Code(err) {
		case code.NoPermission:
			res.Code = common.Code_NoPermission
		case code.NotFoundMail:
			res.Code = common.Code_NotFound
		default:
			res.Code = common.Code_Failed
		}
		log.Errorf("read mail failed, err: %v", err)
		return
	}

	res.Code = common.Code_OK
}

// 一键读取所有邮件
func (l *mail) readAllMail(r node.Request) {
	res := &pb.ReadAllMailRes{}
	defer func() {
		if err := r.Response(res); err != nil {
			log.Errorf("read all mail response failed, err: %v", err)
		}
	}()

	err := l.mailSvc.ReadAllMail(r.UID())
	if err != nil {
		switch errors.Code(err) {
		case code.NoPermission:
			res.Code = common.Code_NoPermission
		case code.NotFoundMail:
			res.Code = common.Code_NotFound
		default:
			res.Code = common.Code_Failed
		}
		log.Errorf("read all mail failed, err: %v", err)
		return
	}

	res.Code = common.Code_OK
}

// 删除邮件
func (l *mail) deleteMail(r node.Request) {
	req := &pb.DeleteMailReq{}
	res := &pb.DeleteMailRes{}
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

	err := l.mailSvc.DeleteMail(req.MailID, r.UID(), false)
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

// 一键删除所有邮件
func (l *mail) deleteAllMail(r node.Request) {
	res := &pb.ReadAllMailRes{}
	defer func() {
		if err := r.Response(res); err != nil {
			log.Errorf("read all mail response failed, err: %v", err)
		}
	}()

	err := l.mailSvc.DeleteAllMail(r.UID(), false)
	if err != nil {
		res.Code = common.Code_Failed
		log.Errorf("read all mail failed, err: %v", err)
		return
	}

	res.Code = common.Code_OK
}
