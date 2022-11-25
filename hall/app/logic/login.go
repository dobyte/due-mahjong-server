package logic

import (
	"context"
	"due-mahjong-server/shared/code"
	"due-mahjong-server/shared/components/jwt"
	"due-mahjong-server/shared/components/mongo"
	dao "due-mahjong-server/shared/dao/user"
	pb "due-mahjong-server/shared/pb/login"
	"due-mahjong-server/shared/route"
	"due-mahjong-server/shared/service"
	mailargs "due-mahjong-server/shared/service/args/mail"
	"github.com/dobyte/due/cluster/node"
	"github.com/dobyte/due/errors"
	"github.com/dobyte/due/log"
)

type login struct {
	proxy    node.Proxy
	ctx      context.Context
	jwt      *jwt.JWT
	userDao  *dao.User
	loginSvc *service.Login
}

func NewLogin(proxy node.Proxy) *login {
	return &login{
		proxy:    proxy,
		ctx:      context.Background(),
		jwt:      jwt.Instance(),
		userDao:  dao.NewUser(mongo.DB()),
		loginSvc: service.NewLogin(proxy),
	}
}

func (l *login) Init() {
	// 用户注册
	l.proxy.AddRouteHandler(route.Register, false, l.register)
	// 用户登录
	l.proxy.AddRouteHandler(route.Login, false, l.login)
}

// 用户注册
func (l *login) register(r node.Request) {

}

// 用户登录
func (l *login) login(r node.Request) {
	req := &pb.LoginReq{}
	res := &pb.LoginRes{}
	defer func() {
		if err := r.Response(res); err != nil {
			log.Errorf("login response failed, err: %v", err)
		}
	}()

	if err := r.Parse(req); err != nil {
		log.Errorf("invalid login message, err: %v", err)
		res.Code = pb.Code_Abnormal
		return
	}

	clientIP, err := r.GetIP()
	if err != nil {
		log.Errorf("get client ip failed, err: %v", err)
		res.Code = pb.Code_Abnormal
		return
	}

	if req.GetDeviceID() == "" {
		res.Code = pb.Code_IllegalParams
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
		res.Code = pb.Code_IllegalParams
		return
	}
	if err != nil {
		switch errors.Code(err) {
		case code.NotFoundUser, code.WrongPassword:
			res.Code = pb.Code_IncorrectAccountOrPassword
		case code.TokenExpired:
			res.Code = pb.Code_TokenExpired
		default:
			res.Code = pb.Code_Failed
		}
		log.Errorf("login failed, err: %v", err)
		return
	}

	token, err := l.jwt.GenerateToken(jwt.Payload{
		l.jwt.IdentityKey(): uid,
	})
	if err != nil {
		log.Errorf("token generate failed, err: %v", err)
		res.Code = pb.Code_Failed
		return
	}

	if err = r.BindGate(uid); err != nil {
		log.Errorf("bind gate failed, err: %v", err)
		res.Code = pb.Code_Failed
		return
	}

	res.Code = pb.Code_OK
	res.Token = token.Token

	go func() {
		_, err = service.NewMail(l.proxy).Send(uid, mailargs.Sender{
			ID: -1,
		}, mailargs.Mail{
			Title:   "Welcome to mahjong world",
			Content: "Hi, Dear player, Welcome to mahjong world",
		})
		if err != nil {
			log.Errorf("send mail failed: %v", err)
		}
	}()
}
