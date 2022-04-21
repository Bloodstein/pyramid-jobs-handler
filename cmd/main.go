package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/Bloodstein/pyramid-jobs-handler/domain"
	"github.com/Bloodstein/pyramid-jobs-handler/pkg/handler"
	"github.com/Bloodstein/pyramid-jobs-handler/pkg/repository"
	"github.com/Bloodstein/pyramid-jobs-handler/pkg/service"
)

func main() {

	port, _ := strconv.Atoi(os.Getenv("POSTGRES_DB_PORT"))
	db := repository.NewPostgreSQLDB(domain.PostgreSQLDBConfiguration{
		Host:     os.Getenv("POSTGRES_DB_HOST"),
		Port:     port,
		Username: os.Getenv("POSTGRES_DB_USERNAME"),
		Password: os.Getenv("POSTGRES_DB_PASSWORD"),
		Database: os.Getenv("POSTGRES_DB_DATABASE"),
	})
	store := repository.NewStore()

	srv := service.NewService(store, db)
	h := handler.NewHandler(srv)

	if err := http.ListenAndServe(":8000", h.Routes()); err != nil {
		log.Fatalf("Fail to launch web server: %s", err.Error())
	}
}
