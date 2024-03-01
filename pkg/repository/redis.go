package repository

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

type ConfigRedis struct {
	Addr     string
	Password string
	DB       int
}

func NewRedisClinet(cfg ConfigRedis) (*redis.Client, error) {
	fmt.Println(cfg.Addr, cfg.Password, cfg.DB)
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379", //cfg.Addr,
		Password: "",               // no password set
		DB:       0,                // use default DB
	})

	return client, nil
}
