package event

import (
	"due-mahjong-server/client/app/store"
	pb "due-mahjong-server/shared/pb/login"
	"due-mahjong-server/shared/route"
	"github.com/dobyte/due/cluster"
	"github.com/dobyte/due/cluster/client"
	"github.com/dobyte/due/log"
)

func Init(proxy client.Proxy) {
	proxy.AddEventListener(cluster.Connect, func(proxy client.Proxy) {
		err := proxy.Push(0, route.Login, &pb.LoginReq{
			Mode:     pb.LoginMode_Guest,
			DeviceID: store.DeviceID,
		})
		if err != nil {
			log.Errorf("push message failed: %v", err)
		}
	})

	proxy.AddEventListener(cluster.Reconnect, func(proxy client.Proxy) {
		if store.Token == "" {
			err := proxy.Push(0, route.Login, &pb.LoginReq{
				Mode:     pb.LoginMode_Guest,
				DeviceID: store.DeviceID,
			})
			if err != nil {
				log.Errorf("push message failed: %v", err)
			}
		} else {
			err := proxy.Push(0, route.Login, &pb.LoginReq{
				Mode:     pb.LoginMode_Token,
				DeviceID: store.DeviceID,
				Token:    &store.Token,
			})
			if err != nil {
				log.Errorf("push message failed: %v", err)
			}
		}
	})

	//proxy.AddEventListener(cluster.Disconnect, func(proxy client.Proxy) {
	//	for {
	//		err := proxy.Reconnect()
	//		if err == nil {
	//			return
	//		}
	//
	//		time.Sleep(time.Second)
	//
	//		log.Errorf("reconnect failed: %v", err)
	//	}
	//})
}
