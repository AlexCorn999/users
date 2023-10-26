package transport

import (
	"context"
	"net/http"
	"time"

	"github.com/AlexCorn999/users/internal/config"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"
)

type APIServer struct {
	config *config.Config
	router *chi.Mux
	//store  *repository.UserStore
	//users  *service.Users
}

func NewAPIServer(config *config.Config) *APIServer {
	return &APIServer{
		config: config,
		router: chi.NewRouter(),
	}
}

func (s *APIServer) Start(ctx context.Context) error {
	s.configureRouter()

	/*
		if err := s.configureStore(); err != nil {
			return err
		}*/

	//s.users = service.NewUsers(s.store)

	log.Info("server starting...")

	go func() {
		if err := http.ListenAndServe(s.config.Port, s.router); err != nil {
			log.Fatal(err)
		}
	}()
	<-ctx.Done()
	log.Info("shutting down server gracefully")

	/*
		b, err := json.Marshal(&s.store)
		if err != nil {
			return err
		}

		if err = os.WriteFile(s.config.FileStore, b, fs.ModePerm); err != nil {
			return err
		}*/

	return nil
}

func (s *APIServer) configureRouter() {
	s.router.Use(middleware.RequestID)
	s.router.Use(middleware.RealIP)
	s.router.Use(middleware.Logger)
	s.router.Use(middleware.Timeout(60 * time.Second))

	s.router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(time.Now().String()))
	})
	s.router.Post("/redis/incr", nil)
	s.router.Post("/sign/hmacsha512", nil)
	s.router.Post("/postgres/users", nil)
}

/*
func (s *APIServer) configureStore() error {
	store := repository.NewUserStore()
	s.store = store
	return nil
}*/
