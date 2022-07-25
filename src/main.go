package main

import (
	"fmt"
	"root/src/api/routers"
	"root/src/db"
	"root/src/env"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := env.LoadConfig()
	if err != nil {
		fmt.Println(err)
	}

	g := gin.Default()
	g.SetTrustedProxies([]string{""})

	store, err := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	if err != nil {
		fmt.Println(err)
	}

	g.Use(sessions.Sessions("mysession", store))

	db.Init(cfg)

	routers.Routers(g)

	fmt.Println("live @ http://localhost" + cfg.Port)
	g.Run(cfg.Port)
}
