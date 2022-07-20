package users

import (
	"mm/pkg/src/db/models"
	"mm/pkg/src/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c controller) GetUsers(g *gin.Context) {

	var users []models.Users

	if res := c.DB.Find(&users); res.Error != nil {
		g.JSON(http.StatusNotFound, util.JSON_MSG{
			Status:  "failure",
			Message: "no users found",
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
