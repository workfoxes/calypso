package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/workfoxes/gobase/pkg/config"
)

var R *redis.Client
var ctx = context.Background()

func New(config *config.Config) *redis.Client {
	R = redis.NewClient(&redis.Options{
		Addr:     config.REDIS_HOST,
		Password: config.REDIS_PASSWORD, // no password set
		DB:       0,                     // use default DB
	})
	return R
}

func Set(key string, value interface{}) {
	err := R.Set(ctx, key, value, 0).Err()
	if err != nil {
		panic(err)
	}
}

func Get(key string) string {
	val, err := R.Get(ctx, key).Result()
	if err != nil {
		panic(err)
	}
	return val
}
