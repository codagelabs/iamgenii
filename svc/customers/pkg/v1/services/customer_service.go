package services

import (
	"context"

	"github.com/iamgenii/email"
	imgError "github.com/iamgenii/error"
	log "github.com/iamgenii/logs"
	"github.com/iamgenii/models"
	repository "github.com/iamgenii/svc/customers/pkg/v1/repositories"
	"github.com/iamgenii/utils/crypto_utils"
)

// CustomerService describes the service.
type CustomerService interface {
	CreateCustomer(ctx context.Context, createReq models.Customer) (interface{}, *imgError.IMGNError)
	GetCustomerByID(context.Context, string) (*models.CustomerResp, *imgError.IMGNError)
	GetCustomers(context.Context, models.GetAllCustReq) ([]*models.CustomerResp, *imgError.IMGNError)
	UpdateCustomer(context.Context, string, models.UpdateCustReq) (interface{}, *imgError.IMGNError)
	DeleteCustomer(context.Context, string) (interface{}, *imgError.IMGNError)
}

// customerService having dependencies fro service
type customerService struct {
	customerRepo         repository.CustomerRepository
	sendBlue             *email.SendInBlue
	hashedUtils          crypto_utils.HashUtils
	repoErrorInterceptor imgError.RepoErrorInterceptor
}

// NewCustomerService inject customer repository
func NewCustomerService(customerRepo repository.CustomerRepository,
	sendBlue *email.SendInBlue,
	hashedUtils crypto_utils.HashUtils,
	repoErrorInterceptor imgError.RepoErrorInterceptor,
) CustomerService {

	return &customerService{
		customerRepo:         customerRepo,
		sendBlue:             sendBlue,
		hashedUtils:          hashedUtils,
		repoErrorInterceptor: repoErrorInterceptor,
	}
}

// CreateCustomer checks if email id the customer is exist in database or
func (b *customerService) CreateCustomer(ctx context.Context, customer models.Customer) (interface{}, *imgError.IMGNError) {
	log.Logger(ctx).Debug("CustomerService.CreateCustomer: Create Customer Request : ", customer)
	_, err := b.customerRepo.GetCustomerByEmail(ctx, customer.Email)
	if err == nil {
		log.Logger(ctx).Error("CustomerService.CreateCustomer: Error email is already exist: Error: ", err)
		return customer, imgError.ErrorEmailAlreadyExist
	}
	_, err = b.customerRepo.GetCustomerByPhone(ctx, customer.CustomersPhone)
	if err == nil {
		log.Logger(ctx).Error("CustomerService.CreateCustomer: Error phone is already exist: Error: ", err)
		return customer, imgError.ErrorMobileNumberAlreadyExist
	}
	_, err = b.customerRepo.GetCustomerByUsername(ctx, customer.UserName)
	if err == nil {
		log.Logger(ctx).Error("CustomerService.CreateCustomer: Error username is already exist: Error: ", err)
		return customer, imgError.ErrorUsernameAlreadyExist
	}

	customer.Password, err = b.hashedUtils.GenerateBcrtptHash(customer.Password)
	if err != nil {
		log.Logger(ctx).Error("CustomerService.CreateCustomer: Error in creating password hash: Error: ", err)
		return nil, b.repoErrorInterceptor.ErrorMapper(ctx, err)
	}
	resp, err := b.customerRepo.CreateCustomer(ctx, customer)
	if err != nil {
		log.Logger(ctx).Error("CustomerService.CreateCustomer: Error in creating customer in repository: Error: ", err)
		return nil, b.repoErrorInterceptor.ErrorMapper(ctx, err)
	}
	return resp, nil
}

// GetCustomerByID calls an GetCustomerByID method of repository to retrieve records
func (b *customerService) GetCustomerByID(ctx context.Context, id string) (*models.CustomerResp, *imgError.IMGNError) {
	log.Logger(ctx).Info("CustomerService.GetCustomerByID: Request id: ", id)
	resp, err := b.customerRepo.GetCustomerByID(ctx, id)
	if err != nil {
		log.Logger(ctx).Error("CustomerService.GetCustomerByID: Error in retrieving customer in repository: Error: ", err)
		return nil, b.repoErrorInterceptor.ErrorMapper(ctx, err)
	}
	return resp, nil
}

// GetCustomers call GetCustomer function of repository to get all customers
func (b *customerService) GetCustomers(ctx context.Context, req models.GetAllCustReq) ([]*models.CustomerResp, *imgError.IMGNError) {
	log.Logger(ctx).Info("CustomerService.GetCustomers. Get all customer request: ", req)
	resp, err := b.customerRepo.GetCustomers(ctx, req)
	if err != nil {
		log.Logger(ctx).Error("CustomerService.GetCustomers: Error in retrieving customers in repository: Error: ", err)
		return nil, b.repoErrorInterceptor.ErrorMapper(ctx, err)
	}
	return resp, nil
}

// UpdateCustomer ccc
func (b *customerService) UpdateCustomer(ctx context.Context, id string, customer models.UpdateCustReq) (interface{}, *imgError.IMGNError) {
	log.Logger(ctx).Info("CustomerService.UpdateCustomer: Update request:", customer, " Customer Id: ", id)
	resp, err := b.customerRepo.UpdateCustomer(ctx, id, customer)
	if err != nil {
		log.Logger(ctx).Error("CustomerService.UpdateCustomer: Error in updating customers in repository: Error: ", err)
		return nil, b.repoErrorInterceptor.ErrorMapper(ctx, err)
	}
	return resp, nil
}

// DeleteCustomer  send request to delete customer from db
func (b *customerService) DeleteCustomer(ctx context.Context, id string) (interface{}, *imgError.IMGNError) {
	log.Logger(ctx).Info("CustomerService.DeleteCustomer: Delete request Id: ", id)
	resp, err := b.customerRepo.DeleteCustomer(ctx, id)
	if err != nil {
		log.Logger(ctx).Error("CustomerService.DeleteCustomer: Error in deleting customers in repository: Error: ", err)
		return nil, b.repoErrorInterceptor.ErrorMapper(ctx, err)
	}
	return resp, nil
}
