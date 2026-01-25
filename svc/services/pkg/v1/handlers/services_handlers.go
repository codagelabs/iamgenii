package handlers

import (
	"fmt"
	"net/http"

	log "github.com/iamgenii/logs"
	"github.com/iamgenii/models"
	"github.com/iamgenii/svc/services/pkg/v1/services"
	"github.com/iamgenii/utils/http_utils"
	"github.com/iamgenii/validator"
)

const serviceId = "serviceId"
const pageQueryParam = "page"

// ServicesHandlers for handlers
type ServicesHandlers struct {
	servicesSvc      services.IamgeniiServices
	httpReader       http_utils.HTTPReader
	httpWriter       http_utils.HTTPWriter
	requestValidator validator.RequestValidator
}

func NewServicesHandlers(servicesSvc services.IamgeniiServices, httpReader http_utils.HTTPReader,
	httpWriter http_utils.HTTPWriter, requestValidator validator.RequestValidator) *ServicesHandlers {
	return &ServicesHandlers{
		servicesSvc:      servicesSvc,
		httpReader:       httpReader,
		httpWriter:       httpWriter,
		requestValidator: requestValidator,
	}
}

func (serviceHandler ServicesHandlers) CreateServicesHandleFunc(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	serviceReq := models.Services{}

	err := serviceHandler.httpReader.ReadInput(&serviceReq, req.Body)
	if err != nil {
		log.Logger(ctx).Error("ServicesHandlers.CreateServicesHandleFunc: Error in reading request body. Error:", err)
		serviceHandler.httpWriter.WriteHTTPError(w, err)
		return
	}

	err = serviceHandler.requestValidator.ValidateReq(ctx, serviceReq)
	if err != nil {
		log.Logger(ctx).Error("ServicesHandlers.CreateServicesHandleFunc: Error in reading request body validation. Error:", err)
		serviceHandler.httpWriter.WriteHTTPError(w, err)
		return
	}

	resp, err := serviceHandler.servicesSvc.CreateServices(ctx, serviceReq)
	if err != nil {
		log.Logger(ctx).Error("ServicesHandlers.CreateServicesHandleFunc: Error raised by service. Error:", err)
		serviceHandler.httpWriter.WriteHTTPError(w, err)
		return
	}
	serviceHandler.httpWriter.WriteOKResponse(w, http.StatusCreated, resp)

}

func (serviceHandler ServicesHandlers) UpdateServicesHandleFunc(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	serviceUpdateReq := models.Services{}

	id, err := serviceHandler.httpReader.GetURLParam(req, serviceId)
	if err != nil {
		log.Logger(ctx).Error("ServicesHandlers.UpdateServicesHandleFunc: Error in reading request url param. Error:", err)
		serviceHandler.httpWriter.WriteHTTPError(w, err)
		return
	}
	strID := fmt.Sprint(*id)

	err = serviceHandler.httpReader.ReadInput(&serviceUpdateReq, req.Body)
	if err != nil {
		log.Logger(ctx).Error("ServicesHandlers.UpdateServicesHandleFunc: Error in reading request body. Error:", err)
		serviceHandler.httpWriter.WriteHTTPError(w, err)
		return
	}

	err = serviceHandler.requestValidator.ValidateReq(ctx, serviceUpdateReq)
	if err != nil {
		log.Logger(ctx).Error("ServicesHandlers.UpdateServicesHandleFunc: Error in reading request body validation. Error:", err)
		serviceHandler.httpWriter.WriteHTTPError(w, err)
		return
	}

	resp, err := serviceHandler.servicesSvc.UpdateService(ctx, serviceUpdateReq, strID)
	if err != nil {
		log.Logger(ctx).Error("ServicesHandlers.UpdateServicesHandleFunc: Error raised by service. Error:", err)
		serviceHandler.httpWriter.WriteHTTPError(w, err)
		return
	}
	serviceHandler.httpWriter.WriteOKResponse(w, http.StatusCreated, resp)

}

func (serviceHandler ServicesHandlers) GetServicesByID(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	id, err := serviceHandler.httpReader.GetURLParam(req, serviceId)
	if err != nil {
		log.Logger(ctx).Error("ServicesHandlers.GetServicesByID: Error in reading request url param. Error:", err)
		serviceHandler.httpWriter.WriteHTTPError(w, err)
		return
	}
	strID := fmt.Sprint(*id)
	resp, err := serviceHandler.servicesSvc.GetServiceById(ctx, strID)
	if err != nil {
		log.Logger(ctx).Error("ServicesHandlers.GetServicesByID: Error raised by service. Error:", err)
		serviceHandler.httpWriter.WriteHTTPError(w, err)
		return
	}
	serviceHandler.httpWriter.WriteOKResponse(w, http.StatusOK, resp)
}

// GetCustomers handler Function
func (serviceHandler ServicesHandlers) GetAllServices(w http.ResponseWriter, req *http.Request) {

	ctx := req.Context()
	page, err := serviceHandler.httpReader.GetURLQueryParam(req, pageQueryParam, true)
	if err != nil {
		log.Logger(ctx).Error("ServicesHandlers.GetAllServices: Error in reading request url query param. Error:", err)
		serviceHandler.httpWriter.WriteHTTPError(w, err)
		return
	}

	getReq := models.GetAllServiceRequest{}
	getReq.Limit = 10
	getReq.Page = *page
	resp, err := serviceHandler.servicesSvc.GetAllService(ctx, getReq)
	if err != nil {
		log.Logger(ctx).Error("ServicesHandlers.GetAllServices: Error raised by service. Error:", err)
		serviceHandler.httpWriter.WriteHTTPError(w, err)
		return
	}
	serviceHandler.httpWriter.WriteOKResponse(w, http.StatusOK, resp)
}

func (serviceHandler ServicesHandlers) DeleteServicesByID(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	id, err := serviceHandler.httpReader.GetURLParam(req, serviceId)
	if err != nil {
		log.Logger(ctx).Error("ServicesHandlers.DeleteServicesByID: Error in reading request url param. Error:", err)
		serviceHandler.httpWriter.WriteHTTPError(w, err)
		return
	}
	strID := fmt.Sprint(*id)
	resp, err := serviceHandler.servicesSvc.DeleteServiceById(ctx, strID)
	if err != nil {
		log.Logger(ctx).Error("ServicesHandlers.DeleteServicesByID: Error raised by service. Error:", err)
		serviceHandler.httpWriter.WriteHTTPError(w, err)
		return
	}
	serviceHandler.httpWriter.WriteOKResponse(w, http.StatusOK, resp)
}
