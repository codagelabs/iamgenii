package endpoints

import (
	"github.com/iamgenii/svc/packages/pkg/v1/handlers"
	"github.com/gorilla/mux"
)

func NewPackageServiceMappingRoutes(router *mux.Router, handler *handlers.PackageServiceMappingHandler) {
	router.HandleFunc("/packages/services-mapping", handler.CreatePackageServiceMappingHandlerFunc).Methods("POST", "OPTIONS")
	router.HandleFunc("/packages/get-services/{packageId}", handler.GetPackageServicesFunc).Methods("GET", "OPTIONS")
	router.HandleFunc("/packages/services/{packageId}", handler.DeleteServiceFromPackagesHandleFunc).Methods("DELETE", "OPTIONS")

}
