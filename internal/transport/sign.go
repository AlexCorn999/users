package transport

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/AlexCorn999/users/internal/domain"
)

// SignHmacSha512 возвращает HMAC-SHA512 подпись значения из "text" по ключу "key" в виде hex строки.
func (s *APIServer) SignHmacSha512(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		logError("signHmacSha512", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var signData domain.SignHmacSha512
	if err := json.Unmarshal(data, &signData); err != nil {
		logError("signHmacSha512", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := signData.Validate(); err != nil {
		logError("signHmacSha512", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := s.sign.SignHmacSha512(&signData)
	if err != nil {
		logError("signHmacSha512", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Write([]byte(result))
	w.WriteHeader(http.StatusOK)
}
