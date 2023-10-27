package transport

import (
	"context"
	"net/http"

	"github.com/AlexCorn999/users/internal/config"
	"github.com/AlexCorn999/users/internal/repository"
	"github.com/AlexCorn999/users/internal/service"
	"github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"
)

type APIServer struct {
	config     *config.Config
	router     *chi.Mux
	postgreSQL *repository.PostgreSQL
	redis      *repository.Redis
	users      *service.Users
	sign       *service.Sign
	storage    *service.Storage
}

func NewAPIServer(config *config.Config) *APIServer {
	return &APIServer{
		config: config,
		router: chi.NewRouter(),
	}
}

func (s *APIServer) Start(ctx context.Context) error {
	s.configureRouter()

	db, err := s.configurePostgreSQL()
	if err != nil {
		return err
	}
	s.postgreSQL = db
	defer s.postgreSQL.Close()

	redisDB, err := s.configureRedis()
	if err != nil {
		return err
	}
	s.redis = redisDB

	s.users = service.NewUsers(db)
	s.sign = service.NewSign()
	s.storage = service.NewStorage(redisDB)

	log.Info("server starting...")

	return http.ListenAndServe(s.config.Port, s.router)
}

func (s *APIServer) configureRouter() {
	s.router.Use(withLogging)
	s.router.Post("/redis/incr", s.AddValue)
	s.router.Post("/sign/hmacsha512", s.SignHmacSha512)
	s.router.Post("/postgres/users", s.CreateUser)
}

func (s *APIServer) configurePostgreSQL() (*repository.PostgreSQL, error) {
	db, err := repository.NewPotgreSQL(s.config.PortPostgreSQL)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (s *APIServer) configureRedis() (*repository.Redis, error) {
	db, err := repository.NewRedis(s.config.PortRedis)
	if err != nil {
		return nil, err
	}
	return db, nil
}
