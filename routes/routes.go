package routes

import (
	"experimental/controllers"

	"github.com/go-chi/chi/v5"
)

func RoutesSetup() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", controllers.HomeController)
	r.Get("/data", controllers.DataController)
	r.Get("/env", controllers.EnvController)
	r.Get("/commands", controllers.CommandsController)

	return r
}
