package repositories

import (
	"context"
	"fmt"
	"time"

	imgnError "github.com/iamgenii/error"
	log "github.com/iamgenii/logs"
	"github.com/iamgenii/models"
	"github.com/jinzhu/gorm"
)

// PackagesRepository implements all methods in CitiesRepository
type PackagesRepository interface {
	InsertPackageData(context context.Context, services models.Packages) (resp *models.Packages, err error)
	UpdatePackageData(context context.Context, packages models.Packages, packageId string) (response *models.Packages, err error)
	GetPackageRecordsById(context context.Context, packageId string) (response *models.Packages, err error)
	GetAllPackageRecords(context context.Context, req models.GetAllPackagesRequest) (resp *[]models.Packages, err error)
	DeletePackageRecordsById(context context.Context, packageId string) (err error)
}

type packagesRepository struct {
	dbConn *gorm.DB
}

func NewPackagesRepository(dbConn *gorm.DB) PackagesRepository {
	return &packagesRepository{dbConn: dbConn}
}

// InsertServiceData create Services entry in database
func (packageRepo packagesRepository) InsertPackageData(context context.Context, packages models.Packages) (response *models.Packages, err error) {

	txn := packageRepo.dbConn.Begin()
	if err := txn.Create(&packages).Scan(&packages).Error; err != nil {
		txn.Rollback()
		log.Logger(context).Error("PackagesRepository.InsertPackageData: Error in insert package record. Error: ", err)
		return nil, err
	}

	txn.Commit()
	return &packages, nil
}

// UpdatePackageData create Services entry in database
func (packageRepo packagesRepository) UpdatePackageData(context context.Context, packages models.Packages, packageId string) (response *models.Packages, err error) {

	txn := packageRepo.dbConn.Begin()
	if err := txn.Model(models.Packages{}).Where("packages_id = ?", packageId).Update(&packages).Scan(&packages).Error; err != nil {
		txn.Rollback()
		log.Logger(context).Error("PackagesRepository.UpdatePackageData: Error in update package record. Error: ", err)
		return nil, err
	}
	txn.Commit()
	return &packages, nil
}

func (packageRepo packagesRepository) GetPackageRecordsById(context context.Context, packageId string) (response *models.Packages, err error) {
	dbConn := packageRepo.dbConn
	var packages models.Packages
	err = dbConn.Model(models.Packages{}).Where("packages_id = ?", packageId).Find(&packages).Error
	if gorm.IsRecordNotFoundError(err) {
		log.Logger(context).Error("PackagesRepository.GetPackageRecordsById: Error in get package record by id. Error: ", err)
		return nil, gorm.ErrRecordNotFound
	}
	if err != nil {
		log.Logger(context).Error("PackagesRepository.GetPackageRecordsById: Error in get package record by id. Error: ", err)
		return nil, err
	}
	return &packages, nil
}

func (packageRepo packagesRepository) GetAllPackageRecords(context context.Context, req models.GetAllPackagesRequest) (*[]models.Packages, error) {

	dbConn := packageRepo.dbConn
	packages := make([]models.Packages, 0)
	offset := (req.Page - 1) * req.Limit
	err := dbConn.Model(models.Packages{}).Limit(req.Limit).Offset(offset).Find(&packages).Error
	if gorm.IsRecordNotFoundError(err) {
		log.Logger(context).Error("PackagesRepository.GetAllPackageRecords: Error in reading all package records. Error: ", err)
		return nil, imgnError.ErrRecordNotFound
	}
	if err != nil {
		log.Logger(context).Error("PackagesRepository.GetAllPackageRecords: Error in reading all package records. Error: ", err)
		return nil, err
	}

	return &packages, nil
}

// UpdatePackageData create Services entry in database
func (packageRepo packagesRepository) DeletePackageRecordsById(context context.Context, packageId string) (err error) {

	packageRepo.dbConn.LogMode(true)
	txn := packageRepo.dbConn.Begin()
	fmt.Print("in repos")
	err = txn.Table("packages").Where("packages_id = ?", packageId).Update("deleted_at", time.Now()).Error
	if err != nil {
		txn.Rollback()
		log.Logger(context).Error("PackagesRepository.DeletePackageRecordsById: Error in delete package records. Error: ", err)
		return err
	}
	txn.Commit()
	return nil
}
