package http_utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/iamgenii/configs"
	custErr "github.com/iamgenii/error"

	log "github.com/iamgenii/logs"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

type HTTPReader interface {
	ReadInput(input interface{}, reqBody io.ReadCloser) (imgnError *custErr.IMGNError)
	GetURLParam(request *http.Request, paramName string) (paramValue *uint64, imgnError *custErr.IMGNError)
	GetURLQueryParam(request *http.Request, paramName string, isRequired bool) (paramValue *int, imgnError *custErr.IMGNError)
	GetIDFromToken(req *http.Request) (Id *string, imgnError *custErr.IMGNError)
}

type httpReader struct {
	jwtConfig configs.JwtConfig
}

func NewHTTPReader(jwtConfig configs.JwtConfig) HTTPReader {
	return &httpReader{jwtConfig: jwtConfig}
}

func (httpReader httpReader) ReadInput(input interface{}, reqBody io.ReadCloser) *custErr.IMGNError {

	decoder := json.NewDecoder(reqBody)
	err := decoder.Decode(input)
	if err != nil {
		return custErr.BadRequestErrorFunc(err.Error())
	}
	return nil
}
func (httpReader httpReader) GetURLParam(request *http.Request, paramName string) (paramValue *uint64, imgnErr *custErr.IMGNError) {
	params := mux.Vars(request)
	id, ok := params[paramName]
	if !ok {
		return nil, custErr.ErrorURLParamNotFound
	}

	uintID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return nil, custErr.ErrorPathVariableNotInt
	}

	return &uintID, nil
}
func (httpReader httpReader) GetURLQueryParam(request *http.Request, paramName string, isRequired bool) (paramValue *int, imgnError *custErr.IMGNError) {

	strParam := "0"
	ctx := request.Context()
	keys := request.URL.Query()
	pages, ok := keys[paramName]
	if ok && strings.TrimSpace(pages[0]) != "" {
		strParam = pages[0]
	} else if isRequired {
		return nil, custErr.ErrorURLQueryParamNotFound
	}
	uintID, err := strconv.Atoi(strParam)
	if err != nil {
		log.Logger(ctx).Error(err)
		return nil, custErr.ErrorQueryParamVariableNotInt
	}
	return &uintID, nil
}
func (httpReader httpReader) GetIDFromToken(req *http.Request) (Id *string, imgnError *custErr.IMGNError) {
	ctx := req.Context()
	authorizationHeader := req.Header.Get("authorization")
	if authorizationHeader != "" {
		bearerToken := strings.Split(authorizationHeader, "Bearer")
		reqToken := bearerToken[1]
		reqToken = strings.TrimSpace(reqToken)
		token, err := jwt.Parse(reqToken, func(t *jwt.Token) (interface{}, error) {
			return []byte(httpReader.jwtConfig.GetJwtSecretKey()), nil
		})
		if err != nil {
			log.Logger(ctx).Error("Error in sectionId Parsing", err)
			return nil, custErr.InternalServerErrorFunc(err.Error())
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			log.Logger(ctx).Error("Error in sectionId Parsing", err)
			return nil, custErr.InternalServerErrorFunc(err.Error())
		}

		id, ok := claims["id"].(string)
		if !ok {
			return nil, custErr.InternalServerErrorFunc("Id is not found in token")
		}

		fmt.Println(id)
		return &id, nil
	}
	return nil, custErr.ErrorUnAuthorizedAction
}
