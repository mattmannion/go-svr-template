package db

import (
	"fmt"
	"log"
	"root/src/db/models"
	"root/src/db/sql"
	"root/src/env"

	"github.com/gin-contrib/sessions/redis"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init(cfg env.Cfg) redis.Store {
	session_store, err := redis.NewStore(10, "tcp", cfg.Redis_Addr, "", []byte(cfg.Redis_Secret))
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("redis connected")

	DB, err = gorm.Open(postgres.Open(cfg.DSN), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	err = DB.AutoMigrate(&models.Users{})

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("db connected")

	if cfg.Env == "dev" {
		SeedUsers(DB)
		fmt.Println("db seeded")
	}

	return session_store
}

func SeedUsers(db *gorm.DB) {
	db.Raw(sql.Util_truncate_users_query).Scan(&models.Users{})
	db.Raw(sql.Util_reset_users_id_query).Scan(&models.Users{})
	db.Raw(sql.Util_insert_default_users_query).Scan(&models.Users{})
}
