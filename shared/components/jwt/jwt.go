package jwt

import (
	"github.com/dobyte/due/config"
	"github.com/dobyte/due/log"
	"github.com/dobyte/jwt"
	"sync"
)

var (
	once     sync.Once
	instance jwt.JWT
)

type Payload = jwt.Payload

// JWT 获取JWT
func JWT() jwt.JWT {
	once.Do(func() {
		conf := &struct {
			Issuer      string `json:"issuer"`
			ExpiredTime int    `json:"expiredTime"`
			SecretKey   string `json:"secretKey"`
			IdentityKey string `json:"identityKey"`
		}{}

		err := config.Get("config.jwt").Scan(conf)
		if err != nil {
			log.Fatalf("load jwt config failed: %v", err)
		}

		ins, err := jwt.NewJwt(&jwt.Options{
			Issuer:      conf.Issuer,
			ExpiredTime: conf.ExpiredTime,
			SignMethod:  jwt.HS256,
			SecretKey:   conf.SecretKey,
			IdentityKey: conf.IdentityKey,
		})
		if err != nil {
			log.Fatalf("jwt init failed: %v", err)
		}

		instance = ins
	})

	return instance
}
