package datastore

import (
	"github.com/redis/go-redis/v9"
)

type RedisAddr string

func ProvideRedisClient(redisAddr RedisAddr) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: string(redisAddr),
	})
	return client
}
