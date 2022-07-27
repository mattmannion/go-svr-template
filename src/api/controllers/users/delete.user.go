package users

import (
	"fmt"
	"net/http"
	"root/src/db"
	"root/src/db/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func DeleteUser(c *gin.Context) {
	session := sessions.Default(c)
	username := fmt.Sprint(session.Get("username"))

	var user models.Users

	res := db.DB.First(&user, &models.Users{Username: username})
	if res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Failure",
			"message": fmt.Sprint(res.Error),
		})
		return
	}

	db.DB.Delete(&user)

	session.Clear()

	session.Save()

	c.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "User " + username + " deleted",
	})
}
