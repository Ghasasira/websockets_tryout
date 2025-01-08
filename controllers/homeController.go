package controllers

import (
	"log"
	"net/http"
)

func HomeController(w http.ResponseWriter, r *http.Request) {
	log.Println("Hello World!")
	w.Write([]byte("Hello World!"))
}
