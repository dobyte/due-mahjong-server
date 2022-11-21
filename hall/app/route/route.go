package route

import (
	"due-mahjong-server/hall/app/logic"
	"github.com/dobyte/due/cluster/node"
)

func Init(proxy node.Proxy) {
	// 登录逻辑
	logic.NewLogin(proxy).Init()
	// 邮件逻辑
	logic.NewMail(proxy).Init()
}
