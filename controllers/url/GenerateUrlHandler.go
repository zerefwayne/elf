package controllers

import (
	"fmt"
	"github.com/zerefwayne/elf/database"
	"github.com/zerefwayne/elf/models"
	"github.com/zerefwayne/elf/utils"
	"log"
	"net/http"
)

// handleGenerateElf URL : POST - Input an ElfUrl structure and returns the shortURL ID or error.
func GenerateURLHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("Route Hit: handleGenerateElfURL")

	if err := r.ParseForm(); err != nil {
		_, _ = fmt.Fprintf(w, "Error while parsing the form: %v", err)
		return
	}

	// Creating a new model out of the parameters

	newElfUrl := new(models.ElfUrl)
	newElfUrl.ParseForm(r)

	// Inserting newElfUrl into database

	insertError := database.InsertURL(newElfUrl)


	if insertError != nil {
		utils.RespondError(w, http.StatusInternalServerError, insertError)
	} else {
		utils.RespondSuccess(w, http.StatusOK, newElfUrl)
	}
}