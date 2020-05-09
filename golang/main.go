package main

import (
	"log"
	"net/http"
	"os"
	"github.com/rakhmadbudiono/compare-api-framework/golang/domain/book"
	"github.com/rakhmadbudiono/compare-api-framework/golang/util/pg"
	"time"

	"github.com/gorilla/mux"
)

func welcome(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(204)
	w.Write([]byte{})
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	r := mux.NewRouter()
	r.HandleFunc("/", welcome)
	book.SetupRouter(r.PathPrefix("/book").Subrouter())

	readerConfiguration := pg.Option{
		Host:     os.Getenv("PG_HOST"),
		Port:     os.Getenv("PG_PORT"),
		Database: os.Getenv("PG_DATABASE"),
		User:     os.Getenv("PG_USER"),
		Password: os.Getenv("PG_PASSWORD"),
	}
	writerConfiguration := pg.Option{
		Host:     os.Getenv("PG_HOST"),
		Port:     os.Getenv("PG_PORT"),
		Database: os.Getenv("PG_DATABASE"),
		User:     os.Getenv("PG_USER"),
		Password: os.Getenv("PG_PASSWORD"),
	}

	err := pg.SetupDatabase(readerConfiguration, writerConfiguration)
	if err != nil {
		panic(err)
	}

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
