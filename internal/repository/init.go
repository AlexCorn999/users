package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/pressly/goose"
)

var (
	ErrDuplicate = errors.New("login already in use")
)

type Storage struct {
	db *sql.DB
}

// NewStorage инициализирует хранилище и применяет миграции.
func NewStorage(addr string) (*Storage, error) {
	db, err := goose.OpenDBWithDriver("pgx", addr)
	if err != nil {
		return nil, fmt.Errorf("goose: failed to open DB: %v", err)
	}

	err = goose.Up(db, "./migrations")
	if err != nil {
		return nil, fmt.Errorf("goose: failed to migrate: %v", err)
	}

	return &Storage{
		db: db,
	}, nil
}

// CloseDB закрывает подключение к базе данных.
func (s *Storage) Close() error {
	return s.db.Close()
}
