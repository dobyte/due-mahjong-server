package logic

import (
	"context"
	"due-mahjong-server/shared/code"
	"due-mahjong-server/shared/middleware"
	"due-mahjong-server/shared/pb/common"
	pb "due-mahjong-server/shared/pb/mail"
	"due-mahjong-server/shared/route"
	"due-mahjong-server/shared/service"
	"github.com/dobyte/due/cluster/node"
	"github.com/dobyte/due/errors"
	"github.com/dobyte/due/log"
)

type mail struct {
	proxy   *node.Proxy
	ctx     context.Context
	mailSvc *service.Mail
}

func NewMail(proxy *node.Proxy) *mail {
	return &mail{
		proxy:   proxy,
		ctx:     context.Background(),
		mailSvc: service.NewMail(proxy),
	}
}

func (l *mail) Init() {
	l.proxy.Router().Group(func(group *node.RouterGroup) {
		// 注册中间件
		group.Middleware(middleware.Auth)
		// 拉取邮件列表
		group.AddRouteHandler(route.FetchMailList, false, l.fetchList)
		// 读取邮件
		group.AddRouteHandler(route.ReadMail, false, l.readMail)
		// 一键读取所有邮件
		group.AddRouteHandler(route.ReadAllMail, false, l.readAllMail)
		// 删除邮件
		group.AddRouteHandler(route.DeleteMail, false, l.deleteMail)
		// 删除所有邮件
		group.AddRouteHandler(route.DeleteAllMail, false, l.deleteAllMail)
	})

}

// 拉取邮件列表
func (l *mail) fetchList(ctx *node.Context) {

}

// 读取邮件
func (l *mail) readMail(ctx *node.Context) {
	req := &pb.ReadMailReq{}
	res := &pb.ReadMailRes{}
	defer func() {
		if err := ctx.Response(res); err != nil {
			log.Errorf("read mail response failed, err: %v", err)
		}
	}()

	if err := ctx.Request.Parse(req); err != nil {
		log.Errorf("invalid read mail message, err: %v", err)
		res.Code = common.Code_Abnormal
		return
	}

	err := l.mailSvc.ReadMail(req.MailID, ctx.Request.UID)
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
func (l *mail) readAllMail(ctx *node.Context) {
	res := &pb.ReadAllMailRes{}
	defer func() {
		if err := ctx.Response(res); err != nil {
			log.Errorf("read all mail response failed, err: %v", err)
		}
	}()

	err := l.mailSvc.ReadAllMail(ctx.Request.UID)
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
func (l *mail) deleteMail(ctx *node.Context) {
	req := &pb.DeleteMailReq{}
	res := &pb.DeleteMailRes{}
	defer func() {
		if err := ctx.Response(res); err != nil {
			log.Errorf("delete mail response failed, err: %v", err)
		}
	}()

	if err := ctx.Request.Parse(req); err != nil {
		log.Errorf("invalid delete mail message, err: %v", err)
		res.Code = common.Code_Abnormal
		return
	}

	err := l.mailSvc.DeleteMail(req.MailID, ctx.Request.UID, false)
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
func (l *mail) deleteAllMail(ctx *node.Context) {
	res := &pb.ReadAllMailRes{}
	defer func() {
		if err := ctx.Response(res); err != nil {
			log.Errorf("read all mail response failed, err: %v", err)
		}
	}()

	err := l.mailSvc.DeleteAllMail(ctx.Request.UID, false)
	if err != nil {
		res.Code = common.Code_Failed
		log.Errorf("read all mail failed, err: %v", err)
		return
	}

	res.Code = common.Code_OK
}
