package endpoints

import (
	handler "github.com/iamgenii/svc/vendors/pkg/v1/handlers"
	"github.com/gorilla/mux"
)

func NewVendorsRoute(router *mux.Router, handler *handler.VendorHandlers) {
	router.HandleFunc("/create-vendor", handler.CreateVendor).Methods("POST", "OPTIONS")
	router.HandleFunc("/get-vendor/{vendorId}", handler.GetVendorByID).Methods("GET", "OPTIONS")
	router.HandleFunc("/get-vendors", handler.GetVendors).Methods("GET", "OPTIONS")
	router.HandleFunc("/update-vendor/{vendorId}", handler.UpdateVendor).Methods("PATCH", "OPTIONS")
	router.HandleFunc("/delete-vendor/{vendorId}", handler.DeleteVendor).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/get-vendor-profile", handler.GetVendorProfile).Methods("GET", "OPTIONS")
	router.HandleFunc("/update-vendor-profile", handler.UpdateProfile).Methods("PUT", "OPTIONS")

}
