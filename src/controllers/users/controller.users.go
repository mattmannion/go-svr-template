package users

import (
	"mm/pkg/src/db/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type controller struct {
	DB *gorm.DB
}

type JSON struct {
	Status string       `json:"status"`
	User   models.Users `json:"user"`
}

func RegisterRoutes(g *gin.Engine, db *gorm.DB) {
	c := &controller{
		DB: db,
	}

	r := g.Group("/users")
	{
		r.GET("", c.GetUsers)
		r.POST("", c.PostUser)

		id := r.Group("/:id")
		{
			id.GET("", c.GetUser)
			id.PUT("", c.UpdateUser)
			id.DELETE("", c.DeleteUser)
		}
	}

}
