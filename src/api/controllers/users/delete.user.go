package users

import (
	"fmt"
	"net/http"
	"root/src/db"
	"root/src/db/models"

	"github.com/gin-gonic/gin"
)

func DeleteUser(c *gin.Context) {
	id := c.Query("id")

	if id == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "failure",
			"message": "Please specify a url query of 'id'",
		})
		return
	}

	var user models.Users

	res := db.DB.First(&user, id)
	if res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Failure",
			"message": fmt.Sprintf("%s", res.Error),
		})
		return
	}

	db.DB.Delete(&user)

	c.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "User " + id + " deleted",
	})
}
