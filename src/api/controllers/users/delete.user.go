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
		g.JSON(http.StatusBadRequest, models.JSON_MSG{
			Status:  "Failure",
			Message: fmt.Sprintf("%s", res.Error),
		})
		return
	}

	db.DB.Delete(&user)

	g.JSON(http.StatusOK, models.JSON_MSG{
		Status:  "Success",
		Message: "User " + id + " deleted",
	})
}
