package repositories

import (
	"context"

	"github.com/iamgenii/models"
	"github.com/jinzhu/gorm"
)

// DeviceRepository implements all methods in DeviceRepository
type DeviceRepository interface {
	InsertDeviceRecord(context.Context, models.Device) (interface{}, error)
	GetDeviceRecordByID(context.Context, string) (*models.Device, error)
	GetDeviceRecords(context.Context, string) ([]*models.Device, error)
	DeleteDeviceRecordByID(context.Context, string) (interface{}, error)
	UpdateDeviceRecordByID(context.Context, string, models.Device) (interface{}, error)
}

// DeviceRepositoryImpl **
type deviceRepositoryImpl struct {
	dbConn *gorm.DB
}

func (deviceRepo deviceRepositoryImpl) InsertDeviceRecord(context.Context, models.Device) (interface{}, error) {
	panic("implement me")
}

func (deviceRepo deviceRepositoryImpl) GetDeviceRecordByID(context.Context, string) (*models.Device, error) {
	panic("implement me")
}

func (deviceRepo deviceRepositoryImpl) GetDeviceRecords(context.Context, string) ([]*models.Device, error) {
	panic("implement me")
}

func (deviceRepo deviceRepositoryImpl) DeleteDeviceRecordByID(context.Context, string) (interface{}, error) {
	panic("implement me")
}

func (deviceRepo deviceRepositoryImpl) UpdateDeviceRecordByID(context.Context, string, models.Device) (interface{}, error) {
	panic("implement me")
}

// NewDeviceRepositoryImpl inject dependencies of DataStore
func NewDeviceRepositoryImpl(dbConn *gorm.DB) DeviceRepository {
	return &deviceRepositoryImpl{dbConn: dbConn}
}
