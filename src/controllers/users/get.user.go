package users

import (
	"core/src/db"
	"core/src/db/models"
	"fmt"
	"net/http"

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
