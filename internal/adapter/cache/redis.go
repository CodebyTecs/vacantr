package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
)

var (
	Ctx   = context.Background()
	Redis *redis.Client
)

func NewRedisClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})
	if err := rdb.Ping(Ctx).Err(); err != nil {
		log.Fatalf("cannot connect to Redis: %v", err)
	}
	return rdb
}

func InitRedis() {
	Redis = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})
	if err := Redis.Ping(Ctx).Err(); err != nil {
		log.Fatalf("cannot connect to Redis: %v", err)
	}
}
