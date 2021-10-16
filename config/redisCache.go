package config

import (
	"context"
	"covid_cases_near_me/constants"
	"time"

	"github.com/go-redis/redis"
)

type RedisCache struct {
}

func NewRedisObject() *RedisCache {
	return &RedisCache{}
}

func GetRedisClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: constants.REDIS_ADDRESS,
		DB:   0, // use default DB
	})
	return rdb
}

func GetRedisContext() context.Context {
	var ctx = context.Background()
	return ctx
}

func (redisCache *RedisCache) Get(key string) (string, error) {
	return GetRedisClient().Get(key).Result()
}

func (redisCache *RedisCache) Set(key string, data string, ttl time.Duration) error {
	return GetRedisClient().Set(key, data, ttl*time.Second).Err()
}
