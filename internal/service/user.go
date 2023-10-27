package service

import (
	"context"

	"github.com/AlexCorn999/users/internal/domain"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *domain.User) (int, error)
}

type Users struct {
	repo UserRepository
}

func NewUsers(repo UserRepository) *Users {
	return &Users{
		repo: repo,
	}
}

func (u *Users) CreateUser(ctx context.Context, usr *domain.User) (int, error) {
	return u.repo.CreateUser(ctx, usr)
}
