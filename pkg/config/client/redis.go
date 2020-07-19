package client

import (
	"github.com/workfoxes/gobase/pkg/utils"
	"github.com/go-redis/redis"
	"log"
)

type RedisClient struct {
	client *redis.Client
	prefix string
}

func (rc *RedisClient) get(key string) interface{} {
	value, err := rc.client.Get(key).Result()
	if err == redis.Nil {
		log.Print("Key Does not exist in redis client", err)
	}
	return value
}

func NewRedisClient(prefix string) *RedisClient {
	if RedisConnection == nil {
		return nil
	}
	pong, err := RedisConnection.Ping().Result()
	utils.HandleError(err)
	log.Print(pong, err)
	return &RedisClient{
		prefix: prefix + "::",
		client: RedisConnection,
	}
}
