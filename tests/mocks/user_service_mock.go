package mocks

import (
	"github.com/stretchr/testify/mock"
	"github.com/woeatory/Adamantite-TypeRacer/internal/user/model"
)

// MockUserService is a mock implementation of the UserServiceInterface
type MockUserService struct {
	mock.Mock
}

// GetAll is a mock implementation for the GetAll method
func (m *MockUserService) GetAll() ([]model.User, error) {
	args := m.Called()
	return args.Get(0).([]model.User), args.Error(1)
}

// GetByID is a mock implementation for the GetByID method
func (m *MockUserService) GetByID(userID string) (model.User, error) {
	args := m.Called(userID)
	return args.Get(0).(model.User), args.Error(1)
}

// ChangeUsername is a mock implementation for the ChangeUsername method
func (m *MockUserService) ChangeUsername(newUsername, userID string) error {
	args := m.Called(newUsername, userID)
	return args.Error(0)
}

// DeleteUser is a mock implementation for the DeleteUser method
func (m *MockUserService) DeleteUser(userID string) error {
	args := m.Called(userID)
	return args.Error(0)
}
