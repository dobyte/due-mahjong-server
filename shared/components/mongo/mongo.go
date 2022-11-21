package mongo

import (
	"context"
	"github.com/dobyte/due/config"
	"github.com/dobyte/due/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"sync"
	"time"
)

var (
	once     sync.Once
	client   *mongo.Client
	database *mongo.Database
)

func Client() *mongo.Client {
	once.Do(func() {
		conf := &struct {
			DSN      string `json:"dsn"`
			Database string `json:"database"`
		}{}

		err := config.Get("config.mongo").Scan(conf)
		if err != nil {
			log.Fatalf("load mongo config failed: %v", err)
		}

		cliOpts := options.Client().ApplyURI(conf.DSN)
		baseCtx := context.Background()

		ctx, cancel := context.WithTimeout(baseCtx, 5*time.Second)
		cli, err := mongo.Connect(ctx, cliOpts)
		cancel()
		if err != nil {
			log.Fatalf("mongo connect failed: %v", err)
		}

		ctx, cancel = context.WithTimeout(baseCtx, 5*time.Second)
		err = cli.Ping(ctx, readpref.Primary())
		cancel()
		if err != nil {
			log.Fatalf("mongo connect failed: %v", err)
		}

		client = cli
		database = cli.Database(conf.Database)
	})

	return client
}

// DB 获取数据库
func DB() *mongo.Database {
	Client()
	return database
}
