package handlers

import (
	"fmt"
	"net/http"

	log "github.com/iamgenii/logs"
	"github.com/iamgenii/models"
	"github.com/iamgenii/svc/packages/pkg/v1/services"
	"github.com/iamgenii/utils/http_utils"
	"github.com/iamgenii/validator"
)

const packageIdUrlParam = "packageId"
const pageQueryParam = "page"
const serviceIdUrlParam = "service_id"

// PackagesHandlers for handlers
type PackagesHandlers struct {
	packagesService  services.PackagesServices
	httpReader       http_utils.HTTPReader
	httpWriter       http_utils.HTTPWriter
	requestValidator validator.RequestValidator
}

func NewPackagesHandlers(servicesSvc services.PackagesServices, httpReader http_utils.HTTPReader,
	httpWriter http_utils.HTTPWriter, requestValidator validator.RequestValidator) *PackagesHandlers {
	return &PackagesHandlers{
		packagesService:  servicesSvc,
		httpReader:       httpReader,
		httpWriter:       httpWriter,
		requestValidator: requestValidator,
	}
}

func (packagesHandlers PackagesHandlers) CreatePackagesHandleFunc(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	packages := models.Packages{}

	readerErr := packagesHandlers.httpReader.ReadInput(&packages, req.Body)
	if readerErr != nil {
		log.Logger(ctx).Error("PackagesHandlers.CreatePackagesHandleFunc: Error in read create request body. Error: ", readerErr)
		packagesHandlers.httpWriter.WriteHTTPError(w, readerErr)
		return
	}

	err := packagesHandlers.requestValidator.ValidateReq(ctx, packages)
	if err != nil {
		log.Logger(ctx).Error("PackagesHandlers.CreatePackagesHandleFunc: Error in validating request data. Error: ", err)
		packagesHandlers.httpWriter.WriteHTTPError(w, err)
		return
	}

	resp, err := packagesHandlers.packagesService.CreatePackages(ctx, packages)
	if err != nil {
		log.Logger(ctx).Error("PackagesHandlers.CreatePackagesHandleFunc: Error raised by service. Error: ", err)
		packagesHandlers.httpWriter.WriteHTTPError(w, err)
		return
	}
	packagesHandlers.httpWriter.WriteOKResponse(w, http.StatusCreated, resp)

}

func (packagesHandlers PackagesHandlers) UpdatePackagesHandleFunc(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	updateReq := models.Packages{}

	id, readerErr := packagesHandlers.httpReader.GetURLParam(req, packageIdUrlParam)
	if readerErr != nil {
		log.Logger(ctx).Error("PackagesHandlers.UpdatePackagesHandleFunc: Error in reading url param: ", packageIdUrlParam, " Error: ", readerErr)
		packagesHandlers.httpWriter.WriteHTTPError(w, readerErr)
		return
	}
	var strID = fmt.Sprint(*id)
	readerErr = packagesHandlers.httpReader.ReadInput(&updateReq, req.Body)
	if readerErr != nil {
		log.Logger(ctx).Error("PackagesHandlers.UpdatePackagesHandleFunc: Error in read create request body. Error: ", readerErr)
		packagesHandlers.httpWriter.WriteHTTPError(w, readerErr)
		return
	}

	err := packagesHandlers.requestValidator.ValidateReq(ctx, updateReq)
	if err != nil {
		log.Logger(ctx).Error("PackagesHandlers.UpdatePackagesHandleFunc: Error in validating request data. Error: ", err)
		packagesHandlers.httpWriter.WriteHTTPError(w, err)
		return
	}

	resp, err := packagesHandlers.packagesService.UpdatePackages(ctx, updateReq, strID)
	if err != nil {
		log.Logger(ctx).Error("PackagesHandlers.UpdatePackagesHandleFunc: Error raised by service. Error: ", err)
		packagesHandlers.httpWriter.WriteHTTPError(w, err)
		return
	}
	packagesHandlers.httpWriter.WriteOKResponse(w, http.StatusCreated, resp)

}

func (packagesHandlers PackagesHandlers) GetPackagesByIDHandlerFunc(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	id, readerErr := packagesHandlers.httpReader.GetURLParam(req, packageIdUrlParam)
	if readerErr != nil {
		log.Logger(ctx).Error("PackagesHandlers.GetPackagesByIDHandlerFunc: Error in reading url param: ", packageIdUrlParam, " Error: ", readerErr)
		packagesHandlers.httpWriter.WriteHTTPError(w, readerErr)
		return
	}
	strID := fmt.Sprint(*id)
	resp, err := packagesHandlers.packagesService.GetPackagesById(ctx, strID)
	if err != nil {
		log.Logger(ctx).Error("PackagesHandlers.GetPackagesByIDHandlerFunc: Error raised by service. Error: ", err)
		packagesHandlers.httpWriter.WriteHTTPError(w, err)
		return
	}
	packagesHandlers.httpWriter.WriteOKResponse(w, http.StatusOK, resp)
}

func (packagesHandlers PackagesHandlers) GetAllPackagesHandlerFunc(w http.ResponseWriter, req *http.Request) {

	ctx := req.Context()
	page, readerErr := packagesHandlers.httpReader.GetURLQueryParam(req, pageQueryParam, true)
	if readerErr != nil {
		log.Logger(ctx).Error("PackagesHandlers.GetAllPackagesHandlerFunc: Error in reading url query param: ", packageIdUrlParam, " Error: ", readerErr)
		packagesHandlers.httpWriter.WriteHTTPError(w, readerErr)
		return
	}

	getReq := models.GetAllPackagesRequest{}
	getReq.Limit = 10
	getReq.Page = *page
	resp, err := packagesHandlers.packagesService.GetAllPackages(ctx, getReq)
	if err != nil {
		log.Logger(ctx).Error("PackagesHandlers.GetAllPackagesHandlerFunc: Error raised by service. Error: ", err)
		packagesHandlers.httpWriter.WriteHTTPError(w, err)
		return
	}
	packagesHandlers.httpWriter.WriteOKResponse(w, http.StatusOK, resp)
}

func (packagesHandlers PackagesHandlers) DeletePackagesByIDHandleFunc(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	id, readerErr := packagesHandlers.httpReader.GetURLParam(req, packageIdUrlParam)
	if readerErr != nil {
		log.Logger(ctx).Error("PackagesHandlers.DeletePackagesByIDHandleFunc: Error in reading url param: ", packageIdUrlParam, " Error: ", readerErr)
		packagesHandlers.httpWriter.WriteHTTPError(w, readerErr)
		return
	}
	strID := fmt.Sprint(*id)
	resp, err := packagesHandlers.packagesService.DeletePackageByID(ctx, strID)
	if err != nil {
		log.Logger(ctx).Error("PackagesHandlers.DeletePackagesByIDHandleFunc: Error raised by service. Error: ", err)
		packagesHandlers.httpWriter.WriteHTTPError(w, err)
		return
	}
	packagesHandlers.httpWriter.WriteOKResponse(w, http.StatusOK, resp)
}
