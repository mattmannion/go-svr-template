package users

import (
	"fmt"
	"net/http"
	"root/src/db"
	"root/src/db/models"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {

	var users []models.Users

	res := db.DB.Select("id", "firstname", "lastname", "email", "username").Order("id").Find(&users)
	if res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "failure",
			"message": fmt.Sprintf("%s...", res.Error),
		})
		return
	}

	var jsonUsers []models.JsonUser

	for _, user := range users {
		jsonUsers = append(jsonUsers, models.JsonUser{
			ID:        user.ID,
			Firstname: user.Firstname,
			Lastname:  user.Lastname,
			Email:     user.Email,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"users":  jsonUsers,
	})
}
