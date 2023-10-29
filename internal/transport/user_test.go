package transport

import (
	"bytes"
	"net/http/httptest"
	"testing"

	"github.com/AlexCorn999/users/internal/config"
	"github.com/AlexCorn999/users/internal/domain"
	"github.com/AlexCorn999/users/internal/repository"
	"github.com/AlexCorn999/users/internal/service"
	mock_service "github.com/AlexCorn999/users/internal/service/mocks"
	"github.com/golang/mock/gomock"
	"gopkg.in/go-playground/assert.v1"
)

func TestHandler_CreateUser(t *testing.T) {
	type mockBehavior func(s *mock_service.MockUserRepository, input domain.User)

	testTable := []struct {
		name                string
		inputBody           string
		input               domain.User
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "user1",
			inputBody: `{"name": "Alex", "age": 21}`,
			input: domain.User{
				Login: "Alex",
				Age:   21,
			},
			mockBehavior: func(s *mock_service.MockUserRepository, input domain.User) {
				s.EXPECT().CreateUser(gomock.Any(), input).Return(1, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"id":1}`,
		},
		{
			name:      "user2",
			inputBody: `{"name": "Kate", "age": 18}`,
			input: domain.User{
				Login: "Kate",
				Age:   18,
			},
			mockBehavior: func(s *mock_service.MockUserRepository, input domain.User) {
				s.EXPECT().CreateUser(gomock.Any(), input).Return(2, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"id":2}`,
		},
		{
			name:      "Login already in use",
			inputBody: `{"name": "Kate", "age": 18}`,
			input: domain.User{
				Login: "Kate",
				Age:   18,
			},
			mockBehavior: func(s *mock_service.MockUserRepository, input domain.User) {
				s.EXPECT().CreateUser(gomock.Any(), input).Return(0, repository.ErrDuplicate)
			},
			expectedStatusCode:  409,
			expectedRequestBody: ``,
		},
		{
			name:      "nil",
			inputBody: `{"name": "", "age": 18}`,
			mockBehavior: func(s *mock_service.MockUserRepository, input domain.User) {
			},
			expectedStatusCode:  400,
			expectedRequestBody: ``,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			users := mock_service.NewMockUserRepository(c)
			testCase.mockBehavior(users, testCase.input)

			user := service.Users{
				Repo: users,
			}

			// Test server
			server := NewAPIServer(config.NewConfig())
			server.users = &user
			server.router.Post("/postgres/users", server.CreateUser)

			// Test Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/postgres/users",
				bytes.NewBufferString(testCase.inputBody))

			// Perform Request
			server.router.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}
}
