package users

import (
	"core/src/db"
	"core/src/db/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(g *gin.Context) {

	var users []models.Users

	res := db.DB.Order("id").Find(&users)
	if res.Error != nil {
		g.JSON(http.StatusNotFound, models.JSON_MSG{
			Status:  "failure",
			Message: fmt.Sprintf("%s...", res.Error),
		})
		return
	}

	type JSON struct {
		Status string         `json:"status"`
		Users  []models.Users `json:"users"`
	}

	g.JSON(http.StatusOK, JSON{
		Status: "success",
		Users:  users,
	})
}
