package main

import (
	"github.com/Bobby-P-dev/FinalProject2_kel7/database"
	"github.com/Bobby-P-dev/FinalProject2_kel7/router"
)

func main() {
	database.ConnectToDB()
	r := router.StarApp()
	r.Run()
}
