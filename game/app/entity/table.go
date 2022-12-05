package entity

import (
	"due-mahjong-server/shared/code"
	"github.com/dobyte/due/errors"
	"sync"
)

type Table struct {
	id int

	rw           sync.RWMutex
	room         *Room
	seats        []*Seat
	totalPlayers int
}

func newTable(id int, room *Room) *Table {
	t := &Table{
		id:    id,
		room:  room,
		seats: make([]*Seat, room.opts.TotalSeats),
	}

	for i := 0; i < room.opts.TotalSeats; i++ {
		t.seats[i] = newSeat(i, t)
	}

	return t
}

// GetSeat 获取座位
func (t *Table) GetSeat(seatID int) (*Seat, error) {
	t.rw.RLock()
	defer t.rw.RUnlock()

	if seatID <= 0 || seatID >= len(t.seats) {
		return nil, errors.NewError(code.IllegalParams)
	}

	return t.seats[seatID-1], nil
}

// AddPlayer 把玩家加入牌桌
// code.IllegalParams
// code.TableIsFull
// code.SeatAlreadyTaken
// code.PlayerAlreadySeated
func (t *Table) AddPlayer(player *Player, seatID ...int) error {
	t.rw.Lock()
	defer t.rw.Unlock()

	sid := -1
	if len(seatID) > 0 {
		sid = seatID[0] - 1

		if sid < 0 || sid >= len(t.seats) {
			return errors.NewError(code.IllegalParams)
		}
	} else {
		for n, seat := range t.seats {
			if !seat.HasPlayer() {
				sid = n
				break
			}
		}

		if sid < 0 {
			return errors.NewError(code.TableIsFull)
		}
	}

	err := t.seats[sid].AddPlayer(player)
	if err != nil {
		return err
	}

	t.totalPlayers++

	return nil
}

// RemoveFromSeat 把玩家从座位上移除
// code.IllegalParams
// code.SeatIsEmpty
func (t *Table) RemoveFromSeat(seatID int) error {
	t.rw.Lock()
	defer t.rw.Unlock()

	if seatID < 0 || seatID >= len(t.seats) {
		return errors.NewError(code.IllegalParams)
	}

	if t.seats[seatID] == nil {
		return errors.NewError(code.SeatIsEmpty)
	}

	t.seats[seatID] = nil
	t.totalPlayers--

	return nil
}

// Reset 重置牌桌
func (t *Table) Reset() {
	t.rw.Lock()
	defer t.rw.Unlock()

	if t.room != nil && !t.room.fixed {
		t.room = nil
	}

	for i := range t.seats {
		t.seats[i].Reset()
	}

	t.seats = make([]*Seat, len(t.seats))
	t.totalPlayers = 0
}

// Room 获取牌桌所属房间
func (t *Table) Room() *Room {
	t.rw.RLock()
	defer t.rw.RUnlock()

	return t.room
}

func (t *Table) QuickMatch(player *Player) error {
	return nil
}
