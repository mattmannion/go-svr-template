package users

import (
	"_/src/db"
	"_/src/db/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	body := models.Users{}

	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSON_MSG{
			Status:  "failure",
			Message: fmt.Sprintf("%s...", err),
		})
		return
	}

	var user models.Users

	res := db.DB.First(&user, id)
	if res.Error != nil {
		c.JSON(http.StatusBadRequest, models.JSON_MSG{
			Status:  "failure",
			Message: fmt.Sprintf("%s...", res.Error),
		})
		return
	}

	ID, str_err := strconv.Atoi(id)
	if str_err != nil {
		c.JSON(http.StatusBadRequest, models.JSON_MSG{
			Status:  "failure",
			Message: fmt.Sprintf("%s...", str_err),
		})
		return
	}

	user.ID = (uint(ID))
	user.FirstName = body.FirstName
	user.LastName = body.LastName

	db.DB.Save(&user)

	c.JSON(http.StatusOK, models.JSON{
		Status: "success",
		User:   user,
	})
}
