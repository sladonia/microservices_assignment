package app

import (
	"client_api/src/controllers"
	"client_api/src/middlewares/logging_middleware"
	"github.com/gorilla/mux"
)

// ConfigureRouter creates and configures router, inits server
// Returns server instance
func ConfigureRouter(listenPort string) *mux.Router {
	r := mux.NewRouter()
	r.NotFoundHandler = &controllers.NotFoundHandler{}

	// route to check if service alive
	r.HandleFunc("/", controllers.RootController.Get)
	// get port by it's abbreviation
	r.HandleFunc("/port/{port_id}", controllers.PortsController.Get).Methods("GET")
	// import ports from json
	r.HandleFunc("/ports", controllers.PortsController.Import).Methods("POST")

	r.Use(logging_middlewaer.LoggingMw)
	return r
}
