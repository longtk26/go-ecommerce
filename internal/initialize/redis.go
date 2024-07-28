package initialize

import (
	"context"
	"fmt"

	"github.com/longtk26/go-ecommerce/global"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var ctx = context.Background()

func InitRedis() {
	r := global.Config.Redis

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%v", r.Host, r.Port),
		Password: r.Password, // no password set
		DB:       r.Database, // use default DB
		PoolSize: 10,
	})

	_, err := rdb.Ping(ctx).Result()

	if err != nil {
		global.Logger.Error("Failed to connect to Redis", zap.Error(err))
	}

	fmt.Println("InitRedis is running")

	global.Rdb = rdb

	redisExample()	

}

func redisExample() {
	err := global.Rdb.Set(ctx, "score", 100, 0).Err()

	if err != nil {
		panic(err)
	}

	val, err := global.Rdb.Get(ctx, "score").Result()

	if err != nil {
		panic(err)
	}

	global.Logger.Info("score", zap.String("score", val))
}
