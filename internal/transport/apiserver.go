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
	postgreSQL *repository.Storage
	users      *service.Users
}

func NewAPIServer(config *config.Config) *APIServer {
	return &APIServer{
		config: config,
		router: chi.NewRouter(),
	}
}

func (s *APIServer) Start(ctx context.Context) error {
	s.configureRouter()

	db, err := s.configureStore()
	if err != nil {
		return err
	}
	s.postgreSQL = db
	defer s.postgreSQL.Close()

	// добавить
	s.users = service.NewUsers(db)

	log.Info("server starting...")

	return http.ListenAndServe(s.config.Port, s.router)
}

func (s *APIServer) configureRouter() {
	s.router.Use(withLogging)
	s.router.Post("/redis/incr", nil)
	s.router.Post("/sign/hmacsha512", nil)
	s.router.Post("/postgres/users", nil)
}

func (s *APIServer) configureStore() (*repository.Storage, error) {
	db, err := repository.NewStorage(s.config.PortDB)
	if err != nil {
		return nil, err
	}
	return db, nil
}
