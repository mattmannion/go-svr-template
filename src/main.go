package main

import (
	"fmt"
	"mm/pkg/src/controllers/users"
	"mm/pkg/src/db"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("./src/env/.env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)

	g := gin.Default()
	g.SetTrustedProxies([]string{""})

	db := db.Init(viper.Get("DSN").(string))

	users.RegisterRoutes(g, db)

	fmt.Println("live @ http://localhost" + port)
	g.Run(port)
}
