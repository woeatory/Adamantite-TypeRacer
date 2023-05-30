package server

import (
	"github.com/gin-gonic/gin"
	authController "github.com/woeatory/Adamantite-TypeRacer/internal/auth/controller"
	authServ "github.com/woeatory/Adamantite-TypeRacer/internal/auth/service"
	userController "github.com/woeatory/Adamantite-TypeRacer/internal/user/controller"
	userServ "github.com/woeatory/Adamantite-TypeRacer/internal/user/service"
	"io"
	"log"
	"net/http"
	"os"
)

const PORT = ":8080"
const ADDRESS = "localhost" + PORT

const (
	UserGroupPath  = "user"
	UserGetAllPath = "/getAll"
	AuthGroupPath  = "auth"
	AuthLogin      = "/login"
	AuthSignUp     = "/signup"
)

func SetUpAndBoot() {
	// Disable Console Color, you don't need console color when writing the logs to file.
	gin.DisableConsoleColor()

	// Logging to a file.
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	// Use the following code if you need to write the logs to file and console at the same time.
	// gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// Set up Services
	userService := userServ.NewUserService()
	authService := authServ.NewAuthService()
	// Set up Controllers
	userController := userController.NewUserController(userService)
	authController := authController.NewAuthController(authService)
	router := gin.Default()
	userGroup := router.Group(UserGroupPath)
	{
		userGroup.GET(UserGetAllPath, userController.GetAllUsers)
	}
	authGroup := router.Group(AuthGroupPath)
	{
		authGroup.POST(AuthLogin, authController.LogIn)
		authGroup.POST(AuthSignUp, authController.SignUp)
	}
	log.Fatal(http.ListenAndServe(ADDRESS, router))
}

func userTestHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "user test"})
}
