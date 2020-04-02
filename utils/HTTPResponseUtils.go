package utils

import (
	"encoding/json"
	"github.com/zerefwayne/elf/models"
	"net/http"
)

func RespondSuccess(w http.ResponseWriter, code int, exists bool, payload interface{}) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	response := new(models.Response)

	response.Success = true
	response.Payload = payload
	response.AlreadyExists = exists

	responseStr, _ := json.Marshal(response)

	_, _ = w.Write(responseStr)

}

func RespondError(w http.ResponseWriter, code int, err error) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	response := new(models.Response)

	response.Success = false
	response.Payload = err

	responseStr, _ := json.Marshal(response)

	_, _ = w.Write(responseStr)

}