package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/AlexCorn999/users/internal/domain"
	"github.com/jackc/pgx/v5/pgconn"
)

func (s *PostgreSQL) CreateUser(ctx context.Context, user domain.User) (int, error) {
	var id int
	err := s.db.QueryRowContext(ctx, "INSERT INTO users (login, age) values ($1, $2) RETURNING id",
		user.Login, user.Age).Scan(&id)
	if err != nil {
		var pgErr *pgconn.PgError
		if ok := errors.As(err, &pgErr); ok && pgErr.ConstraintName == "users_login_key" {
			return 0, ErrDuplicate
		}
		return 0, fmt.Errorf("postgreSQL: createUser %s", err)
	}
	return id, nil
}
