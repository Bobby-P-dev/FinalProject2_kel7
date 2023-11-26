package main

import (
	"os"

	"github.com/Bobby-P-dev/FinalProject2_kel7/database"
	"github.com/Bobby-P-dev/FinalProject2_kel7/initiallizers"
	"github.com/Bobby-P-dev/FinalProject2_kel7/router"
)

func init() {
	initiallizers.LoadEnvVariable()

}

func main() {
	PORT := os.Getenv("PORT")
	database.ConnectToDB()
	r := router.StarApp()
	r.Run(PORT)
}
