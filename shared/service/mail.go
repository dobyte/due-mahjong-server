package service

import (
	"context"
	"due-mahjong-server/shared/code"
	"due-mahjong-server/shared/components/mongo"
	maildao "due-mahjong-server/shared/dao/mail"
	userdao "due-mahjong-server/shared/dao/user"
	mailmodel "due-mahjong-server/shared/model/mail"
	mailpb "due-mahjong-server/shared/pb/mail"
	"due-mahjong-server/shared/route"
	mailargs "due-mahjong-server/shared/service/args/mail"
	"github.com/dobyte/due/cluster/node"
	"github.com/dobyte/due/errors"
	"github.com/dobyte/due/session"
)

type Mail struct {
	proxy   node.Proxy
	ctx     context.Context
	userDao *userdao.User
	mailDao *maildao.Mail
}

func NewMail(proxy node.Proxy) *Mail {
	return &Mail{
		proxy:   proxy,
		ctx:     context.Background(),
		userDao: userdao.NewUser(mongo.DB()),
		mailDao: maildao.NewMail(mongo.DB()),
	}
}

// Send 发送邮件
func (s *Mail) Send(receiver int64, sender mailargs.Sender, mail mailargs.Mail) (string, error) {
	user, err := s.userDao.FindOneByUID(s.ctx, receiver)
	if err != nil {
		return "", errors.NewError(code.InternalServerError, err)
	}

	if user == nil {
		return "", errors.NewError(code.NotFoundUser, "not found receiver")
	}

	if sender.ID > 0 {
		user, err = s.userDao.FindOneByUID(s.ctx, sender.ID)
		if err != nil {
			return "", errors.NewError(code.InternalServerError, err)
		}

		if user == nil {
			return "", errors.NewError(code.NotFoundUser, "not found sender")
		}

		if sender.Name == "" {
			sender.Name = user.Nickname
		}

		if sender.Icon == "" {
			sender.Icon = user.Avatar
		}
	}

	model := &mailmodel.Mail{
		Title:       mail.Title,
		Content:     mail.Content,
		Receiver:    receiver,
		Status:      mailmodel.StatusUnread,
		Attachments: make([]mailmodel.Attachment, 0, len(mail.Attachments)),
		Sender: mailmodel.Sender{
			ID:   sender.ID,
			Name: sender.Name,
			Icon: sender.Icon,
		},
	}

	for _, attachment := range mail.Attachments {
		model.Attachments = append(model.Attachments, mailmodel.Attachment{
			PropID:  attachment.PropID,
			PropNum: attachment.PropNum,
		})
	}

	_, err = s.mailDao.InsertOne(s.ctx, model)
	if err != nil {
		return "", errors.NewError(code.InternalServerError, err)
	}

	if s.proxy != nil {
		data := &mailpb.Mail{
			ID:          model.ID.String(),
			Title:       model.Title,
			Content:     model.Content,
			Status:      mailpb.Status(model.Status),
			SendTime:    model.SendTime.Time().Unix(),
			Attachments: make([]*mailpb.Attachment, 0, len(mail.Attachments)),
			Sender: &mailpb.Sender{
				ID:   model.Sender.ID,
				Name: model.Sender.Name,
				Icon: model.Sender.Icon,
			},
		}

		for _, attachment := range mail.Attachments {
			data.Attachments = append(data.Attachments, &mailpb.Attachment{
				PropID:  int32(attachment.PropID),
				PropNum: int32(attachment.PropNum),
			})
		}

		_ = s.proxy.Push(s.ctx, &node.PushArgs{
			Kind:    session.User,
			Target:  receiver,
			Message: &node.Message{Route: route.NewMail, Data: data},
		})
	}

	return model.ID.String(), nil
}

// Delete 删除邮件
func (s *Mail) Delete(mailID string, isForce bool) error {
	mail, err := s.mailDao.FindOneByID(s.ctx, mailID)
	if err != nil {
		return errors.NewError(code.InternalServerError, err)
	}

	if mail == nil {
		return nil
	}

	if !isForce {
		if mail.Status == mailmodel.StatusUnread {
			return errors.NewError(code.NoPermission, "cannot delete unread mail")
		}

		if len(mail.Attachments) > 0 && mail.Status != mailmodel.StatusReceived {
			return errors.NewError(code.NoPermission, "cannot delete unreceived mail")
		}
	}

	_, err = s.mailDao.DeleteOneByID(s.ctx, mailID)
	if err != nil {
		return errors.NewError(code.InternalServerError, err)
	}

	return nil
}

func (s *Mail) List() {

}
