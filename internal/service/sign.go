package service

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"

	"github.com/AlexCorn999/users/internal/domain"
)

type Sign struct {
}

func NewSign() *Sign {
	return &Sign{}
}

func (s *Sign) SignHmacSha512(inputData *domain.SignHmacSha512) (string, error) {
	h := hmac.New(sha512.New, []byte(inputData.Key))
	h.Write([]byte(inputData.Text))
	signature := h.Sum(nil)
	return hex.EncodeToString(signature), nil
}
