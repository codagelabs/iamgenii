package repositories

import (
	"context"
	"time"

	customErr "github.com/iamgenii/error"
	log "github.com/iamgenii/logs"
	"github.com/iamgenii/models"
	"github.com/jinzhu/gorm"
)

type VendorRepository interface {
	CreateVendor(context.Context, models.Vendor) (interface{}, error)
	GetVendorByID(context.Context, string) (*models.VendorResp, error)
	GetVendors(context.Context, models.GetAllVendorReq) ([]*models.VendorResp, error)
	UpdateVendor(context.Context, string, models.VendorUpdateReq) (interface{}, error)
	GetVendorByEmail(context.Context, string) (*models.Vendor, error)
	GetVendorByPhone(context.Context, string) (*models.Vendor, error)
	DeleteVendor(context.Context, string) (interface{}, error)
}

type vendorRepository struct {
	dbConn *gorm.DB
}

func NewVendorRepository(dbConn *gorm.DB) VendorRepository {
	return &vendorRepository{dbConn: dbConn}
}

func (repos vendorRepository) CreateVendor(ctx context.Context, Vendor models.Vendor) (interface{}, error) {
	var resp models.VendorResp
	dbConn := repos.dbConn

	if err := dbConn.Table("vendors").Create(&Vendor).Scan(&resp).Error; err != nil {
		log.Logger(ctx).Error("VendorRepository.CreateVendor: Error in creating vendor record. Error: ", err)
		return nil, err
	}
	return &resp, nil
}

func (repos vendorRepository) GetVendorByID(ctx context.Context, id string) (*models.VendorResp, error) {

	dbConn := repos.dbConn
	Vendor := models.VendorResp{}
	err := dbConn.Table("vendors").Where("vendors_id = ?", id).Find(&Vendor).Error
	if gorm.IsRecordNotFoundError(err) {
		log.Logger(ctx).Error("VendorRepository.CreateVendor: Error in reading vendor record. Error: ", err)
		return nil, gorm.ErrRecordNotFound
	}
	if err != nil {
		log.Logger(ctx).Error("VendorRepository.CreateVendor: Error in reading vendor record. Error: ", err)
		return nil, err
	}
	return &Vendor, nil
}

func (repos vendorRepository) GetVendors(ctx context.Context, req models.GetAllVendorReq) ([]*models.VendorResp, error) {
	dbConn := repos.dbConn
	var vendors []*models.VendorResp
	offset := (req.Page - 1) * req.Limit
	err := dbConn.Table("vendors").
		Limit(req.Limit).Offset(offset).
		Find(&vendors).Error
	if gorm.IsRecordNotFoundError(err) {
		log.Logger(ctx).Error("VendorRepository.CreateVendor: Error in reading vendors record. Error: ", err)
		return nil, customErr.ErrRecordNotFound
	}
	if err != nil {
		log.Logger(ctx).Error("VendorRepository.CreateVendor: Error in reading vendors record. Error: ", err)
		return nil, err
	}
	return vendors, nil
}

func (repos vendorRepository) UpdateVendor(ctx context.Context, id string, vendor models.VendorUpdateReq) (interface{}, error) {
	dbConn := repos.dbConn
	dbConn.LogMode(true)
	err := dbConn.Table("vendors").Where("vendors_id = ?", id).Updates(&vendor).Error
	if err != nil {
		log.Logger(ctx).Error("VendorRepository.UpdateVendor: Error in update vendor record. Error: ", err)
		return nil, err
	}
	return "record updated successfully", nil

}

func (repos vendorRepository) DeleteVendor(ctx context.Context, id string) (interface{}, error) {

	dbConn := repos.dbConn

	err := dbConn.Table("vendors").Where("vendors_id=?", id).
		Update("deleted_at", time.Now()).Error
	if err != nil {
		log.Logger(ctx).Error("VendorRepository.DeleteVendor: Error in delete vendor record. Error: ", err)
		return nil, err
	}
	return "Record deleted successfully", nil
}

func (repos vendorRepository) GetVendorByPhone(ctx context.Context, phone string) (*models.Vendor, error) {

	dbConn := repos.dbConn
	vendor := models.Vendor{}
	err := dbConn.Table("vendors").Where("mobile_no=?", phone).First(&vendor).Error
	if err != nil {
		log.Logger(ctx).Error("VendorRepository.GetVendorByPhone: Error in reading vendor record by phone. Error: ", err)
		return nil, err
	}
	return &vendor, err
}

// GetVendorByEmail search Vendor by emails
func (repos vendorRepository) GetVendorByEmail(ctx context.Context, email string) (*models.Vendor, error) {

	dbConn := repos.dbConn
	vendor := models.Vendor{}

	err := dbConn.Table("vendors").Where("email=?", email).First(&vendor).Error
	if err != nil {
		log.Logger(ctx).Error("VendorRepository.GetVendorByEmail: Error in reading vendor record by email. Error: ", err)
		return nil, err
	}
	return &vendor, nil
}
