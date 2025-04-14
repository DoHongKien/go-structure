package initialize

import (
	"context"
	"fmt"

	"github.com/DoHongKien/go-structure/global"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var ctx = context.Background()

func InitRedis() {
	r := global.Config.Redis

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%v", r.Host, r.Port),
		Password: r.Password,
		DB:       r.DB,
		PoolSize: 10,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		global.Logger.Error("Redis initialization error: ", zap.Error(err))
	}

	global.Logger.Info("Initialization redis successfully")
	global.Rdb = rdb
}
