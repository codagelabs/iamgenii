package endpoints

import (
	handler "github.com/iamgenii/svc/categaries/pkg/v1/handlers"
	"github.com/gorilla/mux"
)

// NewCategoriesRoutes All Application Routes Are define here
func NewCategoriesRoutes(router *mux.Router, handler *handler.CategoriesHandlers) {

	router.HandleFunc("/categories/", handler.CreateCategories).Methods("POST", "OPTIONS")
	router.HandleFunc("/categories/{categoriesId}", handler.GetCategoriesByID).Methods("GET", "OPTIONS")
	router.HandleFunc("/categories/{categoriesId}/subcategories/", handler.GetSubCategories).Methods("GET", "OPTIONS")
	router.HandleFunc("/categories/{categoriesId}", handler.DeleteCategories).Methods("DELETE", "OPTIONS")

}
