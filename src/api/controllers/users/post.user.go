package users

import (
	"fmt"
	"net/http"
	"root/src/db"
	"root/src/db/models"

	"github.com/gin-gonic/gin"
)

func PostUser(g *gin.Context) {
	body := models.Users{}

	err := g.BindJSON(&body)
	if err != nil {
		g.JSON(http.StatusNotFound, gin.H{
			"status":  "failure",
			"message": fmt.Sprintf("%s...", err),
		})
		return
	}

	var user models.Users

	user.ID = body.ID
	user.FirstName = body.FirstName
	user.LastName = body.LastName

	res := db.DB.Create(&user)
	if res.Error != nil {
		g.JSON(http.StatusNotFound, gin.H{
			"status":  "failure",
			"message": fmt.Sprintf("%s...", res.Error),
		})
		return
	}

	g.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"user":   user,
	})
}
