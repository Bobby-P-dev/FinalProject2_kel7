package database

import (
	"log"

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
}

func GetDB() *gorm.DB {
	return DB
}
