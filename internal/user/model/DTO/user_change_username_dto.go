package DTO

type UserChangeUsernameDto struct {
	NewUsername string `json:"newUsername" binding:"required"`
}
