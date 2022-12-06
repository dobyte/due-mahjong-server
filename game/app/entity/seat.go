package entity

import (
	"due-mahjong-server/shared/code"
	"github.com/dobyte/due/errors"
	"sync"
	"sync/atomic"
)

type Seat struct {
	id    int
	table *Table

	rw      sync.RWMutex
	player  *Player // 玩家
	ready   int32   // 准备状态，0：未准备 1：已准备
	offline int32   // 离线状态，0：在线 1：离线
}

func newSeat(id int, table *Table) *Seat {
	return &Seat{
		id:    id,
		table: table,
	}
}

// Reset 重置座位
func (s *Seat) Reset() {
	s.rw.Lock()
	defer s.rw.Unlock()

	if s.player != nil {
		s.player.Reset()
		s.player = nil
	}

	s.offline = 0
}

// ID 获取座位ID
func (s *Seat) ID() int {
	return s.id
}

// Table 获取座位所属牌桌
func (s *Seat) Table() *Table {
	return s.table
}

// Room 获取座位所属房间
func (s *Seat) Room() *Room {
	return s.table.Room()
}

// Player 获取座位上的玩家
func (s *Seat) Player() *Player {
	s.rw.RLock()
	defer s.rw.RUnlock()

	return s.player
}

// AddPlayer 为座位添加玩家
// code.SeatAlreadyTaken
// code.PlayerAlreadySeated
func (s *Seat) AddPlayer(player *Player) error {
	s.rw.Lock()
	defer s.rw.Unlock()

	if s.player != nil {
		return errors.NewError(code.SeatAlreadyTaken)
	}

	err := player.AddSeat(s)
	if err != nil {
		return err
	}

	if s.table.room.IsAutoReady() {
		s.Ready()
	}

	s.player = player

	return nil
}

// RemPlayer 移除玩家座位
// code.SeatIsEmpty
// code.PlayerNotInSeat
func (s *Seat) RemPlayer() error {
	s.rw.Lock()
	defer s.rw.Unlock()

	if s.player == nil {
		return errors.NewError(code.SeatIsEmpty)
	}

	err := s.player.RemSeat()
	if err != nil {
		return err
	}

	s.player = nil

	return nil
}

// HasPlayer 座位上是否有玩家
func (s *Seat) HasPlayer() bool {
	s.rw.RLock()
	defer s.rw.RUnlock()

	return s.player != nil
}

// IsOffline 检测座位上的玩家是否离线
func (s *Seat) IsOffline() bool {
	return atomic.LoadInt32(&s.offline) == 1
}

// IsOnline 检测座位上的玩家是否在线
func (s *Seat) IsOnline() bool {
	return !s.IsOffline()
}

// Offline 标记座位上的玩家离线
func (s *Seat) Offline() {
	atomic.StoreInt32(&s.offline, 1)
}

// Online 标记座位上的玩家上线
func (s *Seat) Online() {
	atomic.StoreInt32(&s.offline, 0)
}

// Ready 准备
func (s *Seat) Ready() {
	atomic.StoreInt32(&s.ready, 1)
}

// CancelReady 取消准备
func (s *Seat) CancelReady() {
	atomic.StoreInt32(&s.ready, 0)
}

// IsReady 是否已准备
func (s *Seat) IsReady() bool {
	return atomic.LoadInt32(&s.ready) == 1
}

// IsUnready 是否未准备好
func (s *Seat) IsUnready() bool {
	return !s.IsReady()
}
