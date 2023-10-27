package transport

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/AlexCorn999/users/internal/domain"
	"github.com/AlexCorn999/users/internal/repository"
)

// CreateUser добавляет пользователя в базу postgreSQL.
func (s *APIServer) CreateUser(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		logError("newUser", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var usr domain.User
	if err := json.Unmarshal(data, &usr); err != nil {
		logError("newUser", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := usr.Validate(); err != nil {
		logError("newUser", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := s.users.CreateUser(r.Context(), &usr)
	if err != nil {
		if errors.Is(err, repository.ErrDuplicate) {
			w.WriteHeader(http.StatusConflict)
			return
		} else {
			logError("newUser", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	userOutput := domain.UserOutput{
		ID: id,
	}

	result, err := json.Marshal(userOutput)
	if err != nil {
		logError("newUser", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Write(result)
	w.WriteHeader(http.StatusOK)
}
