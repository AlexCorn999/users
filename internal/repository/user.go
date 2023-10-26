package repository

import (
	"fmt"

	"github.com/AlexCorn999/users/internal/domain"
)

// CreateUser добавляет пользователя в базу данных.
func (s *Storage) CreateUser(user *domain.User) (int, error) {
	result, err := s.db.Exec("INSERT INTO users (login, age) values ($1, $2) on conflict (login) do nothing RETURNING id",
		user.Login, user.Age)
	if err != nil {
		return 0, fmt.Errorf("postgreSQL: createUser %s", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("postgreSQL: createUser %s", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("postgreSQL: createUser %s", err)
	}

	if rowsAffected == 0 {
		return 0, ErrDuplicate
	}

	return int(id), nil
}
