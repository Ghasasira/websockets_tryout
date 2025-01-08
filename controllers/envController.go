package controllers

import (
	"net/http"
	"os"
)

func EnvController(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, env is " + os.Getenv("Test_Env")))
}
