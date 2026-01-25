package validator

import (
	"context"
	"github.com/go-playground/validator"
	imgnErr "github.com/iamgenii/error"
	log "github.com/iamgenii/logs"
)

type RequestValidator interface {
	ValidateReq(context context.Context, validateReq interface{}) *imgnErr.IMGNError
}

type requestValidator struct {
	validate *validator.Validate
}

func NewRequestValidator() RequestValidator {
	return requestValidator{validate: validator.New()}
}

func (requestValidator requestValidator) ValidateReq(context context.Context, validateReq interface{}) *imgnErr.IMGNError {
	err := requestValidator.validate.Struct(validateReq)
	if err != nil {
		log.Logger(context).Error("RequestValidator.ValidateReq: Error in request validation. Error: ",err)
		return imgnErr.BadRequestErrorFunc(err.Error())
	}
	return nil
}
