package main

import (
	"due-mahjong-server/hall/app/route"
	"github.com/dobyte/due"
	"github.com/dobyte/due/cluster/node"
	"github.com/dobyte/due/locate/redis"
	"github.com/dobyte/due/registry/etcd"
	"github.com/dobyte/due/transport/grpc"
)

func main() {
	// 创建容器
	container := due.NewContainer()
	// 创建网关组件user
	component := node.NewNode(
		node.WithLocator(redis.NewLocator()),
		node.WithRegistry(etcd.NewRegistry()),
		node.WithTransporter(grpc.NewTransporter()),
	)
	// 初始化路由
	route.Init(component.Proxy())
	// 添加网关组件
	container.Add(component)
	// 启动容器
	container.Serve()
}
