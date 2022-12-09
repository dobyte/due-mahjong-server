package logic

import (
	"context"
	"due-mahjong-server/game/app/entity"
	"github.com/dobyte/due/cluster/node"
	"github.com/dobyte/due/config"
	"github.com/dobyte/due/log"
)

type texas struct {
	proxy     node.Proxy
	ctx       context.Context
	roomMgr   *entity.RoomMgr
	playerMgr *entity.PlayerMgr
}

func NewTexas(proxy node.Proxy) *texas {
	opts := make([]*entity.Options, 0)
	err := config.Get("texas.rooms").Scan(&opts)
	if err != nil {
		log.Fatalf("load texas rooms config failed: %v", err)
	}

	return &texas{
		proxy:     proxy,
		ctx:       context.Background(),
		roomMgr:   entity.NewRoomMgr(opts),
		playerMgr: entity.NewPlayerMgr(),
	}
}

func (l *texas) Init() {

}
