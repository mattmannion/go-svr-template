package users

import (
	"mm/pkg/src/db/models"
	"mm/pkg/src/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c controller) PostUser(g *gin.Context) {
	json := models.UsersJSON{}

	if err := g.BindJSON(&json); err != nil {
		g.JSON(http.StatusNotFound, util.JSON_MSG{
			Status:  "failure",
			Message: "Could not bind JSON...",
		})
		return
	}

	var user models.Users

	user.ID = json.ID
	user.FirstName = json.FirstName
	user.LastName = json.LastName

	if res := c.DB.Create(&user); res.Error != nil {
		g.JSON(http.StatusNotFound, util.JSON_MSG{
			Status:  "failure",
			Message: "Could not create user...",
		})
		return
	}

	g.JSON(http.StatusCreated, JSON{
		Status: "success",
		User:   user,
	})
}
