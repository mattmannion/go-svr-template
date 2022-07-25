package users

import (
	"fmt"
	"net/http"
	"root/src/db"
	"root/src/db/models"

	"github.com/gin-gonic/gin"
)

func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	var user models.Users

	res := db.DB.First(&user, id)
	if res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Failure",
			"message": fmt.Sprintf("%s", res.Error),
		})
		c.Abort()
	}

	db.DB.Delete(&user)

	c.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "User " + id + " deleted",
	})
}
