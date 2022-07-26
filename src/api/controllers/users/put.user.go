package users

import (
	"fmt"
	"net/http"
	"root/src/db"
	"root/src/db/models"
	"root/src/util"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	body := models.Users{}

	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failure",
			"message": fmt.Sprintf("%s...", err),
		})
		return
	}

	var user models.Users

	res := db.DB.First(&user, id)
	if res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failure",
			"message": fmt.Sprintf("%s...", res.Error),
		})
		return
	}

	ID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failure",
			"message": fmt.Sprintf("%s...", err),
		})
		return
	}

	user.ID = (uint(ID))
	user.Firstname = body.Firstname
	user.Lastname = body.Lastname
	user.Email = body.Email
	user.Password = util.Hash(body.Password)

	db.DB.Save(&user)

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"user": models.JsonUser{
			ID:        user.ID,
			Firstname: user.Firstname,
			Lastname:  user.Lastname,
			Email:     user.Email,
		},
	})
}
