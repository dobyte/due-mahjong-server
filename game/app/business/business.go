package business

import (
	"context"
	"due-mahjong-server/game/app/logic"
	"due-mahjong-server/shared/event"
	"github.com/dobyte/due/cluster/node"
	"github.com/dobyte/due/eventbus"
	"github.com/dobyte/due/eventbus/kafka"
	"github.com/dobyte/due/log"
)

func Init(proxy *node.Proxy) {
	// 初始化路由
	initRoute(proxy)
	// 初始化事件
	initEvent()
}

// 初始化路由
func initRoute(proxy *node.Proxy) {
	// 麻将逻辑
	logic.NewMahjong(proxy).Init()
}

// 初始化事件
func initEvent() {
	// 初始化事件总线
	eventbus.SetEventbus(kafka.NewEventbus())
	// 订阅用户登录事件
	err := eventbus.Subscribe(context.Background(), event.Login, event.LoginHandler)
	if err != nil {
		log.Fatalf("%s event subscribe failed: %v", event.Login, err)
	}
}
