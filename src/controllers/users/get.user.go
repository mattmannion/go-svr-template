package users

import (
	"fmt"
	"mm/pkg/src/db/models"
	"mm/pkg/src/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c controller) GetUser(g *gin.Context) {
	id := g.Param("id")

	var user models.Users

	res := c.DB.First(&user, id)
	if res.Error != nil {
		g.JSON(http.StatusNotFound, util.JSON_MSG{
			Status:  "failure",
			Message: fmt.Sprintf("%s...", res.Error),
		})
		return
	}

	g.JSON(http.StatusOK, JSON{
		Status: "success",
		User:   user,
	})
}
