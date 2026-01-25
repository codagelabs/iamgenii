package http_utils

import (
	"encoding/json"
	"net/http"

	custErr "github.com/iamgenii/error"
	"github.com/iamgenii/models"
)

var statusText = map[int]string{
	http.StatusContinue:                      "Continue",
	http.StatusSwitchingProtocols:            "Switching Protocols",
	http.StatusProcessing:                    "Processing",
	http.StatusOK:                            "OK",
	http.StatusCreated:                       "Created",
	http.StatusAccepted:                      "Accepted",
	http.StatusNonAuthoritativeInfo:          "Non-Authoritative Information",
	http.StatusNoContent:                     "No Content",
	http.StatusResetContent:                  "Reset Content",
	http.StatusPartialContent:                "Partial Content",
	http.StatusMultiStatus:                   "Multi-Status",
	http.StatusAlreadyReported:               "Already Reported",
	http.StatusIMUsed:                        "IM Used",
	http.StatusMultipleChoices:               "Multiple Choices",
	http.StatusMovedPermanently:              "Moved Permanently",
	http.StatusFound:                         "Found",
	http.StatusSeeOther:                      "See Other",
	http.StatusNotModified:                   "Not Modified",
	http.StatusUseProxy:                      "Use Proxy",
	http.StatusTemporaryRedirect:             "Temporary Redirect",
	http.StatusPermanentRedirect:             "Permanent Redirect",
	http.StatusBadRequest:                    "Bad Request",
	http.StatusUnauthorized:                  "Unauthorized",
	http.StatusPaymentRequired:               "Payment Required",
	http.StatusForbidden:                     "Forbidden",
	http.StatusNotFound:                      "Not Found",
	http.StatusMethodNotAllowed:              "Method Not Allowed",
	http.StatusNotAcceptable:                 "Not Acceptable",
	http.StatusProxyAuthRequired:             "Proxy Authentication Required",
	http.StatusRequestTimeout:                "Request Timeout",
	http.StatusConflict:                      "Conflict",
	http.StatusGone:                          "Gone",
	http.StatusLengthRequired:                "Length Required",
	http.StatusPreconditionFailed:            "Precondition Failed",
	http.StatusRequestEntityTooLarge:         "Request Entity Too Large",
	http.StatusRequestURITooLong:             "Request URI Too Long",
	http.StatusUnsupportedMediaType:          "Unsupported Media Type",
	http.StatusRequestedRangeNotSatisfiable:  "Requested Range Not Satisfiable",
	http.StatusExpectationFailed:             "Expectation Failed",
	http.StatusTeapot:                        "I'm a teapot",
	http.StatusMisdirectedRequest:            "Misdirected Request",
	http.StatusUnprocessableEntity:           "Unprocessable Entity",
	http.StatusLocked:                        "Locked",
	http.StatusFailedDependency:              "Failed Dependency",
	http.StatusTooEarly:                      "Too Early",
	http.StatusUpgradeRequired:               "Upgrade Required",
	http.StatusPreconditionRequired:          "Precondition Required",
	http.StatusTooManyRequests:               "Too Many Requests",
	http.StatusRequestHeaderFieldsTooLarge:   "Request Header Fields Too Large",
	http.StatusUnavailableForLegalReasons:    "Unavailable For Legal Reasons",
	http.StatusInternalServerError:           "Internal Server Error",
	http.StatusNotImplemented:                "Not Implemented",
	http.StatusBadGateway:                    "Bad Gateway",
	http.StatusServiceUnavailable:            "Service Unavailable",
	http.StatusGatewayTimeout:                "Gateway Timeout",
	http.StatusHTTPVersionNotSupported:       "HTTP Version Not Supported",
	http.StatusVariantAlsoNegotiates:         "Variant Also Negotiates",
	http.StatusInsufficientStorage:           "Insufficient Storage",
	http.StatusLoopDetected:                  "Loop Detected",
	http.StatusNotExtended:                   "Not Extended",
	http.StatusNetworkAuthenticationRequired: "Network Authentication Required",
}

type HTTPWriter interface {
	WriteOKResponse(writer http.ResponseWriter, statusCode int, responseData interface{})
	WriteHTTPError(writer http.ResponseWriter, error *custErr.IMGNError)
	WriteCustomHTTPError(writer http.ResponseWriter, statusCode int, customErrMessage string)
}
type httpWriter struct {
}

func NewHTTPWriter() HTTPWriter {
	return &httpWriter{}
}

// WriteOKResponse as a standard JSON response with StatusOK
func (httpWriter httpWriter) WriteOKResponse(writer http.ResponseWriter, statusCode int, data interface{}) {

	resp := models.HTTPResp{}
	resp.Message = statusText[statusCode]
	resp.Status = statusCode
	resp.Data = data

	//status code setting
	writer.WriteHeader(statusCode)

	//encode error message and send it to api response
	//new encoder
	encoder := json.NewEncoder(writer)

	//if error occurs return http error
	if err := encoder.Encode(resp); err != nil {
		http.Error(writer, "error in response parsing", http.StatusInternalServerError)
	}
}

// WriteHTTPError return HTTP Error Message
func (httpWriter httpWriter) WriteHTTPError(writer http.ResponseWriter, error *custErr.IMGNError) {

	//setting of content type done here
	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")

	//write an status code
	writer.WriteHeader(error.StatusCode)

	//encode error message and send it to api response
	//new encoder
	encoder := json.NewEncoder(writer)
	//if error occurs return http error
	if err := encoder.Encode(error.ErrorResponse); err != nil {
		http.Error(writer, error.ErrorResponse.ErrorMessage, error.StatusCode)
	}
}

// WriteCustomHTTPError  return the response in json format with status code
func (httpWriter httpWriter) WriteCustomHTTPError(writer http.ResponseWriter, statusCode int,
	CustomErrMessage string) {

	//setting of content type done here
	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")

	//error message extract statusText map
	errorMsg := statusText[statusCode]

	//set Params to error message
	resp := models.HTTPErrResp{}
	resp.Message = errorMsg + " : " + CustomErrMessage
	resp.Status = statusCode

	//write an status code
	writer.WriteHeader(statusCode)

	//encode error message and send it to api response
	//new encoder
	encoder := json.NewEncoder(writer)

	//if error occurs return http error
	if err := encoder.Encode(resp); err != nil {
		http.Error(writer, CustomErrMessage, http.StatusInternalServerError)
	}

}
