package main

import (
	"fmt"
	"database/sql"

	"github.com/MattRighetti/fdm-repository-backend/config"
	"github.com/MattRighetti/fdm-repository-backend/models"
	"github.com/MattRighetti/fdm-repository-backend/routes"

	"github.com/jinzhu/gorm"
)

var err error

func main() {
	config.DB, err = gorm.Open("mysql", config.DbURL(config.BuildDBConfig()))
	config.BareDB, err = sql.Open("mysql", config.DbURL(config.BuildDBConfig()))

	if err != nil {
		fmt.Println("Status:", err)
	}

	defer config.DB.Close()
	config.DB.AutoMigrate(&models.FloodModel{})
	r := routes.SetupRouter()

	r.Run()
}
