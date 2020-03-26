package main

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/zerefwayne/elf/database"
	"github.com/zerefwayne/elf/models"
	"log"
	"net/http"
	"strconv"
	"time"
)

// Helper Methods

func generateShortUrl(originalUrl string) string {

	md5Sum := md5.Sum([]byte(originalUrl))
	md5Hash := hex.EncodeToString(md5Sum[:])
	shortUrl := md5Hash[:6]

	return shortUrl
}


// CONTROLLERS SECTION

// handleGenerateElf URL : POST - Input an ElfUrl structure and returns the shortURL ID or error.
func handleGenerateElfURL(w http.ResponseWriter, r *http.Request) {

	log.Println("Route Hit: handleGenerateElfURL")

	if err := r.ParseForm(); err != nil {
		_, _ = fmt.Fprintf(w, "Error while parsing the form: %v", err)
		return
	}

	newElfUrl := new(models.ElfUrl)
	newElfUrl.OriginalURL = r.FormValue("originalUrl")
	newElfUrl.CreatedAt = time.Now()

	newElfUrl.ShortURL = generateShortUrl(r.FormValue("originalUrl"))

	stringToDuration, _ := strconv.ParseInt(r.FormValue("expiresAfter"), 10, 64)
	newElfUrl.ExpiresAt = time.Now().Add(time.Second * time.Duration(stringToDuration))

	_, _ = fmt.Fprintf(w, "Post from website! r.PostForm = %v\n", newElfUrl)

}

func handleDefault(w http.ResponseWriter, r *http.Request) {

	log.Println("Route Hit: handleDefault")

	_, _ = w.Write([]byte("Welcome to elf! It'll soon be up!"))
}

func handleRedirectElfURL(w http.ResponseWriter, r *http.Request) {

	log.Println("Route Hit: handleRedirectElfURL")

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
