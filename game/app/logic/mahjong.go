package logic

import (
	"context"
	"due-mahjong-server/game/app/entity"
	"due-mahjong-server/shared/pb/common"
	pb "due-mahjong-server/shared/pb/mahjong"
	"due-mahjong-server/shared/route"
	"github.com/dobyte/due/cluster"
	"github.com/dobyte/due/cluster/node"
	"github.com/dobyte/due/config"
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
	opts := make([]*entity.Options, 0)
	err := config.Get("mahjong.rooms").Scan(&opts)
	if err != nil {
		log.Fatalf("load mahjong rooms config failed: %v", err)
	}

	return &mahjong{
		proxy:     proxy,
		ctx:       context.Background(),
		roomMgr:   entity.NewRoomMgr(opts),
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
	l.proxy.AddRouteHandler(route.StandUp, true, l.standUp)
	// 开始准备
	l.proxy.AddRouteHandler(route.Ready, true, l.ready)
	// 取消准备
	l.proxy.AddRouteHandler(route.Unready, true, l.unready)
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

	l.syncSeatState(player, pb.SeatState_Offline)
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

	l.syncSeatState(player, pb.SeatState_Online)
}

// 快速开始
func (l *mahjong) quickStart(r node.Request) {
	res := &pb.QuickStartRes{}
	defer func() {
		if err := r.Response(res); err != nil {
			log.Errorf("quick start response failed, err: %v", err)
		}
	}()

	if r.UID() <= 0 {
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

	l.syncSeatState(player, pb.SeatState_SitDown)

	res.Code = common.Code_OK
	res.GameInfo = l.makeGameInfo(player)
}

// 坐下
func (l *mahjong) sitDown(r node.Request) {
	req := &pb.SitDownReq{}
	res := &pb.SitDownRes{}
	defer func() {
		if err := r.Response(res); err != nil {
			log.Errorf("sit down response failed, err: %v", err)
		}
	}()

	if r.UID() <= 0 {
		res.Code = common.Code_NotLogin
		return
	}

	if err := r.Parse(req); err != nil {
		log.Errorf("invalid sit down message, err: %v", err)
		res.Code = common.Code_Abnormal
		return
	}

	player, err := l.playerMgr.GetPlayer(r.UID())
	if err != nil {
		res.Code = common.Code_IllegalOperation
		return
	}

	seat := player.Seat()
	if seat != nil {
		res.Code = common.Code_IllegalOperation
		return
	}

}

// 站起
func (l *mahjong) standUp(r node.Request) {
	res := &pb.StandUpRes{}
	defer func() {
		if err := r.Response(res); err != nil {
			log.Errorf("start ready response failed, err: %v", err)
		}
	}()

	if r.UID() <= 0 {
		res.Code = common.Code_NotLogin
		return
	}

	player, err := l.playerMgr.GetPlayer(r.UID())
	if err != nil {
		res.Code = common.Code_IllegalOperation
		return
	}

	seat := player.Seat()
	if seat == nil {
		res.Code = common.Code_OK
		return
	}

	err = seat.RemPlayer()
	if err != nil {
		res.Code = common.Code_IllegalOperation
		return
	}

	l.syncSeatState(player, pb.SeatState_StandUp)

	res.Code = common.Code_OK
}

// 开始准备
func (l *mahjong) ready(r node.Request) {
	res := &pb.ReadyRes{}
	defer func() {
		if err := r.Response(res); err != nil {
			log.Errorf("start ready response failed, err: %v", err)
		}
	}()

	if r.UID() <= 0 {
		res.Code = common.Code_NotLogin
		return
	}

	player, err := l.playerMgr.GetPlayer(r.UID())
	if err != nil {
		res.Code = common.Code_IllegalOperation
		return
	}

	seat := player.Seat()
	if seat == nil {
		res.Code = common.Code_IllegalOperation
		return
	}

	if seat.IsReady() {
		res.Code = common.Code_OK
		return
	}

	seat.Ready()

	l.syncSeatState(player, pb.SeatState_Ready)

	res.Code = common.Code_OK
}

// 取消准备
func (l *mahjong) unready(r node.Request) {
	res := &pb.UnreadyRes{}
	defer func() {
		if err := r.Response(res); err != nil {
			log.Errorf("unready response failed, err: %v", err)
		}
	}()

	if r.UID() <= 0 {
		res.Code = common.Code_NotLogin
		return
	}

	player, err := l.playerMgr.GetPlayer(r.UID())
	if err != nil {
		res.Code = common.Code_IllegalOperation
		return
	}

	seat := player.Seat()
	if seat == nil {
		res.Code = common.Code_IllegalOperation
		return
	}

	if seat.IsUnready() {
		res.Code = common.Code_OK
		return
	}

	seat.Unready()

	l.syncSeatState(player, pb.SeatState_Unready)

	res.Code = common.Code_OK
}

// 同步玩家座位状态
func (l *mahjong) syncSeatState(player *entity.Player, state pb.SeatState) {
	seat := player.Seat()
	seats := seat.Table().Seats()
	targets := make([]int64, 0, len(seats))

	for _, s := range seats {
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

		targets = append(targets, p.UID())
	}

	if len(targets) == 0 {
		return
	}

	data := &pb.SeatStateNotify{
		SeatID:    int32(seat.ID()),
		SeatState: state,
	}

	if state == pb.SeatState_SitDown {
		data.Player = &pb.Player{User: l.makeUserInfo(player)}
	}

	_, _ = l.proxy.Multicast(l.ctx, &node.MulticastArgs{
		Kind:    session.User,
		Targets: targets,
		Message: &node.Message{
			Route: route.SeatState,
			Data:  data,
		},
	})
}

// 根据玩家生成用户信息
func (l *mahjong) makeUserInfo(player *entity.Player) *common.User {
	u := player.User()

	return &common.User{
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
	}
}

// 根据玩家生成游戏信息
func (l *mahjong) makeGameInfo(player *entity.Player) *pb.GameInfo {
	var (
		opts  = player.Room().Options()
		table = player.Table()
		seats = table.Seats()
	)

	info := &pb.GameInfo{
		Room: &pb.Room{
			ID:            int32(opts.ID),
			Name:          opts.Name,
			MinEntryLimit: int32(opts.MinEntryLimit),
			MaxEntryLimit: int32(opts.MaxEntryLimit),
		},
		Table: &pb.Table{
			ID:    int32(table.ID()),
			Seats: make([]*pb.Seat, len(seats)),
		},
	}

	for i, seat := range seats {
		info.Table.Seats[i] = &pb.Seat{
			ID: int32(seat.ID()),
		}

		p := seat.Player()
		if p == nil {
			continue
		}

		info.Table.Seats[i].Player = &pb.Player{
			IsMyself: player.UID() == p.UID(),
			User:     l.makeUserInfo(p),
		}
	}

	return info
}
