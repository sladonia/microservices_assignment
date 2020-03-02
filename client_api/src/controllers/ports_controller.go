package controllers

import (
	"client_api/src/grpc_client"
	"client_api/src/json_parser"
	"client_api/src/logger"
	"client_api/src/services"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

var (
	PortsController PortsControllerInterface = &portsController{}
)

type PortsControllerInterface interface {
	// Import handles Port import requests
	// It streams Port data to port_domain grpc service
	Import(w http.ResponseWriter, r *http.Request)
	// Calls port_domain grpc service to retrieve the Port data from the db
	Get(w http.ResponseWriter, r *http.Request)
}

type portsController struct{}

func (c *portsController) Import(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	// Create and fills in the channel with Port data
	// Request body is being read object by object
	portsCh, err := json_parser.GetPortsChannel(r.Body)
	if err != nil {
		msg := "invalid json body"
		logger.Logger.Infow(msg, "error", err)
		apiErr := NewBadRequestApiError(msg)
		RespondError(w, apiErr)
		return
	}
	// Call port_domain grpc service
	importResp, err := services.PortService.Import(portsCh, grpc_client.Conn)
	if err != nil {
		logger.Logger.Infow("unable to call ports grpc service", "error", err)
		apiErr := NewApiError("service unavailable", "internal server error",
			http.StatusInternalServerError)
		RespondError(w, apiErr)
		return
	}
	RespondJSON(w, http.StatusCreated, importResp)
}

func (c *portsController) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	portId := vars["port_id"]

	// Call port_domain grpc service to retrieve the Port data
	port, err := services.PortService.Get(portId, grpc_client.Conn)
	if err != nil {
		logger.Logger.Debugw("unable to get port", "portId", portId, "error", err)
		apiErr := NewNotFoundApiError(fmt.Sprintf(
			"port with abbreviation=%s not found in the database", portId))
		RespondError(w, apiErr)
		return
	}
	RespondJSON(w, http.StatusOK, port)
}
