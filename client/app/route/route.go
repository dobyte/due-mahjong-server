package route

import (
	"due-mahjong-server/client/app/logic"
	"github.com/dobyte/due/cluster/client"
)

func Init(proxy client.Proxy) {
	// 登录逻辑
	logic.NewLogin(proxy).Init()
	// 邮件逻辑
	logic.NewMail(proxy).Init()
	// 麻将逻辑
	logic.NewMahjong(proxy).Init()
}
