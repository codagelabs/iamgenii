package repositories

import (
	"context"
	"time"

	log "github.com/iamgenii/logs"

	"github.com/iamgenii/models"
	"github.com/jinzhu/gorm"
)

// ServicesRepository implements all methods in CitiesRepository
type ServicesRepository interface {
	InsertServiceData(context context.Context, services models.Services) (insertedServiceResponse *models.Services, err error)
	FetchServiceById(context context.Context, serviceId string) (insertedServiceResponse *models.Services, err error)
	FetchAllServicesData(context context.Context, req models.GetAllServiceRequest) (services *[]models.Services, err error)
	UpdateServiceData(context context.Context, services models.Services, serviceId string) (resp *models.Services, err error)
	SoftDeleteServiceData(context context.Context, serviceId string) (err error)
}

type servicesRepository struct {
	dbConn *gorm.DB
}

func NewServicesRepository(dbConn *gorm.DB) ServicesRepository {
	return &servicesRepository{dbConn: dbConn}
}

// InsertServiceData create Services entry in database
func (svcRepo servicesRepository) InsertServiceData(context context.Context, services models.Services) (insertedServiceResponse *models.Services, err error) {

	txn := svcRepo.dbConn.Begin()

	if err := txn.Create(&services).Scan(&services).Error; err != nil {
		txn.Rollback()
		log.Logger(context).Error("ServicesRepository.InsertServiceData: Error in create service record. Error: ", err)
		return nil, err
	}

	services.ServiceDescriptions.ServicesID = services.ServicesID
	if err := txn.Table("services_description").Create(&services.ServiceDescriptions).Scan(&services).Error; err != nil {
		txn.Rollback()
		log.Logger(context).Error("ServicesRepository.InsertServiceData: Error in create service description record. Error: ", err)
		return nil, err
	}
	txn.Commit()
	return &services, nil
}

func (svcRepo servicesRepository) FetchServiceById(context context.Context, serviceId string) (insertedServiceResponse *models.Services, err error) {
	services := models.Services{}
	err = svcRepo.dbConn.Table("services").Where("services_id=?", serviceId).Find(&services).Error
	if gorm.IsRecordNotFoundError(err) {
		log.Logger(context).Error("ServicesRepository.FetchServiceById: Error record not found. Error: ", err)
		return nil, gorm.ErrRecordNotFound
	}
	if err != nil {
		log.Logger(context).Error("ServicesRepository.FetchServiceById: Error in reading service record. Error: ", err)
		return nil, err
	}
	svcRepo.dbConn.Table("services_description").Where("services_id=?", serviceId).Find(&services.ServiceDescriptions)
	return &services, nil
}

func (svcRepo servicesRepository) FetchAllServicesData(context context.Context, req models.GetAllServiceRequest) (*[]models.Services, error) {

	dbConn := svcRepo.dbConn
	services := []models.Services{}
	offset := (req.Page - 1) * req.Limit
	dbConn = dbConn.Set("gorm:auto_preload", true)
	err := dbConn.Table("services").
		Limit(req.Limit).Offset(offset).
		Find(&services).Error
	if err != nil {
		log.Logger(context).Error("ServicesRepository.FetchAllServicesData: Error in reading all service record. Error: ", err)
		return nil, err
	}
	return &services, nil
}

func (svcRepo servicesRepository) UpdateServiceData(context context.Context, services models.Services, serviceId string) (resp *models.Services, err error) {
	txn := svcRepo.dbConn.Begin()
	if err := txn.Model(models.Services{}).Where("services_id=?", serviceId).Update(&services).Scan(&services).Error; err != nil {
		txn.Rollback()
		log.Logger(context).Error("ServicesRepository.UpdateServiceData: Error in update service record by service id. Error: ", err)
		return nil, err
	}

	if err := txn.Table("services_description").Where("services_id=?", serviceId).Update(&services.ServiceDescriptions).Scan(&services.ServiceDescriptions).Error; err != nil {
		txn.Rollback()
		log.Logger(context).Error("ServicesRepository.UpdateServiceData: Error in update service description by service id. Error: ", err)
		return nil, err
	}
	txn.Commit()
	return &services, nil
}

func (svcRepo servicesRepository) SoftDeleteServiceData(context context.Context, serviceId string) (err error) {
	txn := svcRepo.dbConn.Begin()
	if err := txn.Model(models.Services{}).Where("services_id=?", serviceId).Update("deleted_at", time.Now()).Error; err != nil {
		txn.Rollback()
		log.Logger(context).Error("ServicesRepository.SoftDeleteServiceData: Error in deleting service record by service id. Error: ", err)
		return err
	}

	if err := txn.Table("services_description").Where("services_id=?", serviceId).Update("deleted_at", time.Now()).Error; err != nil {
		txn.Rollback()
		log.Logger(context).Error("ServicesRepository.SoftDeleteServiceData: Error in deleting service description record by service id. Error: ", err)
		return err
	}

	if err := txn.Table("services_to_categories").Where("services_id=?", serviceId).Update("deleted_at", time.Now()).Error; err != nil {
		txn.Rollback()
		log.Logger(context).Error("ServicesRepository.SoftDeleteServiceData: Error in deleting service to categories mapping record by service id. Error: ", err)
		return err
	}
	txn.Commit()
	return nil
}
