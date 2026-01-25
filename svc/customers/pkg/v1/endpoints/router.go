package endpoints

import (
	handler "github.com/iamgenii/svc/customers/pkg/v1/handlers"
	"github.com/gorilla/mux"
)

// NewCustomersRoute All Application Routes Are defend Here
func NewCustomersRoute(router *mux.Router, handler *handler.CustomerHandlers) {
	router.HandleFunc("/customer/", handler.CreateCustomer).Methods("POST", "OPTIONS")
	router.HandleFunc("/customer/{customerId}", handler.GetCustomerByID).Methods("GET", "OPTIONS")
	router.HandleFunc("/customers/", handler.GetCustomers).Methods("GET", "OPTIONS")
	router.HandleFunc("/customer/{customerId}", handler.UpdateCustomer).Methods("PATCH", "OPTIONS")
	router.HandleFunc("/customer/{customerId}", handler.DeleteCustomer).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/customer", handler.CustomerProfile).Methods("GET", "OPTIONS")
	router.HandleFunc("/customer", handler.UpdateProfile).Methods("PUT", "OPTIONS")

}
