package endpoints

import (
	"github.com/iamgenii/svc/services/pkg/v1/handlers"
	"github.com/gorilla/mux"
)

func NewCategoriesToServicesRoutes(router *mux.Router, handler *handlers.ServicesCategoriesHandlers) {

	router.HandleFunc("/services-get/categories/{categoryId}", handler.GetServicesByCategoriesID).Methods("GET", "OPTIONS")
	router.HandleFunc("/services/category-mapping", handler.ServicesCategoriesMapping).Methods("POST", "OPTIONS")
	router.HandleFunc("/services/{serviceId}/delete-category-mapping", handler.ServicesCategoriesMapping).Methods("DELETE", "OPTIONS")

}
