package users

import (
	"fmt"
	"net/http"
	"root/src/db"
	"root/src/db/models"

	"github.com/gin-gonic/gin"
)

func PostUser(c *gin.Context) {
	body := models.Users{}

	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "failure",
			"message": fmt.Sprintf("%s...", err),
		})
		c.Abort()
	}

	var user models.Users

	user.ID = body.ID
	user.FirstName = body.FirstName
	user.LastName = body.LastName

	res := db.DB.Create(&user)
	if res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "failure",
			"message": fmt.Sprintf("%s...", res.Error),
		})
		c.Abort()
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"user":   user,
	})
}
