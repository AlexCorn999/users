package transport

import (
	"bytes"
	"net/http/httptest"
	"testing"

	"github.com/AlexCorn999/users/internal/config"
	"github.com/AlexCorn999/users/internal/domain"
	"github.com/AlexCorn999/users/internal/service"
	mock_service "github.com/AlexCorn999/users/internal/service/mocks"
	"github.com/golang/mock/gomock"
	"gopkg.in/go-playground/assert.v1"
)

func TestHandler_AddValue(t *testing.T) {
	type mockBehavior func(s *mock_service.MockStorageRepository, input domain.RedisInput)

	testTable := []struct {
		name                string
		inputBody           string
		input               domain.RedisInput
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "OK",
			inputBody: `{"key": "age", "value": 19 }`,
			input: domain.RedisInput{
				Key:   "age",
				Value: 19,
			},
			mockBehavior: func(s *mock_service.MockStorageRepository, input domain.RedisInput) {
				s.EXPECT().AddValue(gomock.Any(), input).Return(20, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"value":20}`,
		},
		{
			name:      "Large number",
			inputBody: `{"key": "age", "value": 111111111111111111 }`,
			input: domain.RedisInput{
				Key:   "age",
				Value: 111111111111111111,
			},
			mockBehavior: func(s *mock_service.MockStorageRepository, input domain.RedisInput) {
				s.EXPECT().AddValue(gomock.Any(), input).Return(111111111111111112, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"value":111111111111111112}`,
		},
		{
			name:      "nil",
			inputBody: `{"key": "", "value": 111111111111111111 }`,
			mockBehavior: func(s *mock_service.MockStorageRepository, input domain.RedisInput) {
			},
			expectedStatusCode:  400,
			expectedRequestBody: ``,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			store := mock_service.NewMockStorageRepository(c)
			testCase.mockBehavior(store, testCase.input)

			storage := service.Storage{
				Repo: store,
			}

			// Test server
			server := NewAPIServer(config.NewConfig())
			server.storage = &storage
			server.router.Post("/redis/incr", server.AddValue)

			// Test Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/redis/incr",
				bytes.NewBufferString(testCase.inputBody))

			// Perform Request
			server.router.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}
}
