package handlers

import (
	"net/http"

	"github.com/iamgenii/utils/http_utils"

	"github.com/iamgenii/svc/authorization/pkg/v1/services"
)

// ForgotPasswordHandlers stores all password handlers
type ForgotPasswordHandlers struct {
	passwordService services.ForgotPasswordService
	httpReader      http_utils.HTTPReader
	httpWriter      http_utils.HTTPWriter
}

// NewForgotPasswordHandlers service dependency
func NewForgotPasswordHandlers(passwordService services.ForgotPasswordService, httpReader http_utils.HTTPReader, httpWriter http_utils.HTTPWriter) *ForgotPasswordHandlers {
	return &ForgotPasswordHandlers{
		passwordService: passwordService,
		httpReader:      httpReader,
		httpWriter:      httpWriter,
	}
}

// VerifyAndSendOTP handler Function for mobile muber verification and to send OTP
func (forgotPasswordHandlers ForgotPasswordHandlers) VerifyAndSendOTP(w http.ResponseWriter,
	req *http.Request) {

	//ctx := req.Context()
	//otpReq := models.SendOTPReq{}
	//
	//err := forgotPasswordHandlers.httpReader.ReadInput(&otpReq, req.Body)
	//if err != nil {
	//	log.Logger(ctx).Error(err)
	//	forgotPasswordHandlers.httpWriter.WriteHTTPError(w, http.StatusBadRequest)
	//	return
	//}
	//
	////call to service layer functions
	//resp, err := forgotPasswordHandlers.passwordService.VerifyUserAndSendOTP(ctx, otpReq)
	//if err == customError.ErrInvalidMobileNumber {
	//	log.Logger(ctx).Error(err)
	//	forgotPasswordHandlers.httpWriter.WriteCustomHTTPError(w, http.StatusBadRequest, err.Error())
	//	return
	//}
	//if err == customError.ErrRecordNotFound {
	//	log.Logger(ctx).Error(err)
	//	forgotPasswordHandlers.httpWriter.WriteCustomHTTPError(w, http.StatusNotFound, err.Error())
	//	return
	//}
	//if err == customError.ErrInvalidUserType {
	//	log.Logger(ctx).Error(err)
	//	forgotPasswordHandlers.httpWriter.WriteCustomHTTPError(w, http.StatusBadRequest, err.Error())
	//	return
	//}
	//
	//if err != nil {
	//	log.Logger(ctx).Error(err)
	//	forgotPasswordHandlers.httpWriter.WriteCustomHTTPError(w, http.StatusInternalServerError, err.Error())
	//	return
	//}

	//write an http json resp
	forgotPasswordHandlers.httpWriter.WriteOKResponse(w, http.StatusCreated, "resp")

}

// ValidateOTP handler Function
func (forgotPasswordHandlers ForgotPasswordHandlers) ValidateOTP(w http.ResponseWriter,
	req *http.Request) {
	//
	//ctx := req.Context()
	//validateOTPReq := models.ValidateOtpReq{}
	//
	//err := forgotPasswordHandlers.httpReader.ReadInput(&validateOTPReq, req.Body)
	//if err != nil {
	//	log.Logger(ctx).Error(err)
	//	forgotPasswordHandlers.httpWriter.WriteHTTPError(w, http.StatusBadRequest)
	//	return
	//}
	//
	////call to service layer functions
	//resp, err := forgotPasswordHandlers.passwordService.ValidateOTP(ctx, validateOTPReq)
	//if err == customError.ErrOTPExpires {
	//	log.Logger(ctx).Error(err)
	//	forgotPasswordHandlers.httpWriter.WriteCustomHTTPError(w, http.StatusForbidden, err.Error())
	//	return
	//}
	//if err == customError.ErrInvalidUserType {
	//	log.Logger(ctx).Error(err)
	//	forgotPasswordHandlers.httpWriter.WriteCustomHTTPError(w, http.StatusBadRequest, err.Error())
	//	return
	//}
	//if err != nil {
	//	log.Logger(ctx).Error(err)
	//	forgotPasswordHandlers.httpWriter.WriteCustomHTTPError(w, http.StatusInternalServerError, err.Error())
	//	return
	//}

	//write an http json resp
	forgotPasswordHandlers.httpWriter.WriteOKResponse(w, http.StatusAccepted, "resp")

}

// UpdatePassword handler Function
func (forgotPasswordHandlers ForgotPasswordHandlers) UpdatePassword(w http.ResponseWriter,
	req *http.Request) {
	//
	//ctx := req.Context()
	//updatePassReq := models.UpdatePasswordReq{}
	//
	//err := forgotPasswordHandlers.httpReader.ReadInput(&updatePassReq, req.Body)
	//if err != nil {
	//	log.Logger(ctx).Error(err)
	//	forgotPasswordHandlers.httpWriter.WriteHTTPError(w, http.StatusBadRequest)
	//	return
	//}
	//
	////call to service layer functions
	//resp, err := forgotPasswordHandlers.passwordService.UpdatePassword(ctx, updatePassReq)
	//if err == customError.ErrInvalidPassword {
	//	log.Logger(ctx).Error(err)
	//	forgotPasswordHandlers.httpWriter.WriteCustomHTTPError(w, http.StatusBadRequest, err.Error())
	//	return
	//}
	//if err == customError.ErrInvalidUserType {
	//	log.Logger(ctx).Error(err)
	//	forgotPasswordHandlers.httpWriter.WriteCustomHTTPError(w, http.StatusBadRequest, err.Error())
	//	return
	//}
	//if err != nil {
	//	log.Logger(ctx).Error(err)
	//	forgotPasswordHandlers.httpWriter.WriteCustomHTTPError(w, http.StatusInternalServerError, err.Error())
	//	return
	//}

	//write an http json resp
	forgotPasswordHandlers.httpWriter.WriteOKResponse(w, http.StatusAccepted, "resp")

}
