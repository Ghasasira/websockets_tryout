package main

import (
	"experimental/routes"
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	fmt.Println("Starting at port................................")
	r := routes.RoutesSetup()
	http.ListenAndServe("0.0.0.0:"+port, r)
}
