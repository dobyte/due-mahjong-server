package entity_test

import (
	"due-mahjong-server/game/app/entity"
	"fmt"
	"testing"
)

var (
	roomMgr   = entity.NewRoomMgr()
	playerMgr = entity.NewPlayerMgr()
)

func TestNewRoomMgr(t *testing.T) {
	fmt.Println(roomMgr.GetTable(1, 1))
}

func TestRoomMgr_QuickMatch(t *testing.T) {
	player, err := playerMgr.LoadPlayer(1)
	if err != nil {
		t.Fatalf("load player failed: %v", err)
	}

	err = roomMgr.QuickMatch(player)
	if err != nil {
		t.Fatalf("quick match failed: %v", err)
	}
}
