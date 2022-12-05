package entity

import (
	"due-mahjong-server/shared/code"
	"github.com/dobyte/due/errors"
	"sync"
)

type Seat struct {
	id    int
	table *Table

	rw     sync.RWMutex
	player *Player
}

func newSeat(id int, table *Table) *Seat {
	return &Seat{
		id:    id,
		table: table,
	}
}

// Reset 重置座位
func (s *Seat) Reset() {
	if s.player != nil {
		s.player.Reset()
	}
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
func (s *Seat) AddPlayer(player *Player) error {
	s.rw.Lock()
	defer s.rw.Unlock()

	if s.player != nil {
		return errors.NewError(code.SeatAlreadyTaken)
	}

	err := s.player.AddSeat(s)
	if err != nil {
		return err
	}

	s.player = player

	return nil
}

// HasPlayer 座位上是否有玩家
func (s *Seat) HasPlayer() bool {
	s.rw.RLock()
	defer s.rw.RUnlock()

	return s.player != nil
}
