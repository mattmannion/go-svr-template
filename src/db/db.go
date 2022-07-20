package db

import (
	"log"
	"mm/pkg/src/db/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(url string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	err = db.AutoMigrate(&models.Users{})

	if err != nil {
		log.Fatalln(err)
	}

	return db
}
