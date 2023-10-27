package repository

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
	"github.com/redis/go-redis/v9"
)

var (
	ErrDuplicate = errors.New("login already in use")
)

type PostgreSQL struct {
	db *sql.DB
}

type Redis struct {
	db *redis.Client
}

func NewPotgreSQL(addr string) (*PostgreSQL, error) {
	db, err := goose.OpenDBWithDriver("pgx", addr)
	if err != nil {
		return nil, fmt.Errorf("goose: failed to open DB: %v", err)
	}

	err = goose.Up(db, "./migrations")
	if err != nil {
		return nil, fmt.Errorf("goose: failed to migrate: %v", err)
	}

	return &PostgreSQL{
		db: db,
	}, nil
}

func (s *PostgreSQL) Close() error {
	return s.db.Close()
}

func NewRedis(addr string) (*Redis, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return &Redis{
		db: rdb,
	}, nil
}
