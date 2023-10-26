package service

import (
	"github.com/AlexCorn999/users/internal/domain"
)

type UserRepository interface {
	CreateUser(user *domain.User) (int, error)
}

type Users struct {
	repo UserRepository
}

func NewUsers(repo UserRepository) *Users {
	return &Users{
		repo: repo,
	}
}

// CreateUser отвечает за добавление пользователя в базу данных.
func (u *Users) CreateUser(usr *domain.User) (int, error) {
	user := &domain.User{
		Login: usr.Login,
		Age:   usr.Age,
	}
	return u.repo.CreateUser(user)
}
