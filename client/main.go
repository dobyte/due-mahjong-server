package main

import (
	"due-mahjong-server/client/app/event"
	"due-mahjong-server/client/app/route"
	"github.com/dobyte/due"
	"github.com/dobyte/due/cluster/client"
	"github.com/dobyte/due/network/ws"
)

func main() {
	// 创建容器
	container := due.NewContainer()
	// 创建网关组件
	component := client.NewClient(
		client.WithClient(ws.NewClient()),
	)
	// 初始化事件和路由
	event.Init(component.Proxy())
	route.Init(component.Proxy())
	// 添加网关组件
	container.Add(component)
	// 启动容器
	container.Serve()
}
