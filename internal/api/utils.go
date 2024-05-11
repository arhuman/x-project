package api

import (
	"net/http"

	"go.uber.org/zap"
)

// Response represents a general response structure
type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func getServerContext(r *http.Request) (*zap.Logger, map[string]interface{}) {
	services := r.Context().Value("services").(map[string]interface{})
	log := r.Context().Value("log").(*zap.Logger)
	return log, services
}
