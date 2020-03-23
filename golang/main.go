package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/rakhmadbudiono/compare-api-framework/golang/routers"
	"github.com/rakhmadbudiono/compare-api-framework/golang/utils/pg"
)

func main() {
	configurePG()
	startServer()
}

func configurePG() {
	option := pg.Option{
		Host:     os.Getenv("PG_HOST"),
		Port:     os.Getenv("PG_PORT"),
		Database: os.Getenv("PG_DATABASE"),
		User:     os.Getenv("PG_USER"),
		Password: os.Getenv("PG_PASSWORD"),
	}

	err := pg.SetupDatabase(option)
	if err != nil {
		panic(err)
	}

	log.Println("PostgreSQL connection is successfully established!")
}

func startServer() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	api := routers.NewAPI()

	srv := &http.Server{
		Handler:      api.Router,
		Addr:         ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("Starting server on port %s!", port)

	log.Fatal(srv.ListenAndServe())
}
