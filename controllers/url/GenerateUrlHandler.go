package controllers

import (
	"errors"
	"fmt"
	"github.com/zerefwayne/elf/database"
	"github.com/zerefwayne/elf/models"
	"github.com/zerefwayne/elf/utils"
	"log"
	"net/http"
)

// handleGenerateElf URL : POST - Input an ElfUrl structure and returns the shortURL ID or error.
func GenerateURLHandler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		_, _ = fmt.Fprintf(w, "Error while parsing the form: %v", err)
		return
	}

	// Creating a new model out of the parameters

	newElfUrl := new(models.ElfUrl)
	newElfUrl.ParseForm(r)

	existsAlready, existentialCrisis := database.FetchShortURLExists(newElfUrl.ShortURL)

	if existentialCrisis != nil {
		utils.RespondError(w, http.StatusInternalServerError, existentialCrisis)
	} else {

		if existsAlready == true {

			utils.RespondError(w, http.StatusInternalServerError, errors.New("short URL: "+newElfUrl.ShortURL+" exists in the database."))

		} else {

			// Inserting newElfUrl into database

			insertError := database.InsertURL(newElfUrl)

			if insertError != nil {
				utils.RespondError(w, http.StatusInternalServerError, insertError)
				log.Fatal(insertError)
			} else {
				utils.RespondSuccess(w, http.StatusOK, newElfUrl)
			}

		}

	}


}