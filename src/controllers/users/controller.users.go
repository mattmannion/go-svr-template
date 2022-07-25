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
	c := &controller{DB: db}

	users := g.Group("/users")
	{
		users.GET("", c.GetUsers)
		users.POST("", c.PostUser)

		id := users.Group("/:id")
		{
			id.GET("", c.GetUser)
			id_auth := id.Group("")
			// auth middleware here
			// id_auth.Use()
			{
				id_auth.PUT("", c.UpdateUser)
				id_auth.DELETE("", c.DeleteUser)
			}
		}
	}
}

// session middlware example
// func AuthRequired(c *gin.Context) {
// 	session := sessions.Default(c)
// 	user := session.Get(userkey)
// 	if user == nil {
// 		// Abort the request with the appropriate error code
// 		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
// 		return
// 	}
// 	// Continue down the chain to handler etc
// 	c.Next()
// }
