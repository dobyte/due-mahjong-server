package middleware

import (
	"due-mahjong-server/shared/route"
	"github.com/dobyte/due/cluster/node"
	"github.com/dobyte/due/log"
	"github.com/dobyte/due/session"
)

func Auth(ctx *node.Context) {
	if ctx.Request.UID == 0 {
		err := ctx.Proxy.Push(ctx.Context(), &node.PushArgs{
			GID:     ctx.Request.GID,
			Kind:    session.Conn,
			Target:  ctx.Request.CID,
			Message: &node.Message{Route: route.Unauthorized},
		})
		if err != nil {
			log.Errorf("response message failed, err: %v", err)
		}
	} else {
		ctx.Middleware.Next(ctx)
	}
}
