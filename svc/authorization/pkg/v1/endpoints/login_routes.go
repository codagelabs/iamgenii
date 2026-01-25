package endpoints

import (
	handler "github.com/iamgenii/svc/authorization/pkg/v1/handlers"
	"github.com/gorilla/mux"
)

func NewAuthorizationRoutes(router *mux.Router, handler *handler.LoginHandler) {
	router.HandleFunc("/login/", handler.Login).Methods("POST", "OPTIONS")
	router.HandleFunc("/logout/", handler.Logout).Methods("POST", "OPTIONS")

}
