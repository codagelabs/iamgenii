package repositories

import (
	"context"
	"strings"
	"time"

	imgnError "github.com/iamgenii/error"
	log "github.com/iamgenii/logs"
	"github.com/iamgenii/models"
	"github.com/jinzhu/gorm"
)

type PackageServiceMappingRepository interface {
	InsertPackagesToServicesRecords(context context.Context, req []models.PackageServiceMapping) (resp interface{}, err error)
	GetAllPackageServices(context context.Context, req models.GetPackageServiceMappingReq) ([]models.Services, error)
	DeletePackageServiceRecord(context context.Context, serviceId uint64, packageId uint64) error
}

type packagesToServicesRepository struct {
	dbConn *gorm.DB
}

func NewPackagesToServicesRepository(dbConn *gorm.DB) PackageServiceMappingRepository {
	return &packagesToServicesRepository{dbConn: dbConn}
}

func (repo packagesToServicesRepository) InsertPackagesToServicesRecords(context context.Context, req []models.PackageServiceMapping) (resp interface{}, err error) {
	txn := repo.dbConn.Begin()
	packagesToServices := make([]models.PackageServiceMapping, 0)
	for _, svcCat := range req {
		err = repo.dbConn.Table("packages_to_services").Create(svcCat).Scan(&svcCat).Error
		if err != nil {
			if strings.Contains(err.Error(), "Duplicate entry") {
				log.Logger(context).Error("PackageServiceMappingRepository.InsertPackagesToServicesRecords: Error duplicate entry records. Error: ", err)

				txn.Rollback()
				return nil, imgnError.ErrDuplicateEntry
			}
			txn.Rollback()
			log.Logger(context).Error("PackageServiceMappingRepository.InsertPackagesToServicesRecords: Error in inserting package to service records. Error ", err)
			return nil, err
		}
		packagesToServices = append(packagesToServices, svcCat)
	}
	txn.Commit()
	return packagesToServices, nil
}

func (repo packagesToServicesRepository) GetAllPackageServices(context context.Context, req models.GetPackageServiceMappingReq) ([]models.Services, error) {
	dbConn := repo.dbConn
	services := []models.Services{}
	offset := (req.Page - 1) * req.Limit
	dbConn = dbConn.Set("gorm:auto_preload", true)
	err := dbConn.Table("packages_to_services").
		Joins("INNER JOIN packages ON packages.packages_id = packages_to_services.packages_id").
		Joins("INNER JOIN services ON services.services_id = packages_to_services.services_id").
		Where("packages_to_services.packages_id=?", req.PackageId).
		Limit(req.Limit).Offset(offset).
		Find(&services).Error
	if gorm.IsRecordNotFoundError(err) || len(services) == 0 {
		log.Logger(context).Error("PackageServiceMappingRepository.GetAllPackageServices: Error in reading package to service records. Error ", err)
		return nil, imgnError.ErrRecordNotFound
	}
	return services, nil
}
func (repo packagesToServicesRepository) DeletePackageServiceRecord(context context.Context, serviceId uint64, packageId uint64) error {

	repo.dbConn.LogMode(true)
	txn := repo.dbConn.Begin()
	updateRecord := txn.Table("packages_to_services").
		Where("services_id = ? AND packages_id=?", serviceId, packageId).
		Update("deleted_at", time.Now())
	if updateRecord.RowsAffected == 0 {
		txn.Rollback()
		log.Logger(context).Error("Error in delete record by id : No Record Found Error : ", imgnError.ErrRecordNotFound)
		return updateRecord.Error
	}
	if updateRecord.Error != nil {
		txn.Rollback()
		log.Logger(context).Error("Error in delete record by id : Repository : ", updateRecord.Error)
		return imgnError.ErrRecordNotFound
	}
	txn.Commit()
	return nil

}

func (repo packagesToServicesRepository) UpdatePackagesToServicesRecords(context context.Context, req []models.PackageServiceMapping) (resp interface{}, err error) {
	txn := repo.dbConn.Begin()
	packagesToServices := make([]models.PackageServiceMapping, 0)
	for _, svcCat := range req {
		err = repo.dbConn.Table("packages_to_services").Create(svcCat).Scan(&svcCat).Error
		if err != nil {
			if strings.Contains(err.Error(), "Duplicate entry") {
				log.Logger(context).Error("PackageServiceMappingRepository.UpdatePackagesToServicesRecords: Error in duplicate entry : ", err)

				txn.Rollback()
				return nil, imgnError.ErrDuplicateEntry
			}
			txn.Rollback()
			log.Logger(context).Error("PackageServiceMappingRepository.UpdatePackagesToServicesRecords: Error in inserting package to service records : ", err)
			return nil, err
		}
		packagesToServices = append(packagesToServices, svcCat)
	}
	txn.Commit()
	return packagesToServices, nil
}
