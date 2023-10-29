package transport

import (
	"bytes"
	"net/http/httptest"
	"testing"

	"github.com/AlexCorn999/users/internal/config"
	"github.com/AlexCorn999/users/internal/domain"
	"gopkg.in/go-playground/assert.v1"
)

func Test_SignHmacSha512(t *testing.T) {

	// Test Server
	server := NewAPIServer(config.NewConfig())
	server.router.Post("/sign/hmacsha512", server.SignHmacSha512)

	testTable := []struct {
		name                string
		inputBody           string
		input               domain.SignHmacSha512
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:                "OK",
			inputBody:           `{"text": "test", "key": "test123"}`,
			expectedStatusCode:  200,
			expectedRequestBody: `b596e24739fd44d42ffd25f26ea367dad3a71f61c8c5fab6b6ee6ceeae5a7170b66445d6eaadfb49e6d4e968a2888726ff522e3bf065c966aa66a24153778382`,
		},
		{
			name:                "Large number",
			inputBody:           `{"text": "secret", "key": "346345756673464567"}`,
			expectedStatusCode:  200,
			expectedRequestBody: `ef1877c671b1c865ba9f788224760f50f09b6b64c1ffea04b50c0b1eaaabc8cba7e3abd8af6c5b467891efa2de89a17787343ca6d5aa4f94828f37193f9aee5a`,
		},
		{
			name:                "nil",
			inputBody:           `{"text": "", "key": "346345756673464567"}`,
			expectedStatusCode:  400,
			expectedRequestBody: ``,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			// Test Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/sign/hmacsha512",
				bytes.NewBufferString(testCase.inputBody))

			// Perform Request
			server.router.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}
}
