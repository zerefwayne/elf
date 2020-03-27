package controllers

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/zerefwayne/elf/database"
	"github.com/zerefwayne/elf/utils"
	"log"
	"net/http"
)

func RedirectURLHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("Route Hit: handleRedirectElfURL")

	vars := mux.Vars(r)

	shortUrl := vars["shortUrl"]

	originalUrl, err := database.FetchOriginalURL(shortUrl)

	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError,err)
	} else {
		fmt.Println("Redirecting to:", originalUrl)
		http.Redirect(w, r, originalUrl, 301)
	}

}
