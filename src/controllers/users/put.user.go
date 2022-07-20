package users

import (
	"mm/pkg/src/db/models"
	"mm/pkg/src/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (c controller) UpdateUser(g *gin.Context) {
	id := g.Param("id")
	json := models.UsersJSON{}

	if err := g.BindJSON(&json); err != nil {
		g.JSON(http.StatusBadRequest, util.JSON_MSG{
			Status:  "failure",
			Message: "Could not bind JSON",
		})
		return
	}

	var user models.Users

	res := c.DB.First(&user, id)
	if res.Error != nil {
		g.JSON(http.StatusBadRequest, util.JSON_MSG{
			Status:  "failure",
			Message: "Could not find user...",
		})
		return
	}

	ID, str_err := strconv.Atoi(id)

	if str_err != nil {
		g.JSON(http.StatusBadRequest, util.JSON_MSG{
			Status:  "failure",
			Message: "Could not convert id...",
		})
		return
	}

	user.ID = (uint(ID))
	user.FirstName = json.FirstName
	user.LastName = json.LastName

	c.DB.Save(&user)

	g.JSON(http.StatusOK, JSON{
		Status: "success",
		User:   user,
	})
}
