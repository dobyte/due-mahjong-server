package logic

import (
	"context"
	"due-mahjong-server/shared/code"
	"due-mahjong-server/shared/components/google"
	"due-mahjong-server/shared/components/jwt"
	"due-mahjong-server/shared/components/mongo"
	dao "due-mahjong-server/shared/dao/user"
	model "due-mahjong-server/shared/model/user"
	pb "due-mahjong-server/shared/pb/login"
	"due-mahjong-server/shared/route"
	"due-mahjong-server/shared/utils/xcrypt"
	"due-mahjong-server/shared/utils/xvalidate"
	"github.com/dobyte/due/cluster/node"
	"github.com/dobyte/due/config"
	"github.com/dobyte/due/errors"
	"github.com/dobyte/due/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type login struct {
	proxy   node.Proxy
	ctx     context.Context
	userDao *dao.User
}

func NewLogin(proxy node.Proxy) *login {
	return &login{
		proxy:   proxy,
		ctx:     context.Background(),
		userDao: dao.NewUser(mongo.DB()),
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

	var uid int64
	switch req.GetMode() {
	case pb.LoginMode_Guest:
		uid, err = l.doGuestLogin(req.GetDeviceID(), clientIP)
	case pb.LoginMode_Mobile:
		uid, err = l.doMobileLogin(req.GetMobile(), req.GetCode(), clientIP)
	case pb.LoginMode_Account:
		uid, err = l.doAccountLogin(req.GetAccount(), req.GetPassword(), clientIP)
	case pb.LoginMode_Wechat:
		uid, err = l.doWechatLogin(req.GetOpenid(), req.GetDeviceID(), clientIP)
	case pb.LoginMode_Google:
		uid, err = l.doGoogleLogin(req.GetToken(), req.GetDeviceID(), clientIP)
	case pb.LoginMode_Token:
		uid, err = l.doTokenLogin(req.GetToken(), clientIP)
	default:
		log.Errorf("login failed, err: %v", err)
		res.Code = pb.Code_IllegalParams
		return
	}
	if err != nil {
		switch errors.Code(err) {
		case code.NotFoundUser, code.WrongPassword:
			res.Code = pb.Code_IncorrectAccountOrPassword
		default:
			res.Code = pb.Code_Failed
		}
		log.Errorf("login failed, err: %v", err)
		return
	}

	token, err := l.doGenerateToken(uid)
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
	res.Token = token
}

// 来宾登录
func (l *login) doGuestLogin(deviceID string, clientIP string) (int64, error) {
	user, err := l.userDao.FindOneByDeviceID(l.ctx, model.TypeGuest, deviceID)
	if err != nil {
		return 0, err
	}

	if user == nil {
		user = &model.User{
			Type:        model.TypeGuest,
			DeviceID:    deviceID,
			RegisterIP:  clientIP,
			LastLoginIP: clientIP,
		}
		if err = l.doCreateUser(user); err != nil {
			return 0, err
		}
	} else {
		l.doUpdateLoginRecord(user.ID, clientIP)
	}

	return user.UID, nil
}

// 账号登录
func (l *login) doAccountLogin(account string, password string, clientIP string) (int64, error) {
	var (
		err  error
		user *model.User
	)

	switch {
	case xvalidate.IsEmail(account):
		user, err = l.userDao.FindOneByEmail(l.ctx, account)
	case xvalidate.IsMobile(account):
		user, err = l.userDao.FindOneByMobile(l.ctx, account)
	default:
		user, err = l.userDao.FindOneByAccount(l.ctx, account)
	}
	if err != nil {
		return 0, err
	}

	if user == nil {
		return 0, errors.NewError(code.NotFoundUser)
	}

	ok, err := xcrypt.Compare(user.Password, password, user.Salt)
	if err != nil {
		return 0, err
	}

	if !ok {
		return 0, errors.NewError(code.WrongPassword)
	}

	l.doUpdateLoginRecord(user.ID, clientIP)

	return user.UID, nil
}

// 手机号登录
func (l *login) doMobileLogin(mobile string, code string, clientIP string) (int64, error) {
	user, err := l.userDao.FindOneByMobile(l.ctx, mobile)
	if err != nil {
		return 0, err
	}

	if user == nil {
		user = &model.User{
			Type:        model.TypeGeneral,
			Mobile:      mobile,
			RegisterIP:  clientIP,
			LastLoginIP: clientIP,
		}
		if err = l.doCreateUser(user); err != nil {
			return 0, err
		}
	} else {
		l.doUpdateLoginRecord(user.ID, clientIP)
	}

	return user.UID, nil
}

// 微信登录
func (l *login) doWechatLogin(openid, deviceID, clientIP string) (int64, error) {
	return 0, nil
}

// Google登录
func (l *login) doGoogleLogin(idToken, deviceID, clientIP string) (int64, error) {
	tokenInfo, err := google.OAuth2().Tokeninfo().IdToken(idToken).Do()
	if err != nil {
		return 0, err
	}

	user, err := l.userDao.FindOneByGoogleUserId(l.ctx, tokenInfo.UserId)
	if err != nil {
		return 0, err
	}

	if user == nil {
		user = &model.User{
			Type:           model.TypeGeneral,
			ThirdPlatforms: model.ThirdPlatforms{Google: tokenInfo.UserId},
			DeviceID:       deviceID,
			RegisterIP:     clientIP,
			LastLoginIP:    clientIP,
		}
		if err = l.doCreateUser(user); err != nil {
			return 0, err
		}
	} else {
		l.doUpdateLoginRecord(user.ID, clientIP)
	}

	return user.UID, nil
}

// TOKEN登录
func (l *login) doTokenLogin(token string, clientIP string) (int64, error) {
	return 0, nil
}

// 创建用户
func (l *login) doCreateUser(user *model.User) error {
	user.Coin = 100

	_, err := l.userDao.InsertOne(l.ctx, user)
	return err
}

// 更新用户登录记录
func (l *login) doUpdateLoginRecord(id primitive.ObjectID, clientIP string) {
	_, err := l.userDao.UpdateOne(l.ctx, func(cols *dao.Columns) interface{} {
		return bson.M{cols.ID: id}
	}, func(cols *dao.Columns) interface{} {
		return bson.M{"$set": bson.M{
			cols.LastLoginIP:   clientIP,
			cols.LastLoginTime: primitive.NewDateTimeFromTime(time.Now()),
		}}
	})
	if err != nil {
		log.Errorf("update user's login record failed, err: %v", err)
	}
}

// 生成Token
func (l *login) doGenerateToken(uid int64) (string, error) {
	identityKey := config.Get("config.jwt.identityKey", "uid").String()

	token, err := jwt.JWT().GenerateToken(jwt.Payload{
		identityKey: uid,
	})
	if err != nil {
		return "", err
	}

	return token.Token, nil
}
