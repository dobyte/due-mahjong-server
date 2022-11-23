package code

import "github.com/dobyte/due/code"

var (
	InternalServerError = code.NewCode(1001, "internal server error", nil)
	NotFoundUser        = code.NewCode(1020, "not found user", nil)
	WrongPassword       = code.NewCode(1021, "wrong password", nil)
)
