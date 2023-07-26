package tests

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/woeatory/Adamantite-TypeRacer/middleware"
)

func setUpRouterWithAuth(userID string) *gin.Engine {
	// setup router with needed middleware and etc
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("testsesesion", store))
	router.Use(
		func(c *gin.Context) {
			session := sessions.Default(c)
			session.Set("session_token", "token")
			session.Set("user_id", userID)
			err := session.Save()
			if err != nil {
				return
			}
			c.Next()
		},
	)
	router.Use(middleware.Authentication())
	return router
}

func setUpRouterNotAuth() *gin.Engine {
	// setup router with needed middleware and etc
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("testsesesion", store))
	router.Use(middleware.Authentication())
	return router
}
