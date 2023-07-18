package middleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		sessionID := session.Get("user_session")
		if sessionID == nil {
			c.JSON(
				http.StatusUnauthorized, gin.H{
					"error": "unauthorized",
				},
			)
			c.Abort()
		}
	}
}
