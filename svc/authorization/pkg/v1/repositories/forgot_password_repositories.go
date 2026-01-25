package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"

	customError "github.com/iamgenii/error"
	log "github.com/iamgenii/logs"
	"github.com/iamgenii/models"
)

// ForgotPasswordRepositories has all method regarding password action
type ForgotPasswordRepositories interface {
	SaveOTPDetails(context.Context, models.MobileVarification) (uint64, error)
	VerifyAdminUser(context.Context, string) (*models.Admin, error)
	VerifyCustomer(context.Context, string) (*models.Customer, error)
	ValidateOTP(context.Context, models.ValidateOtpReq) (err error)
	UpdateAdminPassword(context.Context, string, string) (interface{}, error)
	ValidateUpdateReq(context.Context, string, string) error
	UpdateCustomerPassword(context.Context, string, string) (interface{}, error)
}

// ForgotPasswordRepositoriesImpl has database dependancies
type ForgotPasswordRepositoriesImpl struct {
	dbConn *gorm.DB
}

// NewForgotPasswordRepositories inject dependancies for gorm database connection
func NewForgotPasswordRepositories(dbConn *gorm.DB) ForgotPasswordRepositories {

	return ForgotPasswordRepositoriesImpl{dbConn: dbConn}
}

// VerifyAdminUser stores OTP details into database
func (forgotPasswordRepositories ForgotPasswordRepositoriesImpl) VerifyAdminUser(ctx context.Context,
	mobileNumber string) (*models.Admin, error) {

	//repository dbconn
	dbConn := forgotPasswordRepositories.dbConn

	admin := models.Admin{}
	//Query on database
	err := dbConn.Table("administrators").
		Where("phone = ?", mobileNumber).
		Last(&admin).Error

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

// VerifyCustomer stores OTP details into database
func (forgotPasswordRepositories ForgotPasswordRepositoriesImpl) VerifyCustomer(ctx context.Context,
	mobileNumber string) (*models.Customer, error) {

	//repository dbconn
	dbConn := forgotPasswordRepositories.dbConn

	customer := models.Customer{}
	//Query on database
	err := dbConn.Table("customers").
		Where("customers_phone = ?", mobileNumber).
		Last(&customer).Error

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

// SaveOTPDetails stores OTP details into database
func (forgotPasswordRepositories ForgotPasswordRepositoriesImpl) SaveOTPDetails(ctx context.Context,
	req models.MobileVarification) (uint64, error) {

	dbConn := forgotPasswordRepositories.dbConn
	err := dbConn.Table("mobile_otp_details").Create(&req).Scan(&req).Error
	if err != nil {
		return 0, err
	}
	return req.VarificationID, nil
}

// ValidateOTP **
func (forgotPasswordRepositories ForgotPasswordRepositoriesImpl) ValidateOTP(ctx context.Context,
	req models.ValidateOtpReq) (err error) {

	otpdetails := models.MobileVarification{}
	now := time.Now().Add(-10 * time.Minute)
	dbConn := forgotPasswordRepositories.dbConn
	err = dbConn.Table("mobile_otp_details").
		Where("otp = ? AND mobile_number = ? AND varification_id=?", req.OTP, req.ContactNumber, req.JournyID).
		Where("created_at > ?", now).
		Order("created_at DESC").
		First(&otpdetails).
		Error

	fmt.Println(otpdetails)
	if gorm.IsRecordNotFoundError(err) {
		return customError.ErrOTPExpires
	}

	return err
}

// ValidateUpdateReq **
func (forgotPasswordRepositories ForgotPasswordRepositoriesImpl) ValidateUpdateReq(ctx context.Context,
	contactNumber, journyID string) (err error) {

	otpdetails := models.MobileVarification{}
	now := time.Now().Add(-20 * time.Minute)
	dbConn := forgotPasswordRepositories.dbConn
	err = dbConn.Table("mobile_otp_details").
		Where("mobile_number = ? AND varification_id=?", contactNumber, journyID).
		Where("created_at > ?", now).
		First(&otpdetails).
		Error

	fmt.Println(otpdetails)
	if gorm.IsRecordNotFoundError(err) {
		return customError.ErrOTPExpires
	}

	return err
}

// UpdateAdminPassword **
func (forgotPasswordRepositories ForgotPasswordRepositoriesImpl) UpdateAdminPassword(ctx context.Context,
	phone, password string) (interface{}, error) {

	dbConn := forgotPasswordRepositories.dbConn
	err := dbConn.Table("administrators").Where("phone=?", phone).Update("password", password).Error
	return nil, err
}

// UpdateCustomerPassword **
func (forgotPasswordRepositories ForgotPasswordRepositoriesImpl) UpdateCustomerPassword(ctx context.Context,
	phone, password string) (interface{}, error) {

	dbConn := forgotPasswordRepositories.dbConn
	err := dbConn.Table("customers").Where("customers_phone=?", phone).Update("password", password).Error
	return nil, err
}
