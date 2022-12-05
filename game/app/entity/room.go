package entity

import (
	"due-mahjong-server/shared/code"
	"github.com/dobyte/due/errors"
	"sync"
)

type Room struct {
	opts  *RoomOptions
	fixed bool

	rw         sync.RWMutex
	tables     map[int]*Table
	maxTableID int
}

type RoomOptions struct {
	ID            int    `json:"id"`            // 房间ID
	Name          string `json:"name"`          // 房间名称
	MinEntryLimit int    `json:"minEntryLimit"` // 最低进入限制
	MaxEntryLimit int    `json:"maxEntryLimit"` // 最高进入限制
	TotalTables   int    `json:"totalTables"`   // 房间总的牌桌数，为0时则为动态牌桌
	TotalSeats    int    `json:"totalSeats"`    // 牌桌总的座位数
}

func newRoom(opts *RoomOptions) *Room {
	r := &Room{
		opts:       opts,
		fixed:      opts.TotalTables > 0,
		tables:     make(map[int]*Table, opts.TotalTables),
		maxTableID: opts.TotalTables - 1,
	}

	for i := 1; i <= opts.TotalTables; i++ {
		r.tables[i] = newTable(i, r)
	}

	return r
}

// ID 获取房间ID
func (r *Room) ID() int {
	return r.opts.ID
}

// Options 获取房间配置
func (r *Room) Options() *RoomOptions {
	return r.opts
}

// GetTable 获取桌子
// code.NotFoundTable
func (r *Room) GetTable(tableID int) (*Table, error) {
	r.rw.RLock()
	defer r.rw.RUnlock()

	table, ok := r.tables[tableID]
	if !ok {
		return nil, errors.NewError(code.NotFoundTable)
	}

	return table, nil
}

// GetSeat 获取座位
// code.NotFoundTable
// code.IllegalParams
func (r *Room) GetSeat(tableID, seatID int) (*Seat, error) {
	r.rw.RLock()
	defer r.rw.RUnlock()

	table, ok := r.tables[tableID]
	if !ok {
		return nil, errors.NewError(code.NotFoundTable)
	}

	return table.GetSeat(seatID)
}

// CreatTable 创建桌子
// code.RoomIsFull
func (r *Room) CreatTable() (*Table, error) {
	if r.fixed {
		return nil, errors.NewError(code.RoomIsFull)
	}

	r.rw.Lock()
	defer r.rw.Unlock()

	return r.createTable(), nil
}

// 创建牌桌
func (r *Room) createTable() *Table {
	tid := -1
	for i := 0; i <= r.maxTableID; i++ {
		if _, ok := r.tables[i]; !ok {
			tid = i
			break
		}
	}

	if tid < 0 {
		r.maxTableID++
		tid = r.maxTableID
	}

	table := newTable(tid, r)
	r.tables[tid] = table

	return table
}

// DestroyTable 销毁牌桌
// code.NotFoundTable
func (r *Room) DestroyTable(tableID int) error {
	r.rw.Lock()
	defer r.rw.Unlock()

	table, ok := r.tables[tableID]
	if !ok {
		return errors.NewError(code.NotFoundTable)
	}

	table.Reset()

	if r.fixed {
		return nil
	}

	delete(r.tables, tableID)

	if tableID == r.maxTableID {
		r.maxTableID = 0
		for i := range r.tables {
			if i > r.maxTableID {
				r.maxTableID = i
			}
		}
	}

	return nil
}

// QuickMatch 快速匹配
// code.CoinsAreBelowEntryLimit
// code.CoinsAreOverEntryLimit
// code.NoMatchingTable
func (r *Room) QuickMatch(player *Player) error {
	coin := player.Coin()

	if r.opts.MinEntryLimit > 0 && coin < r.opts.MinEntryLimit {
		return errors.NewError(code.CoinsAreBelowEntryLimit)
	}

	if r.opts.MaxEntryLimit > 0 && coin > r.opts.MaxEntryLimit {
		return errors.NewError(code.CoinsAreOverEntryLimit)
	}

	r.rw.RLock()
	defer r.rw.RUnlock()

	for _, table := range r.tables {
		if table.QuickMatch(player) == nil {
			return nil
		}
	}

	if r.fixed {
		return errors.NewError(code.NoMatchingTable)
	}

	return r.createTable().AddPlayer(player)
}
