package auth

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func DeleteAuth(c *gin.Context) {
	session := sessions.Default(c)

	id := session.Get("id")
	username := session.Get("username")

	if id == nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "No User logged in",
		})
		return
	}

	session.Clear()

	session.Save()

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": fmt.Sprintf("%v logged out", username),
	})
}
