package database

import (
	"log"
	"os"

	"github.com/Bobby-P-dev/FinalProject2_kel7/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	DB.AutoMigrate(&models.Users{}, &models.Photo{}, &models.SocialMedia{}, &models.Comment{})
	if err != nil {
		log.Fatal("failed to connect database")
	}
}

func GetDB() *gorm.DB {
	return DB
}
