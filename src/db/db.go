package db

import (
	"fmt"
	"log"
	"mm/pkg/src/db/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func Init(url string) *gorm.DB {
	DB, err = gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	err = DB.AutoMigrate(&models.Users{})

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("db connected")

	return DB
}

func Seed() {
	DB.Raw("truncate users;").Scan(&models.Users{})
	DB.Raw("alter sequence users_id_seq restart; ").Scan(&models.Users{})
	DB.Raw(`
	insert into users(firstname, lastname)
		values
		('matt', 'mannion'),
		('mack', 'gr'),
		('khris', 'rhodes');
	`).Scan(&models.Users{})
}
