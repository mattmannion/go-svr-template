package main

import (
	"fmt"
	"root/src/api/routers"
	"root/src/db"
	"root/src/env"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func main() {

	cfg := env.LoadConfig()

	gin.SetMode(gin.ReleaseMode)

	svr := gin.New()

	svr.Use(gin.Logger())
	svr.Use(gin.Recovery())

	svr.SetTrustedProxies([]string{""})

	session_store := db.Init(cfg)

	svr.Use(sessions.Sessions(cfg.Session_Name, session_store))

	routers.Routers(svr)

	fmt.Println("live @ http://localhost" + cfg.Port)
	svr.Run(cfg.Port)

}
