package controllers

import (
	"log"
	"net/http"
)

func DefaultHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("Route Hit: handleDefault")

	_, _ = w.Write([]byte("Welcome to elf! It'll soon be up!"))
}

