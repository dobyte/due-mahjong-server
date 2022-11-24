module due-mahjong-server

go 1.16

require (
	github.com/dobyte/due v0.0.6
	github.com/dobyte/due/locate/redis v0.0.6
	github.com/dobyte/due/network/ws v0.0.6
	github.com/dobyte/due/registry/etcd v0.0.6
	github.com/dobyte/due/transport/grpc v0.0.6
	github.com/dobyte/gen-mongo-dao v0.0.4
	github.com/dobyte/jwt v0.1.3
	github.com/gorilla/mux v1.8.0
	go.mongodb.org/mongo-driver v1.11.0
	golang.org/x/crypto v0.0.0-20220622213112-05595931fe9d
	google.golang.org/api v0.103.0
	google.golang.org/protobuf v1.28.1
)

replace (
	github.com/dobyte/due => ../due
	github.com/dobyte/due/locate/redis => ../due/locate/redis
	github.com/dobyte/due/network/ws => ../due/network/ws
	github.com/dobyte/due/registry/etcd => ../due/registry/etcd
	github.com/dobyte/due/transport/grpc => ../due/transport/grpc
)
