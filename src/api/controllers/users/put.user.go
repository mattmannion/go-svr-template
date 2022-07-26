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
	id := c.Query("id")

	if id == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "failure",
			"message": "Please specify a url query of 'id'",
		})
		return
	}

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
	if body.Firstname != "" {
		user.Firstname = body.Firstname
	}

	if body.Lastname != "" {
		user.Lastname = body.Lastname
	}

	if body.Email != "" {
		user.Email = body.Email
	}

	if body.Password != "" {
		user.Password = util.Hash(body.Password)
	}

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
