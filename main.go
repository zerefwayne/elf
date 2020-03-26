package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/zerefwayne/elf/database"
	"github.com/zerefwayne/elf/models"
	"log"
	"net/http"
)


// CONTROLLERS SECTION

// Three types of outputs:
// 1: Success (Successfully created, return 200, Short URL ID), 2: Already Exists (Return 200, Short URL ID)
// 3: Error (Couldn't create short URL, Error)

// handleGenerateElf URL : POST - Input an ElfUrl structure and returns the shortURL ID or error.
func handleGenerateElfURL(w http.ResponseWriter, r *http.Request) {

	log.Println("Route Hit: handleGenerateElfURL")

	if err := r.ParseForm(); err != nil {
		_, _ = fmt.Fprintf(w, "Error while parsing the form: %v", err)
		return
	}

	newElfUrl := new(models.ElfUrl)

	newElfUrl.ParseForm(r)

	insertQuery := "INSERT INTO shorturl(original_url, short_url, expires_at, created_at, has_expired) VALUES($1, $2, $3, $4, $5)"

	_, insertError := database.DB.Exec(context.Background(), insertQuery, newElfUrl.OriginalURL, newElfUrl.ShortURL, newElfUrl.ExpiresAt, newElfUrl.CreatedAt, newElfUrl.HasExpired)

	w.Header().Set("Content-Type", "application/json")

	response := new(models.GenerateResponse)

	if insertError != nil {

		response.Success = false
		response.Error = insertError

	} else {

		response.Success = true
		response.AlreadyExists = false

		response.GeneratedUrl.ShortURL = newElfUrl.ShortURL
		response.GeneratedUrl.OriginalURL = newElfUrl.OriginalURL
		response.GeneratedUrl.ExpiresAt = newElfUrl.ExpiresAt
		response.GeneratedUrl.HasExpired = newElfUrl.HasExpired
		response.GeneratedUrl.CreatedAt = newElfUrl.CreatedAt
	}

	jsonResponse, _ := json.Marshal(response)

	_, _ = w.Write(jsonResponse)

}

func handleDefault(w http.ResponseWriter, r *http.Request) {

	log.Println("Route Hit: handleDefault")

	_, _ = w.Write([]byte("Welcome to elf! It'll soon be up!"))
}

func handleRedirectElfURL(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	log.Println("Route Hit: handleRedirectElfURL")

	vars := mux.Vars(r)

	query := `SELECT original_url FROM shorturl WHERE short_url = '`+vars["elfUrl"]+`' LIMIT 1;`

	log.Println(query)

	rows, queryError := database.DB.Query(context.Background(), query)

	if queryError != nil {
		response, _ := json.Marshal(queryError)
		_, _ = w.Write(response)
	} else {

		defer rows.Close()

		for rows.Next() {
			originalUrl := ""
			err := rows.Scan(&originalUrl)

			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("Redirecting to:", originalUrl)
			http.Redirect(w, r, originalUrl, 301)
		}

		_, _ = w.Write([]byte("No URL Found"))

	}

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
