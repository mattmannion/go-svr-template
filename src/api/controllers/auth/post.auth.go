package auth

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func PostAuth(c *gin.Context) {
	session := sessions.Default(c)
	id := session.Get("id")

	if id != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "User already signed in...",
			"user":    session.Get("email"),
		})
		return
	}

	id = uuid.New()

	session.Set("id", id)

	session.Set("email", "test@gmail.com")

	session.Save()

	c.JSON(http.StatusOK, gin.H{
		"message": "User Sign In successfully",
	})
}
