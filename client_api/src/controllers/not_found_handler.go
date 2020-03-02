package controllers

import (
	"client_api/src/logger"
	"fmt"
	"net/http"
)

type NotFoundHandler struct{}

// This handler is called when user requests an unregistered resource
func (h *NotFoundHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logger.Logger.Debugw("incoming request", "method", r.Method, "path", r.URL.Path)
	apiErr := NewNotFoundApiError(fmt.Sprintf("resource %s %s not found", r.Method, r.URL.Path))
	RespondError(w, apiErr)
}
