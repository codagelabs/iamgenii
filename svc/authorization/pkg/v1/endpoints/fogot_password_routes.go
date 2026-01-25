package endpoints

import (
	handler "github.com/iamgenii/svc/authorization/pkg/v1/handlers"
	"github.com/gorilla/mux"
)

// NewForgotPasswordRoutes All Application Routes Are defend Here
func NewForgotPasswordRoutes(router *mux.Router, handler *handler.ForgotPasswordHandlers) {

	router.HandleFunc("/user/verify-send-otp/", handler.VerifyAndSendOTP).Methods("POST", "OPTIONS")

	router.HandleFunc("/user/validate-otp/", handler.ValidateOTP).Methods("POST", "OPTIONS")

	router.HandleFunc("/user/update-password/", handler.UpdatePassword).Methods("POST", "OPTIONS")

}
