package service

import (
	"context"
	"due-mahjong-server/shared/code"
	"due-mahjong-server/shared/components/mongo"
	userdao "due-mahjong-server/shared/dao/user"
	usermodel "due-mahjong-server/shared/model/user"
	"github.com/dobyte/due/cluster/node"
	"github.com/dobyte/due/errors"
)

type User struct {
	ctx     context.Context
	proxy   node.Proxy
	userDao *userdao.User
}

func NewUser(proxy node.Proxy) *User {
	return &User{
		ctx:     context.Background(),
		proxy:   proxy,
		userDao: userdao.NewUser(mongo.DB()),
	}
}

// GetUser 获取用户
// code.NotFoundUser
// code.InternalServerError
func (s *User) GetUser(uid int64) (*usermodel.User, error) {
	user, err := s.userDao.FindOneByUID(s.ctx, uid)
	if err != nil {
		return nil, errors.NewError(code.InternalServerError, err)
	}

	if user == nil {
		return nil, errors.NewError(code.NotFoundUser)
	}

	return user, nil
}
