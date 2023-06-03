package service

import (
	"github.com/woeatory/Adamantite-TypeRacer/internal/repository"
	"github.com/woeatory/Adamantite-TypeRacer/internal/user/model"
)

type UserServiceInterface interface {
	GetAll() ([]model.User, error)
	GetByID(userID string) (model.User, error)
	ChangeUsername(newUsername, userID string) error
	DeleteUser(userID string) error
}

type UserService struct {
	repo repository.Repo
}

func NewUserService(repo *repository.Repo) *UserService {
	return &UserService{repo: *repo}
}

func (userService *UserService) GetAll() ([]model.User, error) {
	rows, err := userService.repo.DB.Query("SELECT * FROM users")
	if err != nil {
		return []model.User{}, err
	}
	defer rows.Close()
	var users []model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.UserID, &user.Username, &user.PasswordHash, &user.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (userService *UserService) GetByID(userID string) (model.User, error) {
	query := "SELECT user_id, username, created_at FROM users WHERE user_id = $1"
	stmt, err := userService.repo.DB.Prepare(query)
	if err != nil {
		return model.User{}, err
	}
	var user model.User
	err = stmt.QueryRow(userID).Scan(&user.UserID, &user.Username, &user.CreatedAt)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (userService *UserService) ChangeUsername(newUsername, userID string) error {
	query := "UPDATE users SET username = $1 WHERE user_id = $2"
	stmt, err := userService.repo.DB.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(newUsername, userID)
	if err != nil {
		return err
	}
	return nil
}

func (userService *UserService) DeleteUser(userID string) error {
	query := "DELETE FROM users WHERE user_id = $1"
	stmt, err := userService.repo.DB.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(userID)
	if err != nil {
		return err
	}
	return nil
}
