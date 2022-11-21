package service

import "due-mahjong-server/shared/service/types/mail"

type Mail struct {
}

func NewMail() *Mail {
	return &Mail{}
}

func (s *Mail) Send(receiver int64, sender mail.Sender, mail mail.Mail) {

}

func (s *Mail) List() {

}
