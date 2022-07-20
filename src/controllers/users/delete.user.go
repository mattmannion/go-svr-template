package users

import (
	"mm/pkg/src/db/models"
	"mm/pkg/src/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c controller) DeleteUser(g *gin.Context) {
	id := g.Param("id")

	var user models.Users

	if res := c.DB.First(&user, id); res.Error != nil {
		g.JSON(http.StatusBadRequest, util.JSON_MSG{
			Status:  "Failure",
			Message: "User not found...",
		})
		return
	}

	c.DB.Delete(&user)

	g.JSON(http.StatusOK, util.JSON_MSG{
		Status:  "Success",
		Message: "User successful",
	})
}
