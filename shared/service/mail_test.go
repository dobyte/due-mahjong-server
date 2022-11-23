package service_test

import (
	"due-mahjong-server/shared/service"
	mailargs "due-mahjong-server/shared/service/args/mail"
	"testing"
)

func TestMail_Send(t *testing.T) {
	svc := service.NewMail(nil)

	mailID, err := svc.Send(1, mailargs.Sender{
		ID: -1,
	}, mailargs.Mail{
		Title:   "A Test Mail",
		Content: "test mail content",
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(mailID)
}
