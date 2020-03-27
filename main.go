package main

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/zerefwayne/elf/config"
	url "github.com/zerefwayne/elf/controllers/url"
	"net/http"
)

func main() {

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

	r.Use(mux.CORSMethodMiddleware(r))

	r.HandleFunc("/", url.DefaultHandler).Methods("GET")
	r.HandleFunc("/{shortUrl}", url.RedirectURLHandler).Methods("GET")
	r.HandleFunc("/api/generate", url.GenerateURLHandler).Methods("POST")

	// SERVER SETUP

	_ = http.ListenAndServe(":5000", r)

}
