package middleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	session := sessions.Default(c)
	sessionID := session.Get("id")

	if sessionID == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "failure",
			"message": "Please login",
		})

		c.Abort()
	}
	c.Next()
}
