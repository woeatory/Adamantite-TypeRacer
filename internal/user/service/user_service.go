package service

import (
	"github.com/woeatory/Adamantite-TypeRacer/internal/repository"
	"github.com/woeatory/Adamantite-TypeRacer/internal/user/model"
)

type UserService struct {
	repo repository.Repo
}

func NewUserService(repo *repository.Repo) *UserService {
	return &UserService{repo: *repo}
}

func (userService *UserService) GetAll() ([]model.User, error) {
	rows, err := userService.repo.DB.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
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

func (userService *UserService) GetByID() {

}

func (userService *UserService) ChangeUsername() {

}

func (userService *UserService) DeleteUser() {

}
