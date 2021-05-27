package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mattrighetti/fdm-repository-backend/config"
)

func newServer() *http.Server {
	// routes setup
	mux := httprouter.New()
	mux.GET("/models", getModels)
	mux.GET("/models/:id", getModelByID)
	mux.GET("/count/models", getNumModels)
	mux.GET("/filers", getFilters)
	mux.GET("/missing-models", getMissingModels)

	// middleware setup
	muxWithLogger := loggerMiddleware(mux)

	return &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: muxWithLogger,
	}
}

func main() {
	var err error
	config.Db, err = sql.Open("mysql", config.DbURL())
	if err != nil {
		log.Fatalf("coud not connect to database: %v", err)
	}
	server := newServer()
	err = server.ListenAndServe()
	if err != nil {
		log.Fatalf("could not serve: %v", err)
	}
}
