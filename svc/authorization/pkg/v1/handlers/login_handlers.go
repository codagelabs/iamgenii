package handlers

import (
	log "github.com/iamgenii/logs"
	"github.com/iamgenii/models"
	"github.com/iamgenii/utils"
	"github.com/iamgenii/utils/http_utils"

	"net/http"

	service "github.com/iamgenii/svc/authorization/pkg/v1/services"
)

// LoginHandler for handler Functions
type LoginHandler struct {
	authorizationSvc service.LoginService
	httpReader       http_utils.HTTPReader
	httpWriter       http_utils.HTTPWriter
	cookies          utils.Cookies
}

// NewLoginHandler inits dependencies for graphQL and Handlers
func NewLoginHandler(authorizationService service.LoginService, httpReader http_utils.HTTPReader, httpWriter http_utils.HTTPWriter, cookies utils.Cookies) *LoginHandler {
	return &LoginHandler{
		authorizationSvc: authorizationService,
		httpReader:       httpReader,
		httpWriter:       httpWriter,
		cookies:          cookies,
	}
}

// LoginByEmailOrUsername handler Function
func (handler LoginHandler) Login(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	log.Logger(ctx).Debug("LoginHandler.LoginByEmailOrUsername: In handler func")
	authorization := models.LoginReq{}
	//TODO request should be encrypted Username and Password
	readerErr := handler.httpReader.ReadInput(&authorization, req.Body)
	if readerErr != nil {
		log.Logger(ctx).Error("LoginHandler.LoginByEmailOrUsername: Error In read Request Body", readerErr)
		handler.httpWriter.WriteHTTPError(w, readerErr)
		return
	}

	//call to service layer functions
	resp, svcErr := handler.authorizationSvc.LoginByEmailOrUsername(ctx, authorization)
	if svcErr != nil {
		log.Logger(ctx).Error("LoginHandler.LoginByEmailOrUsername: Error raised by service: Error: ", svcErr)
		handler.httpWriter.WriteHTTPError(w, svcErr)
		return
	}
	handler.cookies.SetTokenCookies(w, resp.Token, resp.UserType)
	w.WriteHeader(http.StatusNoContent)

}

// Logout handler Function
func (handler LoginHandler) Logout(w http.ResponseWriter,
	req *http.Request) {
	handler.cookies.DeleteTokenCookies(w)
	w.WriteHeader(http.StatusOK)

}
