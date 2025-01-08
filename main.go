package main

import (
	"experimental/routes"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting................................")
	r := routes.RoutesSetup()
	http.ListenAndServe(":3000", r)
}
