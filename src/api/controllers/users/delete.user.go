package users

import (
	"fmt"
	"net/http"
	"root/src/db"
	"root/src/db/models"

	"github.com/gin-gonic/gin"
)

func DeleteUser(g *gin.Context) {
	id := g.Param("id")

	var user models.Users

	res := db.DB.First(&user, id)
	if res.Error != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"status":  "Failure",
			"message": fmt.Sprintf("%s", res.Error),
		})
		return
	}

	db.DB.Delete(&user)

	// g.JSON(http.StatusOK, models.JSON_MSG{
	// 	Status:  "Success",
	// 	Message: "User " + id + " deleted",
	// })

	g.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "User " + id + " deleted",
	})
}
