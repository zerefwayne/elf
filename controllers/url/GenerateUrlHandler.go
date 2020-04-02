package controllers

import (
	"encoding/json"
	"errors"
	"github.com/zerefwayne/elf/database"
	"github.com/zerefwayne/elf/models"
	"github.com/zerefwayne/elf/utils"
	"log"
	"net/http"
)


// handleGenerateElf URL : POST - Input an ElfUrl structure and returns the shortURL ID or error.
func GenerateURLHandler(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	requestBody := new(models.GenerateRequest)

	_ = json.NewDecoder(r.Body).Decode(requestBody)

	newElfUrl := new(models.ElfUrl)
	newElfUrl.ParseForm(requestBody)

	log.Printf("Hit Generate URL Handler %v\n\n", newElfUrl)

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