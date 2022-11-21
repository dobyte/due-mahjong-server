package counter

import (
	"due-mahjong-server/shared/dao/counter/internal"
	"go.mongodb.org/mongo-driver/mongo"
)

type Counter struct {
	*internal.Counter
}

func NewCounter(db *mongo.Database) *Counter {
	return &Counter{Counter: internal.NewCounter(db)}
}
