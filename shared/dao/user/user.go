package user

import (
	"due-mahjong-server/shared/dao/user/internal"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	*internal.User
}

func NewUser(db *mongo.Database) *User {
	return &User{User: internal.NewUser(db)}
}
