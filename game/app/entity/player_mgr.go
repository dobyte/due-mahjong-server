package entity

import (
	"due-mahjong-server/shared/service"
	"sync"
)

type PlayerMgr struct {
	userSvc *service.User

	rw      sync.RWMutex
	players map[int64]*Player
}

func NewPlayerMgr() *PlayerMgr {
	return &PlayerMgr{
		userSvc: service.NewUser(nil),
		players: make(map[int64]*Player),
	}
}

// LoadPlayer 加载玩家
// code.NotFoundUser
// code.InternalServerError
func (mgr *PlayerMgr) LoadPlayer(uid int64) (*Player, error) {
	user, err := mgr.userSvc.GetUser(uid)
	if err != nil {
		return nil, err
	}

	mgr.rw.Lock()
	defer mgr.rw.Unlock()

	player := newPlayer(user)
	mgr.players[uid] = player

	return player, nil
}

// UnloadPlayer 卸载玩家
func (mgr *PlayerMgr) UnloadPlayer(uid int64) {
	mgr.rw.Lock()
	defer mgr.rw.Unlock()

	player, ok := mgr.players[uid]
	if !ok {
		return
	}

	player.Reset()

	delete(mgr.players, uid)
}
