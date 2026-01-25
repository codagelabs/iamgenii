package handlers

import (
	"net/http"

	service "github.com/iamgenii/svc/device/pkg/v1/services"
)

var (
	//DeviceID
	DeviceID = "deviceId"
)

// DeviceHandlersImpl for handler Functions
type DeviceHandlersImpl struct {
	deviceSvc service.DeviceService
}

func NewDeviceHandlerImpl(deviceService service.DeviceService) *DeviceHandlersImpl {
	return &DeviceHandlersImpl{deviceSvc: deviceService}
}

func (deviceHandlersImpl DeviceHandlersImpl) CreateDevice(w http.ResponseWriter, req *http.Request) {
}
