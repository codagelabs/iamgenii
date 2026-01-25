package handlers

import (
	"fmt"
	"net/http"

	log "github.com/iamgenii/logs"
	"github.com/iamgenii/models"
	service "github.com/iamgenii/svc/vendors/pkg/v1/services"
	"github.com/iamgenii/utils/http_utils"
	"github.com/iamgenii/validator"
)

const (
	vendorsID  = "vendorId"
	pageNumber = "page"
)

type VendorHandlers struct {
	vendorSvc        service.VendorService
	httpReader       http_utils.HTTPReader
	httpWriter       http_utils.HTTPWriter
	requestValidator validator.RequestValidator
}

func NewVendorHandlers(vendorService service.VendorService, httpReader http_utils.HTTPReader, httpWriter http_utils.HTTPWriter, requestValidator validator.RequestValidator) *VendorHandlers {
	return &VendorHandlers{
		vendorSvc:        vendorService,
		httpReader:       httpReader,
		httpWriter:       httpWriter,
		requestValidator: requestValidator,
	}
}

func (vendorHandlers VendorHandlers) CreateVendor(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	vendor := models.Vendor{}
	err := vendorHandlers.httpReader.ReadInput(&vendor, req.Body)
	if err != nil {
		log.Logger(ctx).Error("VendorHandlers.CreateVendor: Error in read create vendor request. Error: ", err)
		vendorHandlers.httpWriter.WriteHTTPError(w, err)
		return
	}

	err = vendorHandlers.requestValidator.ValidateReq(ctx, vendor)
	if err != nil {
		log.Logger(ctx).Error("VendorHandlers.CreateVendor: Error in validation of create vendor request. Error: ", err)
		vendorHandlers.httpWriter.WriteHTTPError(w, err)
		return
	}

	resp, err := vendorHandlers.vendorSvc.CreateVendor(ctx, vendor)
	if err != nil {
		log.Logger(ctx).Error("VendorHandlers.CreateVendor: Error raised by service. Error: ", err)
		vendorHandlers.httpWriter.WriteHTTPError(w, err)
		return
	}
	vendorHandlers.httpWriter.WriteOKResponse(w, http.StatusCreated, resp)
}

func (vendorHandlers VendorHandlers) GetVendorByID(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	id, err := vendorHandlers.httpReader.GetURLParam(req, vendorsID)
	if err != nil {
		log.Logger(ctx).Error("VendorHandlers.GetVendorByID: Error in read url parameter. Error: ", err)
		vendorHandlers.httpWriter.WriteHTTPError(w, err)
		return
	}
	strID := fmt.Sprint(*id)
	resp, err := vendorHandlers.vendorSvc.GetVendorByID(ctx, strID)
	if err != nil {
		log.Logger(ctx).Error("VendorHandlers.GetVendorByID: Error raised by service. Error: ", err)
		vendorHandlers.httpWriter.WriteHTTPError(w, err)
		return
	}
	vendorHandlers.httpWriter.WriteOKResponse(w, http.StatusOK, resp)
}

func (vendorHandlers VendorHandlers) GetVendors(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	pageId, err := vendorHandlers.httpReader.GetURLQueryParam(req, pageNumber, true)
	if err != nil {
		log.Logger(ctx).Error("VendorHandlers.GetVendors: Error in read url query parameter. Error: ", err)
		vendorHandlers.httpWriter.WriteHTTPError(w, err)
		return
	}
	getReq := models.GetAllVendorReq{}
	getReq.Limit = 10
	getReq.Page = *pageId
	resp, err := vendorHandlers.vendorSvc.GetVendors(ctx, getReq)
	if err != nil {
		log.Logger(ctx).Error("VendorHandlers.GetVendors: Error raised by service. Error: ", err)
		vendorHandlers.httpWriter.WriteHTTPError(w, err)
		return
	}
	vendorHandlers.httpWriter.WriteOKResponse(w, http.StatusOK, resp)
}

func (vendorHandlers VendorHandlers) UpdateVendor(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	id, err := vendorHandlers.httpReader.GetURLParam(req, vendorsID)
	if err != nil {
		log.Logger(ctx).Error("VendorHandlers.UpdateVendor: Error in read url param. Error: ", err)
		vendorHandlers.httpWriter.WriteHTTPError(w, err)
		return
	}

	vendor := models.VendorUpdateReq{}
	err = vendorHandlers.httpReader.ReadInput(&vendor, req.Body)
	if err != nil {
		log.Logger(ctx).Error("VendorHandlers.UpdateVendor: Error in read update vendor request. Error: ", err)
		vendorHandlers.httpWriter.WriteHTTPError(w, err)
		return
	}

	strID := fmt.Sprint(*id)
	resp, err := vendorHandlers.vendorSvc.UpdateVendor(ctx, strID, vendor)
	if err != nil {
		log.Logger(ctx).Error("VendorHandlers.GetVendors: Error raised by service. Error: ", err)
		vendorHandlers.httpWriter.WriteHTTPError(w, err)
		return
	}
	vendorHandlers.httpWriter.WriteOKResponse(w, http.StatusOK, resp)
}

func (vendorHandlers VendorHandlers) GetVendorProfile(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	id, err := vendorHandlers.httpReader.GetIDFromToken(req)
	if err != nil {
		log.Logger(ctx).Error("VendorHandlers.GetVendorProfile: Error in get id from token. Error: ", err)
		vendorHandlers.httpWriter.WriteHTTPError(w, err)
		return
	}
	strID := fmt.Sprint(*id)
	resp, err := vendorHandlers.vendorSvc.GetVendorByID(ctx, strID)
	if err != nil {
		log.Logger(ctx).Error("VendorHandlers.GetVendorProfile: Error raised by service. Error: ", err)
		vendorHandlers.httpWriter.WriteHTTPError(w, err)
		return
	}
	vendorHandlers.httpWriter.WriteOKResponse(w, http.StatusOK, resp)
}

func (vendorHandlers VendorHandlers) UpdateProfile(w http.ResponseWriter,
	req *http.Request) {
	ctx := req.Context()
	id, err := vendorHandlers.httpReader.GetIDFromToken(req)
	if err != nil {
		log.Logger(ctx).Error("VendorHandlers.UpdateProfile: Error in get id from token. Error: ", err)
		vendorHandlers.httpWriter.WriteHTTPError(w, err)
		return
	}
	vendor := models.VendorUpdateReq{}
	err = vendorHandlers.httpReader.ReadInput(&vendor, req.Body)
	if err != nil {
		log.Logger(ctx).Error("VendorHandlers.UpdateProfile: Error in read update vendor profile request. Error: ", err)
		vendorHandlers.httpWriter.WriteHTTPError(w, err)
		return
	}
	strID := fmt.Sprint(*id)
	resp, err := vendorHandlers.vendorSvc.UpdateVendor(ctx, strID, vendor)
	if err != nil {
		log.Logger(ctx).Error("VendorHandlers.UpdateProfile: Error raised by service. Error: ", err)
		vendorHandlers.httpWriter.WriteHTTPError(w, err)
		return
	}
	vendorHandlers.httpWriter.WriteOKResponse(w, http.StatusOK, resp)
}

func (vendorHandlers VendorHandlers) DeleteVendor(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	id, err := vendorHandlers.httpReader.GetURLParam(req, vendorsID)
	if err != nil {
		log.Logger(ctx).Error("VendorHandlers.DeleteVendor: Error in read url parameter. Error: ", err)
		vendorHandlers.httpWriter.WriteHTTPError(w, err)
		return
	}
	strID := fmt.Sprint(*id)
	resp, err := vendorHandlers.vendorSvc.DeleteVendor(ctx, strID)
	if err != nil {
		log.Logger(ctx).Error("VendorHandlers.DeleteVendor: Error raised by service. Error: ", err)
		vendorHandlers.httpWriter.WriteHTTPError(w, err)
		return
	}
	vendorHandlers.httpWriter.WriteOKResponse(w, http.StatusOK, resp)

}
