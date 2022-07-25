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
		g.JSON(http.StatusNotFound, models.JSON_MSG{
			Status:  "failure",
			Message: fmt.Sprintf("%s...", res.Error),
		})
		return
	}

	g.JSON(http.StatusOK, models.JSON{
		Status: "success",
		User:   user,
	})
}
