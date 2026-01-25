package repositories

import (
	"context"

	log "github.com/iamgenii/logs"

	customError "github.com/iamgenii/error"
	"github.com/iamgenii/models"
	"github.com/jinzhu/gorm"
)

// LoginRepository implements all methods in LoginRepository
type LoginRepository interface {
	AdminLogin(context.Context, models.LoginReq) (*models.Admin, error)
	VendorLogin(context.Context, models.LoginReq) (*models.Vendor, error)
	CustomerLogin(context.Context, models.LoginReq) (*models.Customer, error)
}

// loginRepository **
type loginRepository struct {
	dbConn *gorm.DB
}

// NewLoginRepository inject dependencies of DataStore
func NewLoginRepository(dbConn *gorm.DB) LoginRepository {
	return &loginRepository{dbConn: dbConn}
}

// AdminLogin create authorizations entry in database
func (authorizationRepositoryImpl loginRepository) AdminLogin(ctx context.Context,
	authorization models.LoginReq) (*models.Admin, error) {

	admin := models.Admin{}

	//repository dbconn
	dbConn := authorizationRepositoryImpl.dbConn

	//Query on database
	err := dbConn.Table("administrators").
		Where("username = ? OR email = ? OR phone = ?",
			authorization.Username, authorization.Username, authorization.Username).
		First(&admin).Error

	if gorm.IsRecordNotFoundError(err) {
		log.Logger(ctx).Error(err)
		return nil, customError.ErrRecordNotFound
	}
	if err != nil {
		log.Logger(ctx).Error(err)
		return nil, err
	}
	return &admin, nil
}

// CustomerLogin create authorizations entry in database
func (authorizationRepositoryImpl loginRepository) CustomerLogin(ctx context.Context,
	authorization models.LoginReq) (*models.Customer, error) {

	customer := models.Customer{}

	//repository dbconn
	dbConn := authorizationRepositoryImpl.dbConn

	//Query on database
	err := dbConn.Table("customers").
		Where("user_name = ? OR email = ? OR customers_phone = ?",
			authorization.Username, authorization.Username, authorization.Username).
		First(&customer).Error

	if gorm.IsRecordNotFoundError(err) {
		log.Logger(ctx).Error(err)
		return nil, customError.ErrRecordNotFound
	}
	if err != nil {
		log.Logger(ctx).Error(err)
		return nil, err
	}
	return &customer, nil
}

// CustomerLogin create authorizations entry in database
func (authorizationRepositoryImpl loginRepository) VendorLogin(ctx context.Context,
	authorization models.LoginReq) (*models.Vendor, error) {

	vendor := models.Vendor{}

	//repository dbconn
	dbConn := authorizationRepositoryImpl.dbConn

	//Query on database
	err := dbConn.Table("vendors").
		Where("email = ? OR mobile_no = ?",
			authorization.Username, authorization.Username).
		First(&vendor).Error

	if gorm.IsRecordNotFoundError(err) {
		log.Logger(ctx).Error(err)
		return nil, customError.ErrRecordNotFound
	}
	if err != nil {
		log.Logger(ctx).Error(err)
		return nil, err
	}
	return &vendor, nil
}
