package controllers

import (
	"client_api/src/json_parser"
	"client_api/src/logger"
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

	n := 0
	for port := range portsCh {
		_ = port
		n++
	}
	RespondJSON(w, http.StatusOK, struct {
		Message    string `json:"message"`
		PortsCount int    `json:"ports_count"`
	}{Message: "success", PortsCount: n})
}

func (c *portsController) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	portId := vars["port_id"]

	RespondJSON(w, http.StatusOK, struct {
		Message string `json:"message"`
		Id      string `json:"id"`
	}{Message: "success", Id: portId})
}
