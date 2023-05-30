package service

import "github.com/woeatory/Adamantite-TypeRacer/internal/auth/models/DTO"

type AuthService struct {
}

func NewAuthService() *AuthService {
	return &AuthService{}
}
func (authService *AuthService) LogIn(dto DTO.UserDTO) error {
	// todo
	return nil
}
func (authService *AuthService) SignUp(dto DTO.UserDTO) error {
	// todo
	return nil
}
