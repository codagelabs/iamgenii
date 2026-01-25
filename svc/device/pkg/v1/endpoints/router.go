package endpoints

import (
	handler "github.com/iamgenii/svc/device/pkg/v1/handlers"
	"github.com/gorilla/mux"
)

func NewDeviceRoutes(router *mux.Router, handler *handler.DeviceHandlersImpl) {
	router.HandleFunc("/api/v1/devices/", handler.CreateDevice).Methods("POST")
}
