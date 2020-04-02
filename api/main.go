package main

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/zerefwayne/elf/config"
	"github.com/rs/cors"
	url "github.com/zerefwayne/elf/controllers/url"
	"net/http"
	"time"
)

func main() {

	time.Sleep(5*time.Second)

	// DATABASE SETUP

	// Calls the init() function of the database package
	// Initialise database package for the whole app
	_ = config.DB

	// Once the main function closes, the database.Close() will be called
	defer config.DB.Close(context.Background())

	// Pings the database of the database package
	config.TestPing()

	// ROUTER SETUP

	r := mux.NewRouter()

	r.HandleFunc("/", url.DefaultHandler).Methods("GET")
	r.HandleFunc("/{shortUrl}", url.RedirectURLHandler).Methods("GET")
	r.HandleFunc("/api/generate", url.GenerateURLHandler).Methods("POST")

	// SERVER SETUP

	handler := cors.AllowAll().Handler(r)
	_ = http.ListenAndServe(":5000", handler)

}
