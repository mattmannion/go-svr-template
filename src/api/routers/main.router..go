package routers

import "github.com/gin-gonic/gin"

type router struct {
	eng *gin.Engine
}

func Routers(eng *gin.Engine) {
	r := &router{eng: eng}

	r.UserRouter()
	r.AuthRouter()
}
