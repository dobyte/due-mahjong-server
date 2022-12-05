package route

import (
	"due-mahjong-server/game/app/logic"
	"github.com/dobyte/due/cluster/node"
)

func Init(proxy node.Proxy) {
	logic.NewMahjong(proxy).Init()
}
