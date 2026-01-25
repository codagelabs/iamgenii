package services

import (
	"context"

	"github.com/iamgenii/models"
	repository "github.com/iamgenii/svc/device/pkg/v1/repositories"
)

// DeviceService describes the service interface for categories.
type DeviceService interface {
	CreateDevice(ctx context.Context, createReq models.Device) (interface{}, error)
}

type deviceService struct {
	deviceRepo repository.DeviceRepository
}

func (deviceSvc deviceService) CreateDevice(ctx context.Context, createReq models.Device) (interface{}, error) {
	panic("implement me")
}

// NewDeviceServiceImpl inject dependencies device repository
func NewDeviceServiceImpl(deviceRepo repository.DeviceRepository) DeviceService {

	return &deviceService{
		deviceRepo: deviceRepo,
	}
}
