package server

import (
	"context"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"os/signal"
	"syscall"
	"time"

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

const PORT = ":8080"
const ADDRESS = "localhost" + PORT

const (
	UserGroupPath     = "user"
	UserGetAllPath    = "/getAll"
	UserGetByIdPath   = "/:userID"
	ChangeUserName    = "/changeUsername"
	DeleteUser        = "/deleteUser"
	AuthGroupPath     = "auth"
	AuthLogin         = "/login"
	AuthSignUp        = "/signup"
	ScoreGroupPath    = "score"
	NewScoreRecord    = "/newScoreRecord"
	DeleteScoreRecord = "/deleteScoreRecord"
)

func SetUpAndBoot() {

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
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

	scoreGroup := router.Group(ScoreGroupPath)
	scoreGroup.Use(middleware.Authentication())
	{
		scoreGroup.POST(NewScoreRecord, scoreController.AddNewScore)
		scoreGroup.DELETE(DeleteScoreRecord, scoreController.DeleteScoreRecord)
	}

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Listen for the interrupt signal.
	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	log.Println("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}
