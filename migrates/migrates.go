package main

import (
	"github.com/Bobby-P-dev/FinalProject2_kel7/database"
	"github.com/Bobby-P-dev/FinalProject2_kel7/initiallizers"
	"github.com/Bobby-P-dev/FinalProject2_kel7/models"
)

func init() {
	initiallizers.LoadEnvVariable()
	database.ConnectToDB()
}
func main() {
	database.DB.AutoMigrate(&models.Users{})
	database.DB.AutoMigrate(&models.Photo{})
	database.DB.AutoMigrate(&models.SocialMedia{})
	database.DB.AutoMigrate(&models.Comment{})
}
