package error

import "net/http"

const (
	BadRequestErrorType ErrorType = iota + 1
	UnauthorizedErrorType
	ForbiddenErrorType
	NotFoundErrorType
	ConflictErrorType
	InternalServerErrorType
)

const (
	BadRequestErrorCode         = "ERR_IMGN_BAD_REQUEST_ERROR"
	UnauthorizedErrorCode       = "ERR_IMGN_UNAUTHORIZED_ERROR"
	UnauthorizedActionErrorCode = "ERR_IMGN_UNAUTHORIZED_ACTION_ERROR"
	ForbiddenErrorCode          = "ERR_IMGN_FORBIDDEN_ERROR"
	NotFoundErrorCode           = "ERR_IMGN_NOT_FOUND_ERROR"
	MethodNotAllowedErrorCode   = "ERR_IMGN_METHOD_NOT_ALLOWED_ERROR"
	RequestTimeoutErrorCode     = "ERR_IMGN_REQUEST_TIMEOUT_ERROR"
	ConflictErrorCode           = "ERR_IMGN_CONFLICT_ERROR"
	InternalServerErrorCode     = "ERR_IMGN_INTERNAL_SERVER_ERROR"
)

type ErrorType uint

func (errorType ErrorType) New(errorCode ErrorCode, errorMessage string, statusCode int) *IMGNError {
	return &IMGNError{
		ErrorType: errorType,
		ErrorResponse: ErrorResponse{
			ErrorCode:    errorCode,
			ErrorMessage: errorMessage,
		},
		StatusCode: statusCode,
	}
}

type ErrorCode string

type ErrorResponse struct {
	ErrorCode    ErrorCode `json:"error_code"`
	ErrorMessage string    `json:"error_message"`
}

type IMGNError struct {
	StatusCode    int
	ErrorType     ErrorType
	ErrorResponse ErrorResponse
}

var ErrorSQLConn = InternalServerErrorType.New(InternalServerErrorCode, "sql connection error", http.StatusInternalServerError)
var ErrorSQLRecordExist = ConflictErrorType.New(ConflictErrorCode, "record already exist", http.StatusConflict)
var ErrorRecordNotFound = NotFoundErrorType.New(NotFoundErrorCode, "record not found", http.StatusNotFound)
var ErrorURLParamNotFound = BadRequestErrorType.New(BadRequestErrorCode, "url parameter not found", http.StatusBadRequest)
var ErrorURLQueryParamNotFound = BadRequestErrorType.New(BadRequestErrorCode, "url query parameter not found", http.StatusBadRequest)
var ErrorPathVariableNotFound = BadRequestErrorType.New(BadRequestErrorCode, "url path variable not found", http.StatusBadRequest)
var ErrorQueryParamVariableNotInt = BadRequestErrorType.New(BadRequestErrorCode, "url query parameter value not a number", http.StatusBadRequest)
var ErrorPathVariableNotInt = BadRequestErrorType.New(BadRequestErrorCode, "url path variable value  not a number", http.StatusBadRequest)
var ErrorInvalidLoginDetails = UnauthorizedErrorType.New(UnauthorizedErrorCode, "invalid login details", http.StatusUnauthorized)
var ErrorUpdateMobileNumber = InternalServerErrorType.New(InternalServerErrorCode, "update mobile number", http.StatusInternalServerError)
var ErrorOTPExpires = ForbiddenErrorType.New(ForbiddenErrorCode, "your otp is expired", http.StatusForbidden)
var ErrorInvalidPassword = UnauthorizedErrorType.New(UnauthorizedErrorCode, "invalid password", http.StatusUnauthorized)
var ErrorInvalidMobileNumber = UnauthorizedErrorType.New(UnauthorizedErrorCode, "invalid mobile number", http.StatusUnauthorized)
var ErrorNotAuthorizedToUpdatePassword = UnauthorizedErrorType.New(UnauthorizedErrorCode, "not authorized to update password ", http.StatusUnauthorized)
var ErrorUnAuthorizedAction = UnauthorizedErrorType.New(UnauthorizedActionErrorCode, "not authorized for action ", http.StatusUnauthorized)
var ErrorEmailAlreadyExist = BadRequestErrorType.New(BadRequestErrorCode, "email already exist", http.StatusBadRequest)
var ErrorMobileNumberAlreadyExist = BadRequestErrorType.New(BadRequestErrorCode, "mobile number already exist", http.StatusBadRequest)
var ErrorUsernameAlreadyExist = BadRequestErrorType.New(BadRequestErrorCode, "username already exist", http.StatusBadRequest)
var ErrorInvalidUserType = BadRequestErrorType.New(BadRequestErrorCode, "invalid user name", http.StatusBadRequest)
var ErrorDuplicateEntry = BadRequestErrorType.New(BadRequestErrorCode, "DUPLICATE ENTRY ERROR", http.StatusBadRequest)

func BadRequestErrorFunc(errorMessage string) *IMGNError {
	return BadRequestErrorType.New(BadRequestErrorCode, errorMessage, http.StatusBadRequest)
}

func InternalServerErrorFunc(errorMessage string) *IMGNError {
	return BadRequestErrorType.New(BadRequestErrorCode, errorMessage, http.StatusBadRequest)
}

func InvalidToTokenErrorFunc(errorMessage string) *IMGNError {
	return BadRequestErrorType.New(UnauthorizedErrorCode, errorMessage, http.StatusUnauthorized)
}
