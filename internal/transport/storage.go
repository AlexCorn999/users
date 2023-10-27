package transport

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/AlexCorn999/users/internal/domain"
)

// AddValue добавляет значение в базу Redis.
func (s *APIServer) AddValue(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		logError("addValue", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var input domain.RedisInput
	if err := json.Unmarshal(data, &input); err != nil {
		logError("addValue", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := input.Validate(); err != nil {
		logError("addValue", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	value, err := s.redis.AddValue(r.Context(), &input)
	if err != nil {
		logError("addValue", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	redisOutput := domain.RedisOutput{
		Value: value,
	}

	result, err := json.Marshal(redisOutput)
	if err != nil {
		logError("addValue", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Write(result)
	w.WriteHeader(http.StatusOK)
}
