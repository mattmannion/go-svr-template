package auth

import (
	"fmt"
	"net/http"
	"root/src/db"
	"root/src/db/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func PostAuth(c *gin.Context) {
	body := models.Users{}

	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "failure",
			"message": "Please enter a Username and Password",
		})
		return
	}

	session := sessions.Default(c)
	id := session.Get("id")

	user := &models.Users{}

	if body.Username == "" || body.Password == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "failure",
			"message": "Please enter a Username and Password",
		})
		return
	}

	db.DB.Find(&user, &models.Users{Username: body.Username, Password: body.Password})
	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "failure",
			"message": "Username or Password not found",
		})
		return
	}

	if id != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": fmt.Sprintf("%v is already logged in", user.Username),
		})
		return
	}

	id = (uuid.New()).String()

	session.Set("id", id)

	session.Set("username", user.Username)

	session.Save()

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": fmt.Sprintf("%v logged in", user.Username),
	})
}
