package auth

import "github.com/gin-gonic/gin"

func GetAuth(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Everything is ok",
	})
}
