package DTO

type ChangeUsernameDTO struct {
	NewUsername string `json:"newUsername" binding:"required"`
}
