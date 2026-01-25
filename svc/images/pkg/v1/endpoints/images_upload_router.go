package endpoints

import (
	"github.com/iamgenii/svc/images/pkg/v1/handlers"
	"github.com/gorilla/mux"
)

func NewImageUploaderRoutes(router *mux.Router, handler *handlers.ImageHandlers) {
	router.HandleFunc("/images/upload/", handler.UploadImage).Methods("POST", "OPTIONS")
}
