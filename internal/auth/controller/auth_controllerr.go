package controller

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/woeatory/Adamantite-TypeRacer/internal/auth/models/DTO"
	"github.com/woeatory/Adamantite-TypeRacer/internal/auth/service"
	"net/http"
)

const MaxAge = 604800 // 1 week

func saveSession(c *gin.Context, userID string) error {
	session := sessions.Default(c)
	session.Options(
		sessions.Options{
			MaxAge:   MaxAge,
			HttpOnly: true,
			Path:     "/",
		},
	)
	session.Set("user_id", userID)
	session.Set("user_session", "token")
	err := session.Save()
	if err != nil {
		return err
	}
	return nil
}

type AuthController struct {
	authService service.Authenticator
}

func NewAuthController(authService service.Authenticator) *AuthController {
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
	userID, err := AuthController.authService.LogIn(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = saveSession(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully Logged In"})
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
	err = saveSession(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully Signed Up"})
}
