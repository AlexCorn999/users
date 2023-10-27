package repository

import "github.com/redis/go-redis/v9"

type Redis struct {
	db *redis.Client
}

func NewRedis(addr string) (*Redis, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})

	return &Redis{
		db: rdb,
	}, nil
}
