package code

import "github.com/dobyte/due/code"

var (
	InternalServerError = code.NewCode(1, "internal server error", nil)
	NoPermission        = code.NewCode(2, "no permission", nil)
	IllegalParams       = code.NewCode(3, "illegal params", nil)
	TokenExpired        = code.NewCode(4, "token is expired", nil)
	TokenInvalid        = code.NewCode(5, "token is invalid", nil)

	NotFoundUser  = code.NewCode(1000, "not found user", nil)
	WrongPassword = code.NewCode(1001, "wrong password", nil)

	NotFoundMail = code.NewCode(1022, "not found mail", nil)

	NotFoundRoom            = code.NewCode(10000, "not found room", nil)
	RoomIsFull              = code.NewCode(10000, "room is full", nil)
	NotFoundTable           = code.NewCode(10000, "not found table", nil)
	TableIsFull             = code.NewCode(10000, "table is full", nil)
	TableIsEmpty            = code.NewCode(10000, "table is empty", nil)
	SeatIsEmpty             = code.NewCode(10001, "seat is empty", nil)
	SeatAlreadyTaken        = code.NewCode(10002, "seat already taken", nil)
	CoinsAreBelowEntryLimit = code.NewCode(10002, "coins are below the entry limit", nil)
	CoinsAreOverEntryLimit  = code.NewCode(10002, "coins are over the entry limit", nil)
	NoMatchingRoom          = code.NewCode(10000, "no matching room", nil)
	NoMatchingTable         = code.NewCode(10000, "no matching table", nil)
	NotFoundPlayer          = code.NewCode(10000, "not found player", nil)
	PlayerAlreadySeated     = code.NewCode(10000, "player is already seated", nil)
	PlayerNotInSeat         = code.NewCode(10000, "player not in seat", nil)
	PlayerIsReady           = code.NewCode(10000, "player is readying", nil)
)
