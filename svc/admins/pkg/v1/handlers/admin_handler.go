package handlers

import (
	"fmt"
	"net/http"

	log "github.com/iamgenii/logs"
	"github.com/iamgenii/models"
	service "github.com/iamgenii/svc/admins/pkg/v1/services"
	"github.com/iamgenii/utils/http_utils"
	"github.com/iamgenii/validator"
)

// AdminHandlers for handler Functions
type AdminHandlers struct {
	adminSvc   service.AdminsService
	httpReader http_utils.HTTPReader
	httpWriter http_utils.HTTPWriter
	validator  validator.RequestValidator
}

const (
	//adminIDUrlParam stores path variable name for adminId
	adminIDUrlParam  = "adminId"
	pageIDQueryParam = "pageId"
)

// NewAdminHandler inits dependencies for graphQL and Handlers
func NewAdminHandler(adminService service.AdminsService, httpReader http_utils.HTTPReader, httpWriter http_utils.HTTPWriter, validator validator.RequestValidator) *AdminHandlers {
	return &AdminHandlers{
		adminSvc:   adminService,
		httpReader: httpReader,
		httpWriter: httpWriter,
		validator:  validator,
	}
}

// CreateAdmin handler Function
func (handler AdminHandlers) CreateAdmin(w http.ResponseWriter,
	req *http.Request) {
	ctx := req.Context()
	log.Logger(ctx).Debug("AdminHandlers.CreateAdmin: In Create Admin handler function.")
	admin := models.Admin{}

	err := handler.httpReader.ReadInput(&admin, req.Body)
	if err != nil {
		log.Logger(ctx).Error("AdminHandlers.CreateAdmin: Error in Read request body: ", err)
		handler.httpWriter.WriteHTTPError(w, err)
		return
	}
	validationErr := handler.validator.ValidateReq(ctx, admin)
	if validationErr != nil {
		log.Logger(ctx).Error("AdminHandlers.CreateAdmin: Error in requested param validation: ", validationErr)
		handler.httpWriter.WriteHTTPError(w, validationErr)
		return
	}

	//call to service layer functions
	resp, err := handler.adminSvc.CreateAdmin(ctx, admin)
	if err != nil {
		log.Logger(ctx).Error("AdminHandlers.CreateAdmin: Error by service: ", err)
		handler.httpWriter.WriteHTTPError(w, err)
		return
	}
	//write an http json resp
	handler.httpWriter.WriteOKResponse(w, http.StatusCreated, resp)

}

// GetAdminByID handler Function
func (handler AdminHandlers) GetAdminByID(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	log.Logger(ctx).Error("AdminHandlers.GetAdminByID: In get admin by id handler function.")

	id, err := handler.httpReader.GetURLParam(req, adminIDUrlParam)
	if err != nil {
		log.Logger(ctx).Error("AdminHandlers.GetAdminByID: Error in get admin id form url param: ", err)
		handler.httpWriter.WriteHTTPError(w, err)
		return
	}
	strID := fmt.Sprint(*id)
	resp, err := handler.adminSvc.GetAdminByID(ctx, strID)
	if err != nil {
		log.Logger(ctx).Error("AdminHandlers.GetAdminByID: Error by service: ", err)
		handler.httpWriter.WriteHTTPError(w, err)
		return
	}
	//create ok response
	handler.httpWriter.WriteOKResponse(w, http.StatusOK, resp)
}

// GetAdmins handler Function
func (handler AdminHandlers) GetAdmins(w http.ResponseWriter, req *http.Request) {
	//get request context
	ctx := req.Context()
	log.Logger(ctx).Debug("AdminHandlers.GetAdmins: In get list of admin handler function.")

	page, imgnError := handler.httpReader.GetURLQueryParam(req, pageIDQueryParam, true)
	if imgnError != nil {
		log.Logger(ctx).Error("AdminHandlers.GetAdmins: Error in get page id from query param:", imgnError)
		handler.httpWriter.WriteHTTPError(w, imgnError)
		return
	}

	getReq := models.GetAllAdminReq{}
	//get page limit default to 10
	getReq.Limit = 10
	getReq.Page = *page

	resp, svcErr := handler.adminSvc.GetAdmins(ctx, getReq)
	if svcErr != nil {
		log.Logger(ctx).Error("AdminHandlers.GetAdmins: Error by service: ", svcErr)
		handler.httpWriter.WriteHTTPError(w, svcErr)
		return
	}

	handler.httpWriter.WriteOKResponse(w, http.StatusOK, resp)
}

// UpdateAdmin handler Function
func (handler AdminHandlers) UpdateAdmin(w http.ResponseWriter, req *http.Request) {
	//request context
	ctx := req.Context()
	log.Logger(ctx).Debug("AdminHandlers.UpdateAdmin: In Handle function.")
	//get url path variables
	id, urlParamErr := handler.httpReader.GetURLParam(req, adminIDUrlParam)
	if urlParamErr != nil { //
		log.Logger(ctx).Errorf("AdminHandlers.UpdateAdmin: Error In reading %s url parameter", adminIDUrlParam)
		handler.httpWriter.WriteHTTPError(w, urlParamErr)
		return
	}

	//update request
	admin := models.UpdateAdminRequest{}
	parsingErro := handler.httpReader.ReadInput(&admin, req.Body)
	if parsingErro != nil {
		log.Logger(ctx).Error("AdminHandlers.UpdateAdmin: Error in reading update request body: ", parsingErro)
		handler.httpWriter.WriteHTTPError(w, parsingErro)
		return
	}

	strID := fmt.Sprint(*id)
	resp, urlParamErr := handler.adminSvc.UpdateAdmin(ctx, strID, admin)
	if urlParamErr != nil {
		log.Logger(ctx).Error("AdminHandlers.UpdateAdmin: Error by service: ", urlParamErr)
		handler.httpWriter.WriteHTTPError(w, urlParamErr)
		return
	}
	handler.httpWriter.WriteOKResponse(w, http.StatusOK, resp)
}

// UpdateAdminProfile handler Function
func (handler AdminHandlers) UpdateAdminProfile(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	log.Logger(ctx).Debug("AdminHandler.UpdateAdminProfile: In handler function")
	id, tokenError := handler.httpReader.GetIDFromToken(req)
	if tokenError != nil {
		log.Logger(ctx).Error("AdminHandler.UpdateAdminProfile: Error in get get id from token: ", tokenError)
		handler.httpWriter.WriteHTTPError(w, tokenError)
		return
	}
	reqData := models.UpdateAdminRequest{}

	readInputErr := handler.httpReader.ReadInput(&reqData, req.Body)
	if readInputErr != nil {
		log.Logger(ctx).Error("AdminHandler.UpdateAdminProfile: Error in get update request id from token: ", readInputErr)
		handler.httpWriter.WriteHTTPError(w, readInputErr)
		return
	}

	strID := fmt.Sprint(*id)
	resp, err := handler.adminSvc.UpdateAdmin(ctx, strID, reqData)
	if err != nil {
		log.Logger(ctx).Errorf("AdminHandler.UpdateAdmin: Error by service : %v", err)
		handler.httpWriter.WriteHTTPError(w, err)
		return
	}
	handler.httpWriter.WriteOKResponse(w, http.StatusOK, resp)
}

// GetAdminProfile handler Function
func (handler AdminHandlers) GetAdminProfile(w http.ResponseWriter,
	req *http.Request) {
	ctx := req.Context()
	log.Logger(ctx).Debug("AdminHandler.GetAdminProfile: In handler function")
	id, err := handler.httpReader.GetIDFromToken(req)
	if err != nil {
		log.Logger(ctx).Errorf("AdminHandler.GetAdminProfile: Error In get user id from token : %v", err)
		handler.httpWriter.WriteHTTPError(w, err)
		return
	}
	strID := fmt.Sprint(*id)
	//request to service layer
	resp, err := handler.adminSvc.GetAdminByID(ctx, strID)
	if err != nil {
		log.Logger(ctx).Errorf("AdminHandler.GetAdminProfile: Error by service: %v", err)
		handler.httpWriter.WriteHTTPError(w, err)
		return
	}
	handler.httpWriter.WriteOKResponse(w, http.StatusOK, resp)
}

// DeleteAdmin handler Function
func (handler AdminHandlers) DeleteAdmin(w http.ResponseWriter, req *http.Request) {
	//request context
	ctx := req.Context()

	//get url path variables
	id, err := handler.httpReader.GetURLParam(req, adminIDUrlParam)
	if err != nil { //
		log.Logger(ctx).Errorf("AdminHandler.DeleteAdmin: Error in get url parameter name %s : %v", adminIDUrlParam, err)
		handler.httpWriter.WriteHTTPError(w, err)
		return
	}
	strID := fmt.Sprint(*id)
	//request to service layer
	resp, err := handler.adminSvc.DeleteAdmin(ctx, strID)
	if err != nil {
		log.Logger(ctx).Errorf("AdminHandler.DeleteAdmin: Error by service : %v", err)
		handler.httpWriter.WriteHTTPError(w, err)
		return
	}
	//create ok response
	handler.httpWriter.WriteOKResponse(w, http.StatusOK, resp)

}
