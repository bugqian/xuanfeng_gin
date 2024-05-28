package db

import (
	"context"
	"github.com/redis/go-redis/v9"
)

type RedisConf struct {
	Addr         string
	Password     string
	DB           int
	MinIdleConns int
	Cluster      []string
}

func NewRedis(conf *RedisConf) (*redis.Client, error) {
	cli := redis.NewClient(&redis.Options{
		Addr:         conf.Addr,
		Password:     conf.Password,
		DB:           conf.DB,
		MinIdleConns: conf.MinIdleConns,
	})
	_, err := cli.Ping(context.Background()).Result()

	return cli, err
}
