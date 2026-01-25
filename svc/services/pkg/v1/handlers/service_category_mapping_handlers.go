package handlers

import (
	"net/http"

	log "github.com/iamgenii/logs"
	"github.com/iamgenii/models"
	"github.com/iamgenii/svc/services/pkg/v1/services"
	"github.com/iamgenii/utils/http_utils"
	"github.com/iamgenii/validator"
)

const categoryId = "categoryId"

// ServicesCategoriesHandlers for handlers
type ServicesCategoriesHandlers struct {
	categorySvc      services.IamgeniiCategoriesToServices
	httpReader       http_utils.HTTPReader
	httpWriter       http_utils.HTTPWriter
	requestValidator validator.RequestValidator
}

func NewServicesCategoriesHandlers(servicesSvc services.IamgeniiCategoriesToServices, httpReader http_utils.HTTPReader,
	httpWriter http_utils.HTTPWriter, requestValidator validator.RequestValidator) *ServicesCategoriesHandlers {
	return &ServicesCategoriesHandlers{
		categorySvc:      servicesSvc,
		httpReader:       httpReader,
		httpWriter:       httpWriter,
		requestValidator: requestValidator,
	}
}

func (serviceCategoriesHandler ServicesCategoriesHandlers) ServicesCategoriesMapping(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	serviceToCategoriesReq := models.ServicesToCategoriesReq{}

	err := serviceCategoriesHandler.httpReader.ReadInput(&serviceToCategoriesReq, req.Body)
	if err != nil {
		log.Logger(ctx).Error("ServicesCategoriesHandlers.ServicesCategoriesMapping: Error in reading url parameter. Error: ", err)
		serviceCategoriesHandler.httpWriter.WriteHTTPError(w, err)
		return
	}

	err = serviceCategoriesHandler.requestValidator.ValidateReq(ctx, serviceToCategoriesReq)
	if err != nil {
		log.Logger(ctx).Error("ServicesCategoriesHandlers.ServicesCategoriesMapping: Error in request validation. Error: ", err)
		serviceCategoriesHandler.httpWriter.WriteHTTPError(w, err)
		return
	}

	resp, err := serviceCategoriesHandler.categorySvc.ServiceToCategoryMapping(ctx, serviceToCategoriesReq)
	if err != nil {
		log.Logger(ctx).Error("ServicesCategoriesHandlers.ServicesCategoriesMapping: Error raised by service. Error: ", err)
		serviceCategoriesHandler.httpWriter.WriteHTTPError(w, err)
		return
	}
	serviceCategoriesHandler.httpWriter.WriteOKResponse(w, http.StatusCreated, resp)
}

func (serviceCategoriesHandler ServicesCategoriesHandlers) GetServicesByCategoriesID(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	page, err := serviceCategoriesHandler.httpReader.GetURLQueryParam(req, pageQueryParam, true)
	if err != nil {
		log.Logger(ctx).Error("ServicesCategoriesHandlers.GetServicesByCategoriesID: Error in reading url  query parameter. Error: ", err)
		serviceCategoriesHandler.httpWriter.WriteHTTPError(w, err)
		return
	}

	id, err := serviceCategoriesHandler.httpReader.GetURLParam(req, categoryId)
	if err != nil {
		log.Logger(ctx).Error("ServicesCategoriesHandlers.GetServicesByCategoriesID: Error in reading url parameter. Error: ", err)
		serviceCategoriesHandler.httpWriter.WriteHTTPError(w, err)
		return
	}

	getReq := models.GetServicesByCategoriesReq{}
	getReq.Limit = 10
	getReq.Page = *page
	getReq.CategoryId = *id

	resp, err := serviceCategoriesHandler.categorySvc.GetServiceByCategoriesId(ctx, getReq)
	if err != nil {
		log.Logger(ctx).Error("ServicesCategoriesHandlers.GetServicesByCategoriesID: Error raised by service. Error: ", err)
		serviceCategoriesHandler.httpWriter.WriteHTTPError(w, err)
		return
	}
	serviceCategoriesHandler.httpWriter.WriteOKResponse(w, http.StatusOK, resp)
}

func (serviceCategoriesHandler ServicesCategoriesHandlers) DeleteServiceToCategoryHandleFunc(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	categoriesId, err := serviceCategoriesHandler.httpReader.GetURLQueryParam(req, categoryId, true)
	if err != nil {
		log.Logger(ctx).Error("ServicesCategoriesHandlers.DeleteServiceToCategoryHandleFunc: Error in reading url  query parameter. Error: ", err)
		serviceCategoriesHandler.httpWriter.WriteHTTPError(w, err)
		return
	}

	servicesId, err := serviceCategoriesHandler.httpReader.GetURLParam(req, serviceId)
	if err != nil {
		log.Logger(ctx).Error("ServicesCategoriesHandlers.DeleteServiceToCategoryHandleFunc: Error in reading url parameter. Error: ", err)
		serviceCategoriesHandler.httpWriter.WriteHTTPError(w, err)
		return
	}

	resp, err := serviceCategoriesHandler.categorySvc.DeleteServiceToCategoryMapping(ctx, *servicesId, uint64(*categoriesId))
	if err != nil {
		log.Logger(ctx).Error("ServicesCategoriesHandlers.DeleteServiceToCategoryHandleFunc: Error raised by service. Error: ", err)
		serviceCategoriesHandler.httpWriter.WriteHTTPError(w, err)
		return
	}
	serviceCategoriesHandler.httpWriter.WriteOKResponse(w, http.StatusOK, resp)
}
