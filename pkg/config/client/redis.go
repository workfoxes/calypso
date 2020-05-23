package client

import (
	"github.com/go-redis/redis"
	"log"
)

type RedisClient struct {
	client *redis.Client
	prefix string
}

func (rc *RedisClient) get(key string) interface{}  {
	value, err := rc.client.Get(key).Result()
	if err == redis.Nil {
		log.Print("Key Does not exist in redis client", err)
	}
	return value
}

func NewRedisClient(accountId string)  *RedisClient  {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		DB:       0,  // use default DB
	})
	pong, err := client.Ping().Result()
	log.Print(pong, err)
	return &RedisClient{
		prefix: accountId+"::",
		client: client,
	}
}