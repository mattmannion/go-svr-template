package users

import (
	"fmt"
	"mm/pkg/src/db/models"
	"mm/pkg/src/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (c controller) UpdateUser(g *gin.Context) {
	id := g.Param("id")
	body := models.Users{}

	err := g.BindJSON(&body)
	if err != nil {
		g.JSON(http.StatusBadRequest, util.JSON_MSG{
			Status:  "failure",
			Message: fmt.Sprintf("%s...", err),
		})
		return
	}

	var user models.Users

	res := c.DB.First(&user, id)
	if res.Error != nil {
		g.JSON(http.StatusBadRequest, util.JSON_MSG{
			Status:  "failure",
			Message: fmt.Sprintf("%s...", res.Error),
		})
		return
	}

	ID, str_err := strconv.Atoi(id)
	if str_err != nil {
		g.JSON(http.StatusBadRequest, util.JSON_MSG{
			Status:  "failure",
			Message: fmt.Sprintf("%s...", str_err),
		})
		return
	}

	user.ID = (uint(ID))
	user.FirstName = body.FirstName
	user.LastName = body.LastName

	c.DB.Save(&user)

	g.JSON(http.StatusOK, JSON{
		Status: "success",
		User:   user,
	})
}
