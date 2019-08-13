package main

import (
	"log"
	"time"

	"../../pkgs/http"
	"../../pkgs/http/rest/gin"
	"../../pkgs/storage/postgres"
)

func main() {
	dbConfig := postgres.Config{"localhost", "postgres", 5432, "postgres", ""}
	repo, err := postgres.NewDB(dbConfig)
	if err != nil {
		log.Fatal(err)
	}
	app := gin.Server()
	config := http.Config{8000, time.Second * 5, time.Second * 10}
	server := app.Start(config, repo, repo)
	server.ListenAndServe()
}
