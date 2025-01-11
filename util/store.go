package util

import (
	"context"
	"fmt"
	"github.com/rbcervilla/redisstore/v9"
	"github.com/redis/go-redis/v9"
	"log"
	"store/types/cfg"
)

var Store *redisstore.RedisStore

func InitStore(config *cfg.Config) *redisstore.RedisStore {
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", config.Redis.Host, config.Redis.Port),
	})

	var err error

	Store, err = redisstore.NewRedisStore(context.Background(), client)
	if err != nil {
		log.Fatal(err)
	}

	Store.KeyPrefix(config.Redis.Prefix)
	return Store
}
