package main

import (
	"database/sql"
	"github.com/mattrighetti/fdm-repository-backend/config"
	"github.com/mattrighetti/fdm-repository-backend/routes"
	"log"
)

func main() {
	var err error
	config.Db, err = sql.Open("mysql", config.DbURL())
	if err != nil {
		log.Fatalf("coud not connect to database: %v", err)
	}
	r := routes.SetupRouter()
	_ = r.Run()
}
