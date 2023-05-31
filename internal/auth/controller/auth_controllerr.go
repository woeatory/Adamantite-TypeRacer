package controller

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/woeatory/Adamantite-TypeRacer/internal/auth/models/DTO"
	"github.com/woeatory/Adamantite-TypeRacer/internal/auth/service"
	"net/http"
)

type AuthController struct {
	authService *service.AuthService
}

func NewAuthController(authService *service.AuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}
func (AuthController *AuthController) LogIn(c *gin.Context) {
	var input DTO.UserDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	session := sessions.Default(c)
	userID, err := AuthController.authService.LogIn(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	session.Set("id", userID)
	err = session.Save()
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "successfully Logged In"})
}
func (AuthController *AuthController) SignUp(c *gin.Context) {
	var input DTO.UserDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, err := AuthController.authService.SignUp(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	session := sessions.Default(c)
	session.Set("id", userID)
	err = session.Save()
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "successfully Sign Up"})
}
