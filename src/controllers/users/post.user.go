package users

import (
	"fmt"
	"mm/pkg/src/db/models"
	"mm/pkg/src/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c controller) PostUser(g *gin.Context) {
	body := models.Users{}

	err := g.BindJSON(&body)
	if err != nil {
		g.JSON(http.StatusNotFound, util.JSON_MSG{
			Status:  "failure",
			Message: fmt.Sprintf("%s...", err),
		})
		return
	}

	var user models.Users

	user.ID = body.ID
	user.FirstName = body.FirstName
	user.LastName = body.LastName

	res := c.DB.Create(&user)
	if res.Error != nil {
		g.JSON(http.StatusNotFound, util.JSON_MSG{
			Status:  "failure",
			Message: fmt.Sprintf("%s...", res.Error),
		})
		return
	}

	g.JSON(http.StatusCreated, JSON{
		Status: "success",
		User:   user,
	})
}
