package controllers

import (
	"encoding/json"
	"github.com/zerefwayne/elf/database"
	"github.com/zerefwayne/elf/models"
	"github.com/zerefwayne/elf/utils"
	"net/http"
)


// handleGenerateElf URL : POST - Input an ElfUrl structure and returns the shortURL ID or error.
func GenerateURLHandler(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	requestBody := new(models.GenerateRequest)

	_ = json.NewDecoder(r.Body).Decode(requestBody)

	newElfUrl := new(models.ElfUrl)
	newElfUrl.ParseForm(requestBody)

	newElfUrl.ShortURL = models.GenerateShortUrl(newElfUrl.OriginalURL)

	exists, _, existentialCrisis := database.FetchShortURLExists(newElfUrl.ShortURL)

	if existentialCrisis != nil {
		utils.RespondError(w, http.StatusInternalServerError, existentialCrisis)
		return
	}

	if exists == false {
		if err := database.InsertURL(newElfUrl); err != nil {
			utils.RespondError(w, http.StatusInternalServerError, err)
		}
	}

	utils.RespondSuccess(w, http.StatusOK, exists, newElfUrl)

}