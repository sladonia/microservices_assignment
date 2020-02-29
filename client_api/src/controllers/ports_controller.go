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
	Import(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

type portsController struct{}

func (c *portsController) Import(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	portsCh, err := json_parser.GetPortsChannel(r.Body)
	if err != nil {
		msg := "invalid json body"
		logger.Logger.Infow(msg, "error", err)
		apiErr := NewBadRequestApiError(msg)
		RespondError(w, apiErr)
		return
	}

	importResp, err := services.PortService.Import(portsCh, grpc_client.Conn)
	if err != nil {
		logger.Logger.Infow("unable to call ports grpc service", "error", err)
		apiErr := NewApiError("service unavailable", "internal server error",
			http.StatusInternalServerError)
		RespondError(w, apiErr)
		return
	}
	RespondJSON(w, http.StatusOK, importResp)
}

func (c *portsController) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	portId := vars["port_id"]

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
