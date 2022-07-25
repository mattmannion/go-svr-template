package users

import (
	"fmt"
	"net/http"
	"root/src/db"
	"root/src/db/models"

	"github.com/gin-gonic/gin"
)

func GetUser(g *gin.Context) {
	id := g.Param("id")

	var user models.Users

	res := db.DB.First(&user, id)
	if res.Error != nil {
		g.JSON(http.StatusNotFound, gin.H{
			"status":  "failure",
			"message": fmt.Sprintf("%s...", res.Error),
		})
		return
	}

	g.JSON(http.StatusOK, gin.H{
		"status": "success",
		"user":   user,
	})
}
