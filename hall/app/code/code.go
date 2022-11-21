package code

import "github.com/dobyte/due/code"

var (
	NotFoundUser  = code.NewCode(1001, "not found user", nil)
	WrongPassword = code.NewCode(1002, "wrong password", nil)
)
