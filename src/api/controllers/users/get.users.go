package users

import (
	"fmt"
	"net/http"
	"root/src/db"
	"root/src/db/models"

	"github.com/gin-gonic/gin"
)

func GetUsers(g *gin.Context) {

	var users []models.Users

	res := db.DB.Order("id").Find(&users)
	if res.Error != nil {
		g.JSON(http.StatusNotFound, gin.H{
			"status":  "failure",
			"message": fmt.Sprintf("%s...", res.Error),
		})
		return
	}

	g.JSON(http.StatusOK, gin.H{
		"status": "success",
		"users":  users,
	})
}
