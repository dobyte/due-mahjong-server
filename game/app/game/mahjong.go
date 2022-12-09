package game

import (
	"due-mahjong-server/game/app/entity"
)

type Mahjong struct {
	table *entity.Table
}

func NewMahjong(table *entity.Table) *Mahjong {
	return &Mahjong{
		table: table,
	}
}
