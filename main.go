package main

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/zerefwayne/elf/database"
	"net/http"
)

// CONTROLLERS SECTION

func handleGenerateElfURL(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Generating the URL"))
}

func handleDefault(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Welcome to elf! It'll soon be up!"))
}

func handleRedirectElfURL(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	_, _ = w.Write([]byte("Redirecting you from "+vars["elfUrl"]+"!"))
}

func main() {

	// DATABASE SETUP

	// Calls the init() function of the database package
	// Initialise database package for the whole app
	_ = database.DB

	// Once the main function closes, the database.Close() will be called
	defer database.DB.Close(context.Background())

	// Pings the database of the database package
	database.TestPing()

	// ROUTER SETUP

	r := mux.NewRouter()

	r.HandleFunc("/", handleDefault).Methods("GET")
	r.HandleFunc("/{elfUrl}", handleRedirectElfURL).Methods("GET")
	r.HandleFunc("/api/generate", handleGenerateElfURL).Methods("POST")

	r.Use(mux.CORSMethodMiddleware(r))

	// SERVER SETUP

	_ = http.ListenAndServe(":5000", r)

}
