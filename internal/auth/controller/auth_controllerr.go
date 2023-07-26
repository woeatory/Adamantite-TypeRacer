package controller

import (
	"encoding/hex"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/woeatory/Adamantite-TypeRacer/internal/auth/models/DTO"
	"github.com/woeatory/Adamantite-TypeRacer/internal/auth/service"
	"log"
	"math/rand"
	"net/http"
)

const MaxAge = 604800 // 1 week

const TOKEN_LENGTH = 80
const ALPHA_UPPER = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const ALPHA_LOWER = "abcdefghijklmnopqrstuvwxyz"
const ALPHA = ALPHA_UPPER + ALPHA_LOWER
const DIGIT = "0123456789"
const ALPHA_DIGIT = ALPHA + DIGIT

func generateToken() (string, error) {
	b := make([]byte, TOKEN_LENGTH)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

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
	token, err := generateToken()
	if err != nil {
		log.Println("error generating token")
		return err
	}
	session.Set("session_token", token)
	err = session.Save()
	if err != nil {
		log.Println("error saving token")
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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

func (AuthController *AuthController) LogOut(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Options(
		sessions.Options{
			Path:   "/",
			MaxAge: -1,
		},
	)
	err := session.Save()
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully Logged Out"})
	c.Redirect(http.StatusMovedPermanently, "/")
}
