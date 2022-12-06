package route

import (
	"due-mahjong-server/game/app/logic"
	"github.com/dobyte/due/cluster/node"
)

func Init(proxy node.Proxy) {
	// 麻将逻辑
	logic.NewMahjong(proxy).Init()
	// 德州逻辑
	logic.NewTexas(proxy).Init()
}
