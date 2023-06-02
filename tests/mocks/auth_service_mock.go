package mocks

import (
	"github.com/stretchr/testify/mock"
	"github.com/woeatory/Adamantite-TypeRacer/internal/auth/models/DTO"
)

type MockAuthService struct {
	mock.Mock
}

func (m *MockAuthService) LogIn(dto DTO.UserDTO) (string, error) {
	args := m.Called(dto)
	return args.String(0), args.Error(1)
}

func (m *MockAuthService) SignUp(dto DTO.UserDTO) (string, error) {
	args := m.Called(dto)
	return args.String(0), args.Error(1)
}
