package entity

import (
	"due-mahjong-server/shared/code"
	"due-mahjong-server/shared/model/user"
	"github.com/dobyte/due/errors"
	"sync"
)

type Player struct {
	user           *user.User
	coinInitAmount int

	rw             sync.RWMutex
	seat           *Seat
	coinIncrAmount int
}

func newPlayer(user *user.User) *Player {
	return &Player{
		user:           user,
		coinInitAmount: user.Coin,
	}
}

// Reset 重置用户
func (p *Player) Reset() {
	p.rw.Lock()
	defer p.rw.Unlock()

	p.syncToDB()
	p.seat = nil
}

// User 获取用户
func (p *Player) User() *user.User {
	u := p.user
	u.Coin = p.Coin()
	return u
}

// Coin 获取金币
func (p *Player) Coin() int {
	p.rw.RLock()
	defer p.rw.RUnlock()

	return p.coinInitAmount + p.coinIncrAmount
}

// Seat 获取玩家所属座位
func (p *Player) Seat() *Seat {
	p.rw.RLock()
	defer p.rw.RUnlock()

	return p.seat
}

// Table 获取玩家所属牌桌
func (p *Player) Table() *Table {
	p.rw.RLock()
	defer p.rw.RUnlock()

	if p.seat == nil {
		return nil
	}

	return p.seat.Table()
}

// Room 获取玩家所属房间
func (p *Player) Room() *Room {
	p.rw.RLock()
	defer p.rw.RUnlock()

	if p.seat == nil {
		return nil
	}

	return p.seat.Room()
}

// AddSeat 为玩家添加一个座位
// code.PlayerAlreadySeated
func (p *Player) AddSeat(seat *Seat) error {
	p.rw.Lock()
	defer p.rw.Unlock()

	if p.seat != nil {
		return errors.NewError(code.PlayerAlreadySeated)
	}

	p.seat = seat

	return nil
}

// IncrCoin 增加金币
func (p *Player) IncrCoin(incr int) {
	if incr == 0 {
		return
	}

	p.rw.Lock()
	defer p.rw.Unlock()

	p.coinIncrAmount += incr
}

// DecrCoin 扣减金币
func (p *Player) DecrCoin(decr int) {
	p.IncrCoin(0 - decr)
}

// SyncToDB 同步到数据库
func (p *Player) SyncToDB() {
	p.rw.Lock()
	defer p.rw.Unlock()

	p.syncToDB()
}

func (p *Player) syncToDB() {
	if p.coinIncrAmount == 0 {
		return
	}

	// TODO: 同步到数据库
	go func(coin int) {

	}(p.coinIncrAmount)

	p.coinInitAmount += p.coinIncrAmount
	p.coinIncrAmount = 0
	p.user.Coin = p.coinInitAmount
}
