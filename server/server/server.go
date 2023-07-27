package server

import (
	"context"
	"io"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/woeatory/Adamantite-TypeRacer/middleware"

	authContr "github.com/woeatory/Adamantite-TypeRacer/internal/auth/controller"
	authServ "github.com/woeatory/Adamantite-TypeRacer/internal/auth/service"
	"github.com/woeatory/Adamantite-TypeRacer/internal/repository"
	userContr "github.com/woeatory/Adamantite-TypeRacer/internal/user/controller"
	userServ "github.com/woeatory/Adamantite-TypeRacer/internal/user/service"
	scoreContr "github.com/woeatory/Adamantite-TypeRacer/internal/user_scores/controller"
	scoreServ "github.com/woeatory/Adamantite-TypeRacer/internal/user_scores/service"

	"log"
	"net/http"
)

const (
	UserGroupPath           = "user"
	UserGetAllPath          = "/getAll"
	UserGetByIdPath         = "/:userID"
	ChangeUserName          = "/changeUsername"
	DeleteUser              = "/deleteUser"
	AuthGroupPath           = "auth"
	AuthLogin               = "/login"
	AuthSignUp              = "/signup"
	AuthLogOut              = "/logout"
	ScoreGroupPath          = "score"
	NewScoreRecord          = "/newScoreRecord"
	GetAllUsersScoreRecords = "/getAllUsersScoreRecords"
	DeleteScoreRecord       = "/deleteScoreRecord"
)

func SetUpAndBoot() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading env: %s\n", err)
	}
	port := os.Getenv("PORT")

	c, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Set up Services
	repo := repository.NewRepo()
	userService := userServ.NewUserService(repo)
	authService := authServ.NewAuthService(repo)
	scoreService := scoreServ.NewScoreService(repo)
	// Set up Controllers
	userController := userContr.NewUserController(userService)
	authController := authContr.NewAuthController(authService)
	scoreController := scoreContr.NewScoreController(scoreService)

	router := SetUpRouter()
	storageSecret := os.Getenv("STORAGE_KEY")
	if storageSecret == "" {
		log.Fatalln("error getting env STORAGE_KEY")
	}
	store := cookie.NewStore([]byte(storageSecret))
	router.Use(sessions.Sessions("session", store))

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
		authGroup.POST(AuthLogOut, authController.LogOut)
	}

	scoreGroup := router.Group(ScoreGroupPath)
	scoreGroup.Use(middleware.Authentication())
	{
		scoreGroup.POST(NewScoreRecord, scoreController.AddNewScoreRecord)
		scoreGroup.GET(GetAllUsersScoreRecords, scoreController.GetAllUsersRecords)
		scoreGroup.DELETE(DeleteScoreRecord, scoreController.DeleteScoreRecord)
	}

	srv := &http.Server{
		Addr:    port,
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Listen for the interrupt signal.
	<-c.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	log.Println("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	c, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(c); err != nil {
		log.Fatal("server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
	err = repo.CloseRepo()
	if err != nil {
		log.Fatalf("error closing DB: %s\n", err)
	}
}

func SetUpRouter() *gin.Engine {
	// Logging into file and console
	gin.DisableConsoleColor()
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	router := gin.Default()
	return router
}
