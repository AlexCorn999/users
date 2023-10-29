package repository

import (
	"context"
	"strconv"
	"sync/atomic"

	"github.com/AlexCorn999/users/internal/domain"
)

func (r *Redis) AddValue(ctx context.Context, input domain.RedisInput) (int, error) {
	err := r.db.Set(ctx, input.Key, input.Value, 0).Err()
	if err != nil {
		return 0, err
	}

	val, err := r.db.Get(ctx, input.Key).Result()
	if err != nil {
		return 0, err
	}

	value, err := strconv.Atoi(val)
	if err != nil {
		return 0, err
	}

	value32 := int32(value)
	atomic.AddInt32(&value32, 1)

	return int(value32), nil
}
