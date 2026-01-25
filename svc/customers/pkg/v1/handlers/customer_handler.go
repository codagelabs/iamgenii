package handlers

import (
	"fmt"

	log "github.com/iamgenii/logs"
	"github.com/iamgenii/utils/http_utils"

	"net/http"

	"github.com/iamgenii/models"
	service "github.com/iamgenii/svc/customers/pkg/v1/services"
)

const (
	customersIdUrlParam = "customerId"
	pageQueryParam      = "page"
)

// CustomerHandlers for handler Functions
type CustomerHandlers struct {
	customerSvc service.CustomerService
	httpReader  http_utils.HTTPReader
	httpWriter  http_utils.HTTPWriter
}

// NewCustomerHandlers inits dependencies for graphQL and Handlers
func NewCustomerHandlers(customerService service.CustomerService, httpReader http_utils.HTTPReader, httpWriter http_utils.HTTPWriter) *CustomerHandlers {
	return &CustomerHandlers{
		customerSvc: customerService,
		httpReader:  httpReader,
		httpWriter:  httpWriter,
	}
}

// CreateCustomer handler Function
func (customerHandlers CustomerHandlers) CreateCustomer(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	customer := models.Customer{}
	readerErr := customerHandlers.httpReader.ReadInput(&customer, req.Body)
	if readerErr != nil {
		log.Logger(ctx).Error("CustomerHandlers.CreateCustomer: Error in read  request body. Error: ", readerErr)
		customerHandlers.httpWriter.WriteHTTPError(w, readerErr)
		return
	}
	resp, err := customerHandlers.customerSvc.CreateCustomer(ctx, customer)
	if err != nil {
		log.Logger(ctx).Error("CustomerHandlers.CreateCustomer: Error raised by service. Error: ", err)
		customerHandlers.httpWriter.WriteHTTPError(w, err)
		return
	}
	customerHandlers.httpWriter.WriteOKResponse(w, http.StatusCreated, resp)
}

// GetCustomerByID handler Function
func (customerHandlers CustomerHandlers) GetCustomerByID(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	id, readerErr := customerHandlers.httpReader.GetURLParam(req, customersIdUrlParam)
	if readerErr != nil {
		log.Logger(ctx).Error("CustomerHandlers.GetCustomerByID: Error in read url parameter. Error: ", readerErr)
		customerHandlers.httpWriter.WriteHTTPError(w, readerErr)
		return
	}
	strID := fmt.Sprint(*id)
	resp, err := customerHandlers.customerSvc.GetCustomerByID(ctx, strID)
	if err != nil {
		log.Logger(ctx).Error("CustomerHandlers.GetCustomerByID: Error raised by service. Error: ", err)
		customerHandlers.httpWriter.WriteHTTPError(w, err)
		return
	}
	customerHandlers.httpWriter.WriteOKResponse(w, http.StatusOK, resp)
}

// GetCustomers handler Function
func (customerHandlers CustomerHandlers) GetCustomers(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	pageNo, readerErr := customerHandlers.httpReader.GetURLQueryParam(req, pageQueryParam, true)
	if readerErr != nil {
		log.Logger(ctx).Error("CustomerHandlers.GetCustomers: Error in read url query parameter. Error: ", readerErr)
		customerHandlers.httpWriter.WriteHTTPError(w, readerErr)
		return
	}
	getReq := models.GetAllCustReq{}
	getReq.Limit = 10
	getReq.Page = *pageNo
	resp, err := customerHandlers.customerSvc.GetCustomers(ctx, getReq)
	if err != nil {
		log.Logger(ctx).Error("CustomerHandlers.GetCustomers: Error raised by service. Error: ", err)
		customerHandlers.httpWriter.WriteHTTPError(w, err)
		return
	}
	customerHandlers.httpWriter.WriteOKResponse(w, http.StatusOK, resp)
}

// UpdateCustomer handler Function
func (customerHandlers CustomerHandlers) UpdateCustomer(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	id, readerErr := customerHandlers.httpReader.GetURLParam(req, customersIdUrlParam)
	if readerErr != nil {
		log.Logger(ctx).Error("CustomerHandlers.UpdateCustomer: Error in read url parameter. Error: ", readerErr)
		customerHandlers.httpWriter.WriteHTTPError(w, readerErr)
		return
	}

	customer := models.UpdateCustReq{}
	readerErr = customerHandlers.httpReader.ReadInput(&customer, req.Body)
	if readerErr != nil {
		log.Logger(ctx).Error("CustomerHandlers.UpdateCustomer: Error in read  request body. Error: ", readerErr)
		customerHandlers.httpWriter.WriteHTTPError(w, readerErr)
		return
	}

	strID := fmt.Sprint(*id)
	resp, err := customerHandlers.customerSvc.UpdateCustomer(ctx, strID, customer)
	if err != nil {
		log.Logger(ctx).Error("CustomerHandlers.CreateCustomer: Error raised by service. Error: ", err)
		customerHandlers.httpWriter.WriteHTTPError(w, err)
		return
	}
	customerHandlers.httpWriter.WriteOKResponse(w, http.StatusOK, resp)
}

// CustomerProfile handler Function
func (customerHandlers CustomerHandlers) CustomerProfile(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	id, readerErr := customerHandlers.httpReader.GetIDFromToken(req)
	if readerErr != nil {
		log.Logger(ctx).Error("CustomerHandlers.CustomerProfile: Error in read customer id from token. Error: ", readerErr)
		customerHandlers.httpWriter.WriteHTTPError(w, readerErr)
		return
	}
	strID := fmt.Sprint(*id)
	resp, err := customerHandlers.customerSvc.GetCustomerByID(ctx, strID)
	if err != nil {
		log.Logger(ctx).Error("CustomerHandlers.CustomerProfile: Error raised by service. Error: ", err)
		customerHandlers.httpWriter.WriteHTTPError(w, err)
		return
	}
	customerHandlers.httpWriter.WriteOKResponse(w, http.StatusOK, resp)
}

// UpdateProfile handler Function
func (customerHandlers CustomerHandlers) UpdateProfile(w http.ResponseWriter,
	req *http.Request) {
	ctx := req.Context()
	id, readerErr := customerHandlers.httpReader.GetIDFromToken(req)
	if readerErr != nil {
		log.Logger(ctx).Error("CustomerHandlers.UpdateProfile: Error in read customer id from token. Error: ", readerErr)
		customerHandlers.httpWriter.WriteHTTPError(w, readerErr)
		return
	}
	customer := models.UpdateCustReq{}
	readerErr = customerHandlers.httpReader.ReadInput(&customer, req.Body)
	if readerErr != nil {
		log.Logger(ctx).Error("CustomerHandlers.UpdateProfile: Error in read request body. Error: ", readerErr)
		customerHandlers.httpWriter.WriteHTTPError(w, readerErr)
		return
	}
	strID := fmt.Sprint(*id)
	resp, err := customerHandlers.customerSvc.UpdateCustomer(ctx, strID, customer)
	if err != nil {
		log.Logger(ctx).Error("CustomerHandlers.UpdateProfile: Error raised by service. Error: ", err)
		customerHandlers.httpWriter.WriteHTTPError(w, err)
		return
	}
	customerHandlers.httpWriter.WriteOKResponse(w, http.StatusOK, resp)
}

// DeleteCustomer handler Function
func (customerHandlers CustomerHandlers) DeleteCustomer(w http.ResponseWriter, req *http.Request) {
	//request context
	ctx := req.Context()

	//get url path variables
	id, readerErr := customerHandlers.httpReader.GetURLParam(req, customersIdUrlParam)
	if readerErr != nil {
		log.Logger(ctx).Error("CustomerHandlers.DeleteCustomer: Error in read customer id from token. Error: ", readerErr)
		customerHandlers.httpWriter.WriteHTTPError(w, readerErr)
		return
	}
	strID := fmt.Sprint(*id)
	resp, err := customerHandlers.customerSvc.DeleteCustomer(ctx, strID)
	if err != nil {
		log.Logger(ctx).Error("CustomerHandlers.DeleteCustomer: Error raised by service. Error: ", err)
		customerHandlers.httpWriter.WriteHTTPError(w, err)
		return
	}
	customerHandlers.httpWriter.WriteOKResponse(w, http.StatusOK, resp)

}
