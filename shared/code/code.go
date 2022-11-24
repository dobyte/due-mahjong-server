package code

import "github.com/dobyte/due/code"

var (
	InternalServerError = code.NewCode(1001, "internal server error", nil)
	NoPermission        = code.NewCode(1002, "no permission", nil)
	TokenExpired        = code.NewCode(1003, "token is expired", nil)
	TokenInvalid        = code.NewCode(1004, "token is invalid", nil)
	NotFoundUser        = code.NewCode(1020, "not found user", nil)
	WrongPassword       = code.NewCode(1021, "wrong password", nil)
	NotFoundMail        = code.NewCode(1022, "not found mail", nil)
)
