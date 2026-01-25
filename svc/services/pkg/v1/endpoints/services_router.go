package endpoints

import (
	"github.com/iamgenii/svc/services/pkg/v1/handlers"
	"github.com/gorilla/mux"
)

func NewServicesRoutes(router *mux.Router, handler *handlers.ServicesHandlers) {
	router.HandleFunc("/services-create/", handler.CreateServicesHandleFunc).Methods("POST", "OPTIONS")
	router.HandleFunc("/services-get/{serviceId}", handler.GetServicesByID).Methods("GET", "OPTIONS")
	router.HandleFunc("/services-get/", handler.GetAllServices).Methods("GET", "OPTIONS")
	router.HandleFunc("/services-delete/{serviceId}", handler.DeleteServicesByID).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/services-update/{serviceId}", handler.UpdateServicesHandleFunc).Methods("PATCH", "OPTIONS")

}
