package logging_middlewaer

import (
	"client_api/src/logger"
	"net/http"
)

// LoggingMw logs incoming request
func LoggingMw(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Logger.Infow("incoming request", "method", r.Method, "path", r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
