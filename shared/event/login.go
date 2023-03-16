package event

import (
	"github.com/dobyte/due/eventbus"
	"github.com/dobyte/due/log"
)

type LoginPayload struct {
	ID      int32
	Account string
}

// LoginHandler 登录事件处理器
func LoginHandler(event *eventbus.Event) {
	log.Infof("%+v", event)
}
