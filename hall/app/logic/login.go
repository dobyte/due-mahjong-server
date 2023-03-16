package logic

import (
	"context"
	"due-mahjong-server/shared/code"
	"due-mahjong-server/shared/pb/common"
	pb "due-mahjong-server/shared/pb/login"
	"due-mahjong-server/shared/route"
	"due-mahjong-server/shared/service"
	"github.com/dobyte/due/cluster/node"
	"github.com/dobyte/due/errors"
	"github.com/dobyte/due/log"
)

type login struct {
	proxy    *node.Proxy
	ctx      context.Context
	loginSvc *service.Login
}

func NewLogin(proxy *node.Proxy) *login {
	return &login{
		proxy:    proxy,
		ctx:      context.Background(),
		loginSvc: service.NewLogin(proxy),
	}
}

func (l *login) Init() {
	l.proxy.Router().Group(func(group *node.RouterGroup) {
		// 用户注册
		group.AddRouteHandler(route.Register, false, l.register)
		// 用户登录
		group.AddRouteHandler(route.Login, false, l.login)
	})
}

// 用户注册
func (l *login) register(ctx *node.Context) {

}

// 用户登录
func (l *login) login(ctx *node.Context) {
	req := &pb.LoginReq{}
	res := &pb.LoginRes{}
	defer func() {
		if err := ctx.Response(res); err != nil {
			log.Errorf("login response failed, err: %v", err)
		}
	}()

	if err := ctx.Request.Parse(req); err != nil {
		log.Errorf("invalid login message, err: %v", err)
		res.Code = common.Code_Abnormal
		return
	}

	clientIP, err := ctx.GetIP()
	if err != nil {
		log.Errorf("get client ip failed, err: %v", err)
		res.Code = common.Code_Abnormal
		return
	}

	if req.GetDeviceID() == "" {
		res.Code = common.Code_IllegalParams
		return
	}

	var uid int64
	switch req.GetMode() {
	case pb.LoginMode_Guest:
		uid, err = l.loginSvc.GuestLogin(req.GetDeviceID(), clientIP)
	case pb.LoginMode_Mobile:
		uid, err = l.loginSvc.MobileLogin(req.GetMobile(), req.GetCode(), clientIP)
	case pb.LoginMode_Account:
		uid, err = l.loginSvc.AccountLogin(req.GetAccount(), req.GetPassword(), clientIP)
	case pb.LoginMode_Wechat:
		uid, err = l.loginSvc.WechatLogin(req.GetOpenid(), req.GetDeviceID(), clientIP)
	case pb.LoginMode_Google:
		uid, err = l.loginSvc.GoogleLogin(req.GetToken(), req.GetDeviceID(), clientIP)
	case pb.LoginMode_Token:
		uid, err = l.loginSvc.TokenLogin(req.GetToken(), clientIP)
	default:
		log.Errorf("login failed, err: %v", err)
		res.Code = common.Code_IllegalParams
		return
	}
	if err != nil {
		switch errors.Code(err) {
		case code.NotFoundUser, code.WrongPassword:
			res.Code = common.Code_IncorrectAccountOrPassword
		case code.TokenExpired:
			res.Code = common.Code_TokenExpired
		default:
			res.Code = common.Code_Failed
		}
		log.Errorf("login failed, err: %v", err)
		return
	}

	if err = ctx.BindGate(uid); err != nil {
		log.Errorf("bind gate failed, err: %v", err)
		res.Code = common.Code_Failed
		return
	}

	token, err := l.loginSvc.GenerateToken(uid)
	if err != nil {
		log.Errorf("token generate failed, err: %v", err)
		res.Code = common.Code_Failed
		return
	}

	res.Code = common.Code_OK
	res.Token = token
}
