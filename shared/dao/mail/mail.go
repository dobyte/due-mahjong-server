package mail

import (
	"context"
	"due-mahjong-server/shared/dao/mail/internal"
	"due-mahjong-server/shared/model/mail"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Columns = internal.Columns

type Mail struct {
	*internal.Mail
}

func NewMail(db *mongo.Database) *Mail {
	return &Mail{Mail: internal.NewMail(db)}
}

// FindOneByID 根据ID查找
func (dao *Mail) FindOneByID(ctx context.Context, id string) (*mail.Mail, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	return dao.FindOne(ctx, func(cols *internal.Columns) interface{} {
		return bson.M{cols.ID: objectID}
	})
}

// DeleteOneByID 根据ID删除
func (dao *Mail) DeleteOneByID(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = dao.DeleteOne(ctx, func(cols *internal.Columns) interface{} {
		return bson.M{cols.ID: objectID}
	})
	return err
}
