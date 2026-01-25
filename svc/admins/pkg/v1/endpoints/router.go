package endpoints

import (
	handler "github.com/iamgenii/svc/admins/pkg/v1/handlers"
	"github.com/gorilla/mux"
)

// NewAdminRoute All Application Routes Are defend Here
func NewAdminRoute(routerV1 *mux.Router, handler *handler.AdminHandlers) {

	routerV1.HandleFunc("/admin/", handler.CreateAdmin).Methods("POST", "OPTIONS")
	routerV1.HandleFunc("/admin/{adminId}", handler.GetAdminByID).Methods("GET", "OPTIONS")
	routerV1.HandleFunc("/admins/", handler.GetAdmins).Methods("GET", "OPTIONS")
	routerV1.HandleFunc("/admin/{adminId}", handler.UpdateAdmin).Methods("PATCH", "OPTIONS")
	routerV1.HandleFunc("/admin/", handler.GetAdminProfile).Methods("GET", "OPTIONS")
	routerV1.HandleFunc("/admin/", handler.UpdateAdminProfile).Methods("PUT", "OPTIONS")
	routerV1.HandleFunc("/admin/{adminId}", handler.UpdateAdmin).Methods("PATCH", "OPTIONS")
	routerV1.HandleFunc("/admin/{adminId}", handler.DeleteAdmin).Methods("DELETE", "OPTIONS")

}
