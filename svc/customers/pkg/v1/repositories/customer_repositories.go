package repositories

import (
	"context"
	"time"

	log "github.com/iamgenii/logs"
	"github.com/iamgenii/models"
	"github.com/jinzhu/gorm"
)

// CustomerRepository implements all methods in CustomerRepository
type CustomerRepository interface {
	CreateCustomer(context.Context, models.Customer) (interface{}, error)
	GetCustomerByID(context.Context, string) (*models.CustomerResp, error)
	GetCustomers(context.Context, models.GetAllCustReq) ([]*models.CustomerResp, error)
	UpdateCustomer(context.Context, string, models.UpdateCustReq) (interface{}, error)
	GetCustomerByEmail(context.Context, string) (models.Customer, error)
	GetCustomerByPhone(context.Context, string) (models.Customer, error)
	GetCustomerByUsername(context.Context, string) (models.Customer, error)
	DeleteCustomer(context.Context, string) (interface{}, error)
}

// customerRepository **
type customerRepository struct {
	dbConn *gorm.DB
}

// NewCustomerRepository inject dependencies of DataStore
func NewCustomerRepository(dbConn *gorm.DB) CustomerRepository {
	return &customerRepository{dbConn: dbConn}
}

// CreateCustomer create customers entry in database
func (customerRepositoryImpl customerRepository) CreateCustomer(ctx context.Context, customer models.Customer) (interface{}, error) {
	var resp models.CustomerResp
	dbConn := customerRepositoryImpl.dbConn
	if err := dbConn.Table("customers").Create(&customer).Scan(&resp).Error; err != nil {
		log.Logger(ctx).Error("CustomerRepository.CreateCustomer: Error: ", err)
		return nil, err
	}
	return &resp, nil
}

// GetCustomerByID retries customers records by provided customer Id
func (customerRepositoryImpl customerRepository) GetCustomerByID(ctx context.Context, id string) (*models.CustomerResp, error) {

	dbConn := customerRepositoryImpl.dbConn
	customer := models.CustomerResp{}
	err := dbConn.Table("customers").Where("customers_id = ?", id).Find(&customer).Error
	if err != nil {
		log.Logger(ctx).Error("CustomerRepository.CreateCustomer: Error: ", err)
		return nil, err
	}
	return &customer, nil
}

// GetCustomers retrieve all customer record by page and limit filters
func (customerRepositoryImpl customerRepository) GetCustomers(ctx context.Context, req models.GetAllCustReq) ([]*models.CustomerResp, error) {
	dbConn := customerRepositoryImpl.dbConn
	var customers []*models.CustomerResp
	offset := (req.Page - 1) * req.Limit
	err := dbConn.Table("customers").Limit(req.Limit).Offset(offset).Find(&customers).Error
	if err != nil {
		log.Logger(ctx).Error("CustomerRepository.GetCustomers: Error: ", err)
		return nil, err
	}
	if len(customers) == 0 {
		log.Logger(ctx).Error("CustomerRepository.GetCustomers: Error: record not found error. ")
		return nil, gorm.ErrRecordNotFound
	}
	return customers, nil
}

// UpdateCustomer retrieve customers from database
func (customerRepositoryImpl customerRepository) UpdateCustomer(ctx context.Context, id string, customer models.UpdateCustReq) (interface{}, error) {
	dbConn := customerRepositoryImpl.dbConn
	dbConn.LogMode(true)
	err := dbConn.Table("customers").Where("customers_id = ?", id).Updates(&customer).Error
	if err != nil {
		log.Logger(ctx).Error("CustomerRepository.UpdateCustomer: Error: ", err)
		return nil, err
	}
	return "updated sucessfully", nil

}

// DeleteCustomer **
func (customerRepositoryImpl customerRepository) DeleteCustomer(ctx context.Context, id string) (interface{}, error) {
	dbConn := customerRepositoryImpl.dbConn
	err := dbConn.Table("customers").Where("customers_id=?", id).
		Update("deleted_at", time.Now()).Error
	if err != nil {
		log.Logger(ctx).Error("CustomerRepository.DeleteCustomer: Error: ", err)
		return nil, err
	}
	return "Record deleted successfully", nil
}

// GetCustomerByEmail search customer by emails
func (customerRepositoryImpl customerRepository) GetCustomerByUsername(ctx context.Context,
	username string) (customer models.Customer, err error) {
	dbConn := customerRepositoryImpl.dbConn
	err = dbConn.Table("customers").Where("user_name=?", username).First(&customer).Error
	if err != nil {
		log.Logger(ctx).Error("CustomerRepository.GetCustomerByUsername: Error: ", err)
		return models.Customer{}, err
	}
	return customer, nil
}

// GetCustomerByMobile search  customer by mobile
func (customerRepositoryImpl customerRepository) GetCustomerByPhone(ctx context.Context,
	phone string) (customer models.Customer, err error) {

	dbConn := customerRepositoryImpl.dbConn
	err = dbConn.Table("customers").Where("customers_phone=?", phone).First(&customer).Error
	if err != nil {
		log.Logger(ctx).Error("CustomerRepository.GetCustomerByPhone: Error: ", err)
		return models.Customer{}, err
	}
	return customer, nil
}

// GetCustomerByEmail search customer by emails
func (customerRepositoryImpl customerRepository) GetCustomerByEmail(ctx context.Context,
	email string) (customer models.Customer, err error) {
	dbConn := customerRepositoryImpl.dbConn
	err = dbConn.Table("customers").Where("email=?", email).First(&customer).Error
	if err != nil {
		log.Logger(ctx).Error("CustomerRepository.GetCustomerByEmail: Error: ", err)
		return models.Customer{}, err
	}
	return customer, nil
}
