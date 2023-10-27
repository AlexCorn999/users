package repository

import (
	"context"
	"strconv"

	"github.com/AlexCorn999/users/internal/domain"
)

func (r *Redis) AddValue(input *domain.RedisInput) (int, error) {
	err := r.db.Set(context.Background(), input.Key, input.Value, 0).Err()
	if err != nil {
		return 0, err
	}

	val, err := r.db.Get(context.Background(), input.Key).Result()
	if err != nil {
		return 0, err
	}

	value, err := strconv.Atoi(val)
	if err != nil {
		return 0, err
	}

	// сделать потокобезопасным
	value++

	return value, nil
}
