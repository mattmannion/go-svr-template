package main

import (
	"fmt"
	"root/src/api/routers"
	"root/src/db"
	"root/src/env"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := env.LoadConfig()
	if err != nil {
		fmt.Println(err)
	}

	g := gin.Default()
	g.SetTrustedProxies([]string{""})

	db.Init(cfg)

	routers.Routers(g)

	fmt.Println("live @ http://localhost" + cfg.Port)
	g.Run(cfg.Port)
}
