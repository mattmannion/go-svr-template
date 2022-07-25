package db

import (
	"_/src/db/models"
	"_/src/db/sql"
	"_/src/env"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func Init(cfg env.Cfg) {
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
		Seed(DB)
	}
}

func Seed(db *gorm.DB) {
	db.Raw(sql.Util_truncate_tables_query).Scan(&models.Users{})
	db.Raw(sql.Util_reset_primary_id_query).Scan(&models.Users{})
	db.Raw(sql.Util_insert_default_users_query).Scan(&models.Users{})
	fmt.Println("db seeded")
}
