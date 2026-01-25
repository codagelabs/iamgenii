package repositories

import (
	"context"
	"time"

	log "github.com/iamgenii/logs"

	"github.com/iamgenii/models"
	"github.com/jinzhu/gorm"
)

// AdminRepository implements all methods in AdminRepository
type AdminRepository interface {
	CreateAdmin(context.Context, models.Admin) (interface{}, error)
	GetAdminByID(context.Context, string) (*models.AdminResp, error)
	GetAdmins(context.Context, models.GetAllAdminReq) (*models.GetAdminsResponse, error)
	UpdateAdmin(context.Context, string, models.UpdateAdminRequest) error
	GetAdminByEmail(context.Context, string) (models.Admin, error)
	DeleteAdmin(context.Context, string) error
	CheckAdminExistOrNot(ctx context.Context, email string, username string, mobile string) (models.Admin, error)
	GetAdminByUsername(ctx context.Context, username string) (*models.Admin, error)
}

// adminRepository **
type adminRepository struct {
	dbConn *gorm.DB
}

// NewAdminRepositoryImpl inject dependencies of DataStore
func NewAdminRepositoryImpl(dbConn *gorm.DB) AdminRepository {
	return &adminRepository{dbConn: dbConn}
}

// CreateAdmin create administrators entry in database
func (adminRepo adminRepository) CreateAdmin(ctx context.Context, admin models.Admin) (interface{}, error) {
	log.Logger(ctx).Debug("AdminRepository.CreateAdmin: in create admin record ")

	adminResp := models.AdminResp{}
	dbConn := adminRepo.dbConn
	if err := dbConn.Table("administrators").Create(&admin).Find(&adminResp).Error; err != nil {
		log.Logger(ctx).Error("AdminRepository.CreateAdmin: Error in create admin record: ", err)
		return nil, err
	}
	return &adminResp, nil
}

// GetAdminByID retries users records by provided admin Id
func (adminRepo adminRepository) GetAdminByID(ctx context.Context,
	id string) (*models.AdminResp, error) {
	log.Logger(ctx).Debug("AdminRepository.GetAdminByID: in get admin record by id")
	dbConn := adminRepo.dbConn
	admin := models.AdminResp{}
	err := dbConn.Table("administrators").Where("administrators_id=?", id).First(&admin).Error
	if err != nil {
		log.Logger(ctx).Error("AdminRepository.GetAdminByID: Error in get admin record by id: ", err)
		return nil, err
	}
	return &admin, nil
}

// GetAdmins retrieve all admin record by page and limit filters
func (adminRepo adminRepository) GetAdmins(ctx context.Context,
	req models.GetAllAdminReq) (*models.GetAdminsResponse, error) {
	log.Logger(ctx).Debug("AdminRepository.GetAdmins: in get admin records.")
	dbConn := adminRepo.dbConn
	var administrators []models.AdminResp
	offset := (req.Page - 1) * req.Limit
	err := dbConn.Table("administrators").
		Where("deleted_at IS NULL").
		Limit(req.Limit).Offset(offset).
		Find(&administrators).Error
	if err != nil {
		log.Logger(ctx).Error("AdminRepository.GetAdmins: Error in get admin records: ", err)
		return nil, err
	}
	return &models.GetAdminsResponse{
		Count:  len(administrators),
		Admins: administrators,
	}, nil
}

// UpdateAdmin retrieve administrators from database
func (adminRepo adminRepository) UpdateAdmin(ctx context.Context, id string,
	updateReq models.UpdateAdminRequest) error {
	log.Logger(ctx).Debug("AdminRepository.UpdateAdmin: in update admin records by id.")
	dbConn := adminRepo.dbConn
	var findAdmin models.Admin
	if err := dbConn.Table("administrators").Where("administrators_id=?", id).First(&findAdmin).Error; err != nil {
		log.Logger(ctx).Error("AdminRepository.UpdateAdmin: in update admin records by id: record not found.")
		return err
	}
	err := dbConn.Table("administrators").Where("administrators_id=?", id).Update(&updateReq).Error
	if err != nil {
		log.Logger(ctx).Error("AdminRepository.UpdateAdmin:Error in update admin records by id: ", err)
		return err
	}
	return nil
}

// DeleteAdmin deletes admin data
func (adminRepo adminRepository) DeleteAdmin(ctx context.Context, id string) error {
	log.Logger(ctx).Debug("AdminRepository.DeleteAdmin: In delete admin records by id.")
	dbConn := adminRepo.dbConn
	err := dbConn.Table("administrators").Where("administrators_id=?", id).
		Update("deleted_at", time.Now())
	if err.Error != nil {
		log.Logger(ctx).Error("AdminRepository.DeleteAdmin:Error in delete admin records by id: ", err)
		return err.Error
	}
	if err.RowsAffected == 0 {
		log.Logger(ctx).Error("AdminRepository.DeleteAdmin:Error not found: ", err)
		return gorm.ErrRecordNotFound
	}
	return nil
}

// GetAdminByEmail find admin details by email id
func (adminRepo adminRepository) GetAdminByEmail(ctx context.Context,
	email string) (models.Admin, error) {
	log.Logger(ctx).Debug("AdminRepository.GetAdminByEmail: In get admin records by email.")
	var admin models.Admin
	dbConn := adminRepo.dbConn
	err := dbConn.Table("administrators").Where("email=?", email).First(&admin).Error
	if err != nil {
		log.Logger(ctx).Error("AdminRepository.GetAdminByEmail: Error in get admin records by email:", err)
		return admin, err
	}
	return admin, nil
}

// CheckAdminExistOrNot find admin details by email id
func (adminRepo adminRepository) CheckAdminExistOrNot(ctx context.Context,
	email, username, phone string) (models.Admin, error) {
	var admin models.Admin
	log.Logger(ctx).Debug("AdminRepository.CheckAdminExistOrNot: In check admin records exist or not.")
	dbConn := adminRepo.dbConn
	err := dbConn.Table("administrators").Where("email= ? OR username = ? OR phone = ? ", email, username, phone).First(&admin).Error
	if err != nil {
		log.Logger(ctx).Error("AdminRepository.CheckAdminExistOrNot: Error in check admin records exist or not: ", err)
		return admin, err
	}
	log.Logger(ctx).Info(admin)
	return admin, nil
}

func (adminRepo adminRepository) GetAdminByUsername(ctx context.Context, username string) (*models.Admin, error) {
	//AdminLogin create authorizations entry in database
	log.Logger(ctx).Debug("AdminRepository.GetAdminByUsername: In GetAdminByUsername method.")
	admin := models.Admin{}

	//repository dbconn
	dbConn := adminRepo.dbConn

	//Query on database
	err := dbConn.Table("administrators").
		Where("username = ?", username).
		First(&admin).Error

	if err != nil {
		log.Logger(ctx).Error(err)
		return nil, err
	}
	return &admin, nil
}
