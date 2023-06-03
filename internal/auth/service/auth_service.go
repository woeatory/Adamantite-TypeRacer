package service

import (
	"github.com/google/uuid"
	"github.com/woeatory/Adamantite-TypeRacer/internal/auth/models/DTO"
	"github.com/woeatory/Adamantite-TypeRacer/internal/repository"
	"github.com/woeatory/Adamantite-TypeRacer/internal/user/model"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Authenticator interface {
	LogIn(dto DTO.UserDTO) (string, error)
	SignUp(dto DTO.UserDTO) (string, error)
}

type AuthService struct {
	repo *repository.Repo
}

func NewAuthService(repo *repository.Repo) *AuthService {
	return &AuthService{repo: repo}
}
func (authService *AuthService) LogIn(dto DTO.UserDTO) (string, error) {
	query := "SELECT user_id, password_hash FROM users WHERE username = $1"
	stmt, err := authService.repo.DB.Prepare(query)
	if err != nil {
		return "", err
	}
	defer stmt.Close()
	var userID string
	var hash string
	err = stmt.QueryRow(dto.Username).Scan(&userID, &hash)
	if err != nil {
		return "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(dto.Password))
	if err != nil {
		return "", err
	}
	return userID, nil
}
func (authService *AuthService) SignUp(dto DTO.UserDTO) (string, error) {
	userID, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	user := model.User{
		UserID:       userID,
		Username:     dto.Username,
		PasswordHash: string(hashedPass),
		CreatedAt:    time.Now(),
	}
	query := "INSERT INTO users (user_id, username, password_hash, created_at) VALUES ($1, $2, $3, $4)"
	_, err = authService.repo.DB.Exec(query, user.UserID.String(), user.Username, user.PasswordHash, user.CreatedAt)
	if err != nil {
		return "", err
	}
	return userID.String(), nil
}
