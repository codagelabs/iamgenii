package repositories

import (
	"context"
	"time"

	customErr "github.com/iamgenii/error"
	log "github.com/iamgenii/logs"

	"github.com/iamgenii/models"
	"github.com/jinzhu/gorm"
)

type ServicesToCategoriesRepository interface {
	InsertPackagesToServicesRecords(context context.Context, req []models.ServicesToCategories) (resp interface{}, err error)
	FetchServiceByCategoriesID(context context.Context, req models.GetServicesByCategoriesReq) (services *[]models.Services, err error)
	DeleteServiceCategoryRecord(context context.Context, serviceId uint64, categoryId uint64) error
}

type servicesToCategoriesRepository struct {
	dbConn *gorm.DB
}

func NewServicesToCategoriesRepository(dbConn *gorm.DB) ServicesToCategoriesRepository {
	return &servicesToCategoriesRepository{dbConn: dbConn}
}

func (svcRepo servicesToCategoriesRepository) InsertPackagesToServicesRecords(context context.Context, req []models.ServicesToCategories) (resp interface{}, err error) {

	txn := svcRepo.dbConn.Begin()
	svcToCategories := make([]models.ServicesToCategories, 0)
	for _, svcCat := range req {
		err = svcRepo.dbConn.Table("services_to_categories").Create(svcCat).Scan(&svcCat).Error
		if err != nil {
			txn.Rollback()
			log.Logger(context).Error("ServicesToCategoriesRepository.InsertPackagesToServicesRecords: Error in creating service to categories mapping service. Error: ", err)
			return nil, err
		}
		svcToCategories = append(svcToCategories, svcCat)

	}
	txn.Commit()
	return svcToCategories, nil
}

func (svcRepo servicesToCategoriesRepository) FetchServiceByCategoriesID(context context.Context, req models.GetServicesByCategoriesReq) (*[]models.Services, error) {

	dbConn := svcRepo.dbConn
	services := []models.Services{}
	offset := (req.Page - 1) * req.Limit
	dbConn = dbConn.Set("gorm:auto_preload", true)
	err := dbConn.Table("services").
		Joins("INNER JOIN services_to_categories ON services_to_categories.services_id = services.services_id").
		Where("categories_id=?", req.CategoryId).
		Limit(req.Limit).Offset(offset).
		Find(&services).Error
	if gorm.IsRecordNotFoundError(err) || len(services) == 0 {
		log.Logger(context).Error("ServicesToCategoriesRepository.FetchServiceByCategoriesID: Error in reading services by categories id. Error: ", err)
		return nil, gorm.ErrRecordNotFound
	}
	if err != nil {
		log.Logger(context).Error("ServicesToCategoriesRepository.FetchServiceByCategoriesID: Error in reading services by categories id. Error: ", err)
		return nil, err
	}
	return &services, nil
}

func (svcRepo servicesToCategoriesRepository) DeleteServiceCategoryRecord(context context.Context, serviceId uint64, categoryId uint64) error {

	txn := svcRepo.dbConn.Begin()
	updateRecord := txn.Table("services_to_categories").
		Where("services_id = ? AND categories_id=?", serviceId, categoryId).
		Update("deleted_at", time.Now())
	if updateRecord.RowsAffected == 0 {
		txn.Rollback()
		log.Logger(context).Error("ServicesToCategoriesRepository.DeleteServiceCategoryRecord: Error in delete record by id: No Record Found Error: ", customErr.ErrRecordNotFound)
		return updateRecord.Error
	}
	if updateRecord.Error != nil {
		txn.Rollback()
		log.Logger(context).Error("ServicesToCategoriesRepository.DeleteServiceCategoryRecord: Error in delete record by id: Repository: ", updateRecord.Error)
		return customErr.ErrRecordNotFound
	}
	txn.Commit()
	return nil

}
