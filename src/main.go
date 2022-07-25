package main

import (
	"_/src/db"
	"_/src/env"
	"_/src/routers"
	"fmt"

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
