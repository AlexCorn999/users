package service

import (
	"context"

	"github.com/AlexCorn999/users/internal/domain"
)

//go:generate mockgen -source=user.go -destination=mocks/mockUser.go

type UserRepository interface {
	CreateUser(ctx context.Context, user domain.User) (int, error)
}

type Users struct {
	Repo UserRepository
}

func NewUsers(repo UserRepository) *Users {
	return &Users{
		Repo: repo,
	}
}

func (u *Users) CreateUser(ctx context.Context, usr domain.User) (int, error) {
	return u.Repo.CreateUser(ctx, usr)
}
