package main

import (
	"fmt"
	"root/src/api/middleware"
	"root/src/api/routers"
	"root/src/db"
	"root/src/env"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := env.LoadConfig()

	gin.SetMode(gin.ReleaseMode)

	s := gin.New()

	// s.Use(gin.Logger())
	s.Use(gin.Recovery())

	s.SetTrustedProxies([]string{""})

	session_store := db.Init(cfg)

	s.Use(sessions.Sessions(cfg.Session_Name, session_store))

	s.Use(middleware.Logger)

	routers.Routers(s)

	fmt.Println("live @ http://localhost" + cfg.Port)
	s.Run(cfg.Port)
}
