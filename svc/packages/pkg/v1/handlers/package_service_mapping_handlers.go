package handlers

import (
	"net/http"

	log "github.com/iamgenii/logs"
	"github.com/iamgenii/models"
	"github.com/iamgenii/svc/packages/pkg/v1/services"
	"github.com/iamgenii/utils/http_utils"
	"github.com/iamgenii/validator"
)

type PackageServiceMappingHandler struct {
	packagesService  services.PackagesServicesMappingService
	httpReader       http_utils.HTTPReader
	httpWriter       http_utils.HTTPWriter
	requestValidator validator.RequestValidator
}

func NewPackageServiceMappingHandler(
	packagesService services.PackagesServicesMappingService,
	httpReader http_utils.HTTPReader,
	httpWriter http_utils.HTTPWriter,
	requestValidator validator.RequestValidator) *PackageServiceMappingHandler {
	return &PackageServiceMappingHandler{
		packagesService:  packagesService,
		httpReader:       httpReader,
		httpWriter:       httpWriter,
		requestValidator: requestValidator,
	}
}

func (handler PackageServiceMappingHandler) CreatePackageServiceMappingHandlerFunc(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	packageServiceMappingReq := models.PackageServiceMappingReq{}

	readerErr := handler.httpReader.ReadInput(&packageServiceMappingReq, req.Body)
	if readerErr != nil {
		log.Logger(ctx).Error("PackageServiceMappingHandler.CreatePackageServiceMappingHandlerFunc: Error in reading request data. Error: ", readerErr)
		handler.httpWriter.WriteHTTPError(w, readerErr)
		return
	}

	err := handler.requestValidator.ValidateReq(ctx, packageServiceMappingReq)
	if err != nil {
		log.Logger(ctx).Error("PackageServiceMappingHandler.CreatePackageServiceMappingHandlerFunc: Error in validating request data. Error: ", err)
		handler.httpWriter.WriteHTTPError(w, err)
		return
	}

	resp, err := handler.packagesService.CreatePackageServiceMapping(ctx, packageServiceMappingReq)
	if err != nil {
		log.Logger(ctx).Error("PackageServiceMappingHandler.CreatePackageServiceMappingHandlerFunc: Error raised by service. Error: ", err)
		handler.httpWriter.WriteHTTPError(w, err)
		return
	}
	handler.httpWriter.WriteOKResponse(w, http.StatusCreated, resp)
}

func (handler PackageServiceMappingHandler) GetPackageServicesFunc(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	page, readerErr := handler.httpReader.GetURLQueryParam(req, pageQueryParam, true)
	if readerErr != nil {
		log.Logger(ctx).Error("PackageServiceMappingHandler.GetPackageServicesFunc: Error in reading url query param : "+pageQueryParam, "Error: ", readerErr)
		handler.httpWriter.WriteHTTPError(w, readerErr)
		return
	}

	id, readerErr := handler.httpReader.GetURLParam(req, packageIdUrlParam)
	if readerErr != nil {
		log.Logger(ctx).Error("PackageServiceMappingHandler.GetPackageServicesFunc: Error in reading url param : "+packageIdUrlParam, "Error: ", readerErr)
		handler.httpWriter.WriteHTTPError(w, readerErr)
		return
	}
	getReq := models.GetPackageServiceMappingReq{}
	getReq.Limit = 10
	getReq.Page = *page
	getReq.PackageId = *id

	resp, err := handler.packagesService.GetPackageServices(ctx, getReq)
	if err != nil {
		log.Logger(ctx).Error("PackageServiceMappingHandler.GetPackageServicesFunc: Error raised by service. Error: ", err)
		handler.httpWriter.WriteHTTPError(w, err)
		return
	}
	handler.httpWriter.WriteOKResponse(w, http.StatusOK, resp)
}

func (handler PackageServiceMappingHandler) DeleteServiceFromPackagesHandleFunc(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	serviceId, readerErr := handler.httpReader.GetURLQueryParam(req, serviceIdUrlParam, true)
	if readerErr != nil {
		log.Logger(ctx).Error("PackageServiceMappingHandler.DeleteServiceFromPackagesHandleFunc: Error in reading url query param : ", serviceId, " Error: ", readerErr)
		handler.httpWriter.WriteHTTPError(w, readerErr)
		return
	}

	packageId, readerErr := handler.httpReader.GetURLParam(req, packageIdUrlParam)
	if readerErr != nil {
		log.Logger(ctx).Error("PackageServiceMappingHandler.DeleteServiceFromPackagesHandleFunc: Error in reading url param : "+packageIdUrlParam, "Error: ", readerErr)
		handler.httpWriter.WriteHTTPError(w, readerErr)
		return
	}

	resp, err := handler.packagesService.DeleteServiceFromPackage(ctx, *packageId, uint64(*serviceId))
	if err != nil {
		log.Logger(ctx).Error("PackageServiceMappingHandler.DeleteServiceFromPackagesHandleFunc: Error raised by service. Error: ", err)
		handler.httpWriter.WriteHTTPError(w, err)
		return
	}
	handler.httpWriter.WriteOKResponse(w, http.StatusOK, resp)
}
