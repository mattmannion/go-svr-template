package db

import (
	"_/src/db/models"
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
	db.Raw("truncate users;").Scan(&models.Users{})
	db.Raw("alter sequence users_id_seq restart;").Scan(&models.Users{})
	db.Raw(`
	insert into users(firstname, lastname)
		values
		('matt', 'mannion'),
		('mack', 'gr'),
		('khris', 'rhodes');
	`).Scan(&models.Users{})

	fmt.Println("db seeded")
}
