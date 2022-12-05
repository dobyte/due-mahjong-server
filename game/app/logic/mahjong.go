package logic

import (
	"context"
	"due-mahjong-server/game/app/entity"
	"due-mahjong-server/shared/pb/common"
	pb "due-mahjong-server/shared/pb/mahjong"
	"due-mahjong-server/shared/route"
	"github.com/dobyte/due/cluster/node"
	"github.com/dobyte/due/log"
)

type mahjong struct {
	proxy     node.Proxy
	ctx       context.Context
	roomMgr   *entity.RoomMgr
	playerMgr *entity.PlayerMgr
}

func NewMahjong(proxy node.Proxy) *mahjong {
	return &mahjong{
		proxy:     proxy,
		ctx:       context.Background(),
		roomMgr:   entity.NewRoomMgr(),
		playerMgr: entity.NewPlayerMgr(),
	}
}

func (l *mahjong) Init() {
	// 快速开始
	l.proxy.AddRouteHandler(route.QuickStart, false, l.quickStart)
}

// 快速开始
func (l *mahjong) quickStart(r node.Request) {
	res := &pb.QuickStartRes{}
	defer func() {
		if err := r.Response(res); err != nil {
			log.Errorf("read all mail response failed, err: %v", err)
		}
	}()

	if r.UID() <= 0 {
		log.Warnf("invalid user id: %d", r.UID())
		res.Code = common.Code_NotLogin
		return
	}

	player, err := l.playerMgr.LoadPlayer(r.UID())
	if err != nil {
		log.Errorf("load user info failed: uid: %d err: %v", r.UID(), err)
		res.Code = common.Code_Failed
		return
	}

	defer func() {
		if err != nil {
			l.playerMgr.UnloadPlayer(r.UID())
		}
	}()

	if err = l.roomMgr.QuickMatch(player); err != nil {
		log.Errorf("quick match failed: uid: %d err: %v", r.UID(), err)
		res.Code = common.Code_Failed
		return
	}

	if err = r.BindNode(); err != nil {
		log.Errorf("bind node failed: uid: %d err: %v", r.UID(), err)
		res.Code = common.Code_Failed
		return
	}

	res.Code = common.Code_OK
}
