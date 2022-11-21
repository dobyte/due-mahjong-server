/**
 * @Author: fuxiao
 * @Email: 576101059@qq.com
 * @Date: 2022/9/25 1:59 下午
 * @Desc: 网关服务器
 */

package main

import (
	"github.com/dobyte/due"
	"github.com/dobyte/due/cluster/gate"
	"github.com/dobyte/due/locate/redis"
	"github.com/dobyte/due/network/ws"
	"github.com/dobyte/due/registry/etcd"
	"github.com/dobyte/due/transport/grpc"
)

func main() {
	// 创建容器
	container := due.NewContainer()
	// 创建网关组件
	component := gate.NewGate(
		gate.WithServer(ws.NewServer()),
		gate.WithLocator(redis.NewLocator()),
		gate.WithRegistry(etcd.NewRegistry()),
		gate.WithTransporter(grpc.NewTransporter()),
	)
	// 添加网关组件
	container.Add(component)
	// 启动容器
	container.Serve()
}
