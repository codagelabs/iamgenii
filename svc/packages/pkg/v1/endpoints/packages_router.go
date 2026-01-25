package endpoints

import (
	"github.com/iamgenii/svc/packages/pkg/v1/handlers"
	"github.com/gorilla/mux"
)

func NewPackageRoutes(router *mux.Router, handler *handlers.PackagesHandlers) {
	router.HandleFunc("/services/package-create", handler.CreatePackagesHandleFunc).Methods("POST", "OPTIONS")
	router.HandleFunc("/services/package-update/{packageId}", handler.UpdatePackagesHandleFunc).Methods("PATCH", "OPTIONS")
	router.HandleFunc("/services/package-get/{packageId}", handler.GetPackagesByIDHandlerFunc).Methods("GET", "OPTIONS")
	router.HandleFunc("/services/package-get-all", handler.GetAllPackagesHandlerFunc).Methods("GET", "OPTIONS")
	router.HandleFunc("/services/package-delete/{packageId}", handler.DeletePackagesByIDHandleFunc).Methods("DELETE", "OPTIONS")

}
