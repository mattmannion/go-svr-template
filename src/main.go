package main

import (
	"fmt"
	"mm/pkg/src/controllers/users"
	"mm/pkg/src/db"
	"mm/pkg/src/env"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := env.LoadConfig()
	if err != nil {
		fmt.Println(err)
	}

	g := gin.Default()
	g.SetTrustedProxies([]string{""})

	DB := db.Init(cfg.DSN)

	users.RegisterRoutes(g, DB)
	if cfg.Env == "dev" {
		db.Seed()
		fmt.Println("db seeded")
	}

	fmt.Println("live @ http://localhost" + cfg.Port)
	g.Run(cfg.Port)
}
