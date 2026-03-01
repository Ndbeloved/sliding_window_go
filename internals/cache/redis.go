package cache

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

func New(addr string, ctx context.Context) (redisCli *redis.Client, err error) {
	rdb := redis.NewClient(&redis.Options{Addr: addr})
	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, err
	}
	log.Printf("✅ Connected to redis server")
	return rdb, nil
}
