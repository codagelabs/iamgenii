package endpoints

import (
	handler "github.com/iamgenii/svc/cities/pkg/v1/handlers"
	"github.com/gorilla/mux"
)

// NewCitiesRoutes All Application Routes Are defend Here
func NewCitiesRoutes(router *mux.Router, handler *handler.CitiesHandlers) {
	router.HandleFunc("/cities/", handler.CreateCities).Methods("POST", "OPTIONS")
	router.HandleFunc("/cities/", handler.GetCities).Methods("GET", "OPTIONS")
}
