package database

import (
	"fmt"
	"log"

	"github.com/Bobby-P-dev/FinalProject2_kel7/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	dsn := "host=flora.db.elephantsql.com user=kesnvrjv password=IXjhC8WBdPHCOW_9EKZOEENRvTiDX8iZ dbname=kesnvrjv port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed connect to database")
	}
	DB.AutoMigrate(&models.Users{}, &models.Comment{}, &models.Photo{}, &models.SocialMedia{})
	fmt.Println("connected to database")
}

// GetDB mengembalikan instance DB
func GetDB() *gorm.DB {
	return DB
}
