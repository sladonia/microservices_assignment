package app

import (
	"client_api/src/config"
	"client_api/src/controllers"
	"client_api/src/logger"
	"client_api/src/middlewares/logging_middleware"
	"github.com/gorilla/mux"
	"net/http"
)

// ConfigureApp loads config and inits logger
func ConfigureApp() error {
	if err := config.Load(); err != nil {
		return err
	}
	if err := logger.InitLogger(config.Config.ServiceName, config.Config.LogLevel); err != nil {
		return err
	}
	return nil
}

// InitApp creates and configures router
// Returns router instance
func InitApp() http.Handler {
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
