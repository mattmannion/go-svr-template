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

	s := gin.New()

	// custom logger i made
	// s.Use(middleware.Logger)

	s.Use(gin.Logger())
	s.Use(gin.Recovery())

	s.SetTrustedProxies([]string{""})

	session_store := db.Init(cfg)

	s.Use(sessions.Sessions(cfg.Session_Name, session_store))

	routers.Routers(s)

	fmt.Println("live @ http://localhost" + cfg.Port)
	s.Run(cfg.Port)
}
