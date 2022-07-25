package users

import (
	"fmt"
	"mm/pkg/src/db/models"
	"mm/pkg/src/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (db controller) DeleteUser(g *gin.Context) {
	id := g.Param("id")

	var user models.Users

	res := db.DB.First(&user, id)
	if res.Error != nil {
		g.JSON(http.StatusBadRequest, util.JSON_MSG{
			Status:  "Failure",
			Message: fmt.Sprintf("%s", res.Error),
		})
		return
	}

	db.DB.Delete(&user)

	g.JSON(http.StatusOK, util.JSON_MSG{
		Status:  "Success",
		Message: "User " + id + " deleted",
	})
}
