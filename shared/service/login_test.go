package service_test

import (
	"due-mahjong-server/shared/service"
	"testing"
)

func TestLogin_GuestLogin(t *testing.T) {
	svc := service.NewLogin(nil)

	uid, err := svc.GuestLogin("1", "127.0.0.1")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(uid)
}
