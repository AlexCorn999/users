package service

import "github.com/AlexCorn999/users/internal/domain"

type StorageRepository interface {
	AddValue(input *domain.RedisInput) (int, error)
}

type Storage struct {
	repo StorageRepository
}

func NewStorage(repo StorageRepository) *Storage {
	return &Storage{
		repo: repo,
	}
}

func (s *Storage) AddValue(input *domain.RedisInput) (int, error) {
	return s.repo.AddValue(input)
}
