package endpoints

import (
	"github.com/gorilla/mux"
	handler "github.com/iamgenii/svc/probes/pkg/v1/handlers"
)

// NewProbesRoutes All Application Routes Are defend Here
func NewProbesRoutes(router *mux.Router, handler *handler.ProbesHandlers) {
	router.HandleFunc("/health", handler.Health).Methods("GET", "OPTIONS")
	router.HandleFunc("/ready", handler.Readiness).Methods("GET", "OPTIONS")
	router.HandleFunc("/live", handler.Liveness).Methods("GET", "OPTIONS")
}
