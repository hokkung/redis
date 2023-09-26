package v9

import (
	"github.com/hokkung/redis/config"
	"github.com/redis/go-redis/v9"
)

type RedisV9 struct {
	*redis.Client
}

func NewRedisConnection() *RedisV9 {
	cfg := config.New()
	rdb := redis.NewClient(&redis.Options{
		Addr:       cfg.Addr,
		Password:   cfg.Password,
		DB:         cfg.DB,
		MaxRetries: cfg.DB,
	})
	return &RedisV9{rdb}
}

func ProvideRedisConnection() *RedisV9 {
	return NewRedisConnection()
}
