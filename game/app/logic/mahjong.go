package logic

import (
	"context"
	"due-mahjong-server/game/app/entity"
	"due-mahjong-server/shared/pb/common"
	pb "due-mahjong-server/shared/pb/mahjong"
	"due-mahjong-server/shared/route"
	"github.com/dobyte/due/cluster"
	"github.com/dobyte/due/cluster/node"
	"github.com/dobyte/due/log"
	"github.com/dobyte/due/session"
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
	// 断开连接
	l.proxy.AddEventListener(cluster.Disconnect, l.disconnect)
	// 重新连接
	l.proxy.AddEventListener(cluster.Reconnect, l.reconnect)
	// 快速开始
	l.proxy.AddRouteHandler(route.QuickStart, false, l.quickStart)
	// 坐下
	l.proxy.AddRouteHandler(route.SitDown, false, l.sitDown)
	// 起立
	l.proxy.AddRouteHandler(route.StandUp, false, l.standUp)
	// 开始准备
	l.proxy.AddRouteHandler(route.StartReady, true, l.startReady)
	// 取消准备
	l.proxy.AddRouteHandler(route.CancelReady, true, l.startReady)
}

// 断开连接
func (l *mahjong) disconnect(_ string, uid int64) {
	player, err := l.playerMgr.GetPlayer(uid)
	if err != nil {
		return
	}

	seat := player.Seat()
	if seat == nil {
		return
	}

	seat.Offline()

	message := &node.Message{
		Route: route.Offline,
		Data: &pb.OfflineNotify{
			SeatID: int32(seat.ID()),
		},
	}

	for _, s := range seat.Table().Seats() {
		if s.IsOffline() {
			continue
		}

		p := s.Player()
		if p == nil {
			continue
		}

		_ = l.proxy.Push(l.ctx, &node.PushArgs{
			Kind:    session.User,
			Target:  p.UID(),
			Message: message,
		})
	}
}

// 重新连接
func (l *mahjong) reconnect(_ string, uid int64) {
	player, err := l.playerMgr.GetPlayer(uid)
	if err != nil {
		return
	}

	seat := player.Seat()
	if seat == nil {
		return
	}

	seat.Online()

	message := &node.Message{
		Route: route.Online,
		Data: &pb.OnlineNotify{
			SeatID: int32(seat.ID()),
		},
	}

	for _, s := range seat.Table().Seats() {
		if s.IsOffline() {
			continue
		}

		p := s.Player()
		if p == nil {
			continue
		}

		if p.UID() == player.UID() {
			continue
		}

		_ = l.proxy.Push(l.ctx, &node.PushArgs{
			Kind:    session.User,
			Target:  p.UID(),
			Message: message,
		})
	}
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
	res.GameInfo = l.makeGameInfo(player)
}

// 坐下
func (l *mahjong) sitDown(r node.Request) {

}

// 站起
func (l *mahjong) standUp(r node.Request) {

}

// 开始准备
func (l *mahjong) startReady(r node.Request) {

}

// 取消准备
func (l *mahjong) cancelReady(r node.Request) {

}

// 根据玩家生成游戏信息
func (l *mahjong) makeGameInfo(player *entity.Player) *pb.GameInfo {
	var (
		opts  = player.Room().Options()
		table = player.Table()
		seats = table.Seats()
	)

	info := &pb.GameInfo{}
	info.Room = &pb.Room{
		ID:            int32(opts.ID),
		Name:          opts.Name,
		MinEntryLimit: int32(opts.MinEntryLimit),
		MaxEntryLimit: int32(opts.MaxEntryLimit),
	}
	info.Table = &pb.Table{
		ID:    int32(table.ID()),
		Seats: make([]*pb.Seat, len(seats)),
	}

	for i, seat := range seats {
		info.Table.Seats[i] = &pb.Seat{
			ID: int32(seat.ID()),
		}

		p := seat.Player()
		if p == nil {
			continue
		}

		u := p.User()

		info.Table.Seats[i].Player = &pb.Player{
			IsMyself: player.UID() == u.UID,
			User: &common.User{
				UID:           u.UID,
				Account:       u.Account,
				Nickname:      u.Nickname,
				Avatar:        u.Avatar,
				Signature:     u.Signature,
				Gender:        common.Gender(u.Gender),
				Level:         int32(u.Level),
				Experience:    int32(u.Experience),
				Coin:          int32(u.Coin),
				LastLoginIP:   u.LastLoginIP,
				LastLoginTime: u.LastLoginTime.Time().Unix(),
			},
		}
	}

	return info
}
