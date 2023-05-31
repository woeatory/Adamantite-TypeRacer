package server

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	"github.com/woeatory/Adamantite-TypeRacer/middleware"

	authContr "github.com/woeatory/Adamantite-TypeRacer/internal/auth/controller"
	authServ "github.com/woeatory/Adamantite-TypeRacer/internal/auth/service"
	"github.com/woeatory/Adamantite-TypeRacer/internal/repository"
	userContr "github.com/woeatory/Adamantite-TypeRacer/internal/user/controller"
	userServ "github.com/woeatory/Adamantite-TypeRacer/internal/user/service"
	"io"
	"log"
	"net/http"
	"os"
)

const PORT = ":8080"
const ADDRESS = "localhost" + PORT

const (
	UserGroupPath   = "user"
	UserGetAllPath  = "/getAll"
	UserGetByIdPath = "/:userID"
	ChangeUserName  = "/changeUsername"
	DeleteUser      = "/deleteUser"
	AuthGroupPath   = "auth"
	AuthLogin       = "/login"
	AuthSignUp      = "/signup"
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
	repo := repository.NewRepo()
	userService := userServ.NewUserService(repo)
	authService := authServ.NewAuthService(repo)
	// Set up Controllers
	userController := userContr.NewUserController(userService)
	authController := authContr.NewAuthController(authService)

	router := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	userGroup := router.Group(UserGroupPath)
	userGroup.Use(middleware.Authentication())
	{
		userGroup.GET(UserGetAllPath, userController.GetAllUsers)
		userGroup.GET(UserGetByIdPath, userController.GetUserByID)
		userGroup.PUT(ChangeUserName, userController.ChangeUsername)
		userGroup.DELETE(DeleteUser, userController.DeleteUser)
	}
	authGroup := router.Group(AuthGroupPath)
	{
		authGroup.POST(AuthLogin, authController.LogIn)
		authGroup.POST(AuthSignUp, authController.SignUp)
	}
	log.Fatal(http.ListenAndServe(ADDRESS, router))
}
