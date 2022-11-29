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
)
