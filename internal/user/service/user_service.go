package service

import "github.com/woeatory/Adamantite-TypeRacer/internal/user/model"

type UserService struct {
	users []model.User
}

func NewUserService() *UserService {
	return &UserService{users: []model.User{}}
}

func (userService *UserService) GetAll() ([]model.User, error) {
	return userService.users, nil
}

func (userService *UserService) GetByID() {

}

func (userService *UserService) ChangeUsername() {

}

func (userService *UserService) DeleteUser() {

}
