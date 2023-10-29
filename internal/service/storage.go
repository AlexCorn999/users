package service

import (
	"context"

	"github.com/AlexCorn999/users/internal/domain"
	"github.com/AlexCorn999/users/internal/repository"
)

//go:generate mockgen -source=storage.go -destination=mocks/mockStorage.go

type StorageRepository interface {
	AddValue(ctx context.Context, input domain.RedisInput) (int, error)
}

type Storage struct {
	Repo StorageRepository
}

func NewStorage(repo *repository.Redis) *Storage {
	return &Storage{
		Repo: repo,
	}
}

func (s *Storage) AddValue(ctx context.Context, input domain.RedisInput) (int, error) {
	return s.Repo.AddValue(ctx, input)
}
