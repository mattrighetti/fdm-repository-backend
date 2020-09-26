package main

import (
	"fmt"

	"github.com/MattRighetti/fdm-repository-backend/Config"
	"github.com/MattRighetti/fdm-repository-backend/Models"
	"github.com/MattRighetti/fdm-repository-backend/Routes"

	"github.com/jinzhu/gorm"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))

	if err != nil {
		fmt.Println("Status:", err)
	}

	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.FloodModel{})
	r := Routes.SetupRouter()
	//running
	r.Run()
}
