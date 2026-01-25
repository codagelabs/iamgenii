package request

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"mime"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
	"strings"

	http2 "github.com/iamgenii/utils/http"
	"github.com/iamgenii/utils/http/client"
	"github.com/go-playground/validator"
)

type HTTPRequest interface {
	WithJsonBody(requestModel interface{}) HTTPRequest
	WithJsonBodyNoEscapeHTML(requestModel interface{}) HTTPRequest
	WithXmlBody(requestModel interface{}) HTTPRequest
	WithFromURLEncoded(formData map[string]interface{}) HTTPRequest
	WithContext(context context.Context) HTTPRequest
	WithBasicAuth(username string, password string) HTTPRequest
	WithJWTAuth(token string) HTTPRequest
	WithCustomValidator(validate *validator.Validate) HTTPRequest
	WithOauth(token string) HTTPRequest
	ResponseAs(responseModel interface{}) HTTPRequest
	ResponseStatusCodeAs(httpStatusCode *int) HTTPRequest
	ResponseCookiesAs(cookies *[]*http.Cookie) HTTPRequest
	ResponseHeadersAs(respHeaders *map[string][]string) HTTPRequest
	AddHeaders(key string, value string) HTTPRequest
	AddQueryParameters(queryParam map[string]string) HTTPRequest
	AddCookies(cookies *http.Cookie) HTTPRequest
	GET(url string) error
	POST(url string) error
	PUT(url string) error
	PATCH(url string) error
	DELETE(url string) error
	Error() error
}

var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

func escapeQuotes(s string) string {
	return quoteEscaper.Replace(s)
}

type httpRequest struct {
	responseStatusCode *int
	responseModel      interface{}
	responseCookies    *[]*http.Cookie
	requestModel       interface{}
	headers            map[string]string
	cookies            []*http.Cookie
	httpClient         client.HttpClient
	context            context.Context
	forWordCookies     bool
	validate           *validator.Validate
	requestByte        []byte
	requestBuildError  error
	queryParam         map[string]string
	responseHeader     *map[string][]string
}

func (r httpRequest) ResponseHeadersAs(respHeaders *map[string][]string) HTTPRequest {
	r.responseHeader = respHeaders
	return r
}

func (r httpRequest) WithJsonBody(requestModel interface{}) HTTPRequest {
	r.requestModel = requestModel
	r.headers["Content-Type"] = "application/json"
	requestBytes, err := json.Marshal(r.requestModel)
	if err != nil {
		r.requestBuildError = err
	}
	r.requestByte = requestBytes
	return r
}

func (r httpRequest) WithJsonBodyNoEscapeHTML(requestModel interface{}) HTTPRequest {
	r.requestModel = requestModel
	r.headers["Content-Type"] = "application/json"

	var buffer bytes.Buffer
	encoder := json.NewEncoder(&buffer)
	encoder.SetEscapeHTML(false)
	encoder.Encode(r.requestModel)
	err := encoder.Encode(r.requestModel)
	if err != nil {
		r.requestBuildError = err
	}
	r.requestByte = buffer.Bytes()
	return r
}

func (r httpRequest) WithXmlBody(requestModel interface{}) HTTPRequest {
	r.requestModel = requestModel
	r.headers["Content-Type"] = "application/json"
	requestBytes, err := xml.Marshal(r.requestModel)
	if err != nil {
		r.requestBuildError = err
	}
	r.requestByte = requestBytes
	return r
}

func (r httpRequest) WithFromURLEncoded(formData map[string]interface{}) HTTPRequest {
	r.requestModel = formData
	bodyBuffer := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuffer)
	for key, value := range formData {
		switch value.(type) {
		case multipart.FileHeader:
			fileHeader := value.(multipart.FileHeader)
			file, err := fileHeader.Open()
			if err != nil {
				r.requestBuildError = err
			}
			header := make(textproto.MIMEHeader)
			header.Set("Content-Disposition", fmt.Sprintf(`from-data; name="%s"; filename="%s"`, escapeQuotes(key), escapeQuotes(fileHeader.Filename)))
			header.Set("Content-Type", fileHeader.Header.Get("Content-Type"))
			formfileWrier, err := bodyWriter.CreatePart(header)
			if err != nil {
				r.requestBuildError = err
			}
			content, err := ioutil.ReadAll(file)
			if err != nil {
				r.requestBuildError = err
			}
			_, err = formfileWrier.Write(content)
			if err != nil {
				r.requestBuildError = err
			}
		case string:
			err := bodyWriter.WriteField(key, value.(string))
			if err != nil {
				r.requestBuildError = err
			}
		case []byte:
			data := value.([]byte)
			header := make(textproto.MIMEHeader)
			header.Set("Content-Disposition", fmt.Sprintf(`from-data; name="%s"`, escapeQuotes(key)))
			header.Set("Content-Type", "application/json")
			formfileWrier, err := bodyWriter.CreatePart(header)
			if err != nil {
				r.requestBuildError = err
			}
			_, err = formfileWrier.Write(data)
			if err != nil {
				r.requestBuildError = err
			}
		case http2.FileUpload:
			fileContent := value.(http2.FileUpload)
			file, err := os.Open(fileContent.FilePath)
			if err != nil {
				r.requestBuildError = err
			}
			formfileWrier, err := bodyWriter.CreateFormFile(key, fileContent.FileName)
			if err != nil {
				r.requestBuildError = err
			}
			content, err := ioutil.ReadAll(file)
			if err != nil {
				r.requestBuildError = err
			}
			_, err = formfileWrier.Write(content)
			if err != nil {
				r.requestBuildError = err
			}
			err = file.Close()
			if err != nil {
				r.requestBuildError = err
			}
		default:
			r.requestBuildError = errors.New("Invalid request type: only multipart, files and string are supported ")
		}
	}
	err := bodyWriter.Close()
	if err != nil {
		r.requestBuildError = err
	}
	r.headers["Content-Type"] = bodyWriter.FormDataContentType()
	r.requestByte = bodyBuffer.Bytes()
	return r
}

func (r httpRequest) WithContext(context context.Context) HTTPRequest {
	r.context = context
	return r
}

func (r httpRequest) WithCustomValidator(validate *validator.Validate) HTTPRequest {
	r.validate = validate
	return r
}
func (r httpRequest) WithBasicAuth(username, password string) HTTPRequest {
	r.headers["Authorization"] = "Basic " + r.basicAuth(username, password)
	return r
}

func (r httpRequest) WithJWTAuth(token string) HTTPRequest {
	r.headers["Authorization"] = "Bearer " + token
	return r
}

func (r httpRequest) WithOauth(token string) HTTPRequest {
	r.headers["Authorization"] = "Bearer " + token
	return r
}

func (r httpRequest) ResponseAs(resp interface{}) HTTPRequest {
	r.responseModel = resp
	return r
}

func (r httpRequest) ResponseStatusCodeAs(httpStatusCode *int) HTTPRequest {
	r.responseStatusCode = httpStatusCode
	return r
}

func (r httpRequest) ResponseCookiesAs(cookies *[]*http.Cookie) HTTPRequest {
	r.responseCookies = cookies
	return r
}

func (r httpRequest) AddHeaders(key, value string) HTTPRequest {
	r.headers[key] = value
	return r
}

func (r httpRequest) AddQueryParameters(queryParam map[string]string) HTTPRequest {
	r.queryParam = queryParam
	return r
}

func (r httpRequest) AddCookies(cookies *http.Cookie) HTTPRequest {
	r.cookies = append(r.cookies, cookies)
	return r
}

func (r httpRequest) GET(url string) error {
	r.makeRequest("GET", url)
	return r.makeRequest("GET", url)
}

func (r httpRequest) POST(url string) error {
	return r.makeRequest("POST", url)
}

func (r httpRequest) PUT(url string) error {
	return r.makeRequest("PUT", url)
}

func (r httpRequest) PATCH(url string) error {
	return r.makeRequest("PATCH", url)
}

func (r httpRequest) DELETE(url string) error {
	return r.makeRequest("DELETE", url)
}

func (r httpRequest) Error() error {
	return r.requestBuildError
}
func (r httpRequest) makeRequest(method, url string) error {
	if r.requestBuildError != nil {
		return r.requestBuildError
	}
	httpRequest, reqErr := http.NewRequest(method, url, bytes.NewBuffer(r.requestByte))
	if reqErr != nil {
		return r.requestBuildError
	}
	for key, value := range r.headers {
		httpRequest.Header.Add(key, value)
	}

	if r.context != nil {
		httpRequest = httpRequest.WithContext(r.context)
	}
	query := httpRequest.URL.Query()
	for paramKey, paramValue := range r.headers {
		query.Add(paramKey, paramValue)
	}
	httpRequest.URL.RawQuery = query.Encode()

	for _, cookie := range r.cookies {
		httpRequest.AddCookie(cookie)
	}
	response, httpErr := r.httpClient.Do(httpRequest)
	if httpErr != nil {
		return httpErr
	}
	if response != nil && r.responseStatusCode != nil {
		*r.responseStatusCode = response.StatusCode
	}
	if response != nil && r.responseHeader != nil {
		*r.responseHeader = response.Header
	}

	if response.StatusCode < 200 && response.StatusCode >= 300 {
		errRespBytes, readErr := ioutil.ReadAll(response.Body)
		if readErr != nil {
			return readErr
		}
		return errors.New(string(errRespBytes))
	}
	if r.responseModel != nil {
		err := r.processResponseModel(response)
		if err != nil {
			return err
		}
	}
	if r.responseCookies != nil {
		*r.responseCookies = response.Cookies()
	}
	closeErr := response.Body.Close()
	if closeErr != nil {
		return closeErr
	}
	return nil
}

func (r httpRequest) basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func (r httpRequest) processResponseModel(resp *http.Response) error {
	if twoDByteArrayPtr, is2DArrayErr := r.responseModel.(*[][]byte); is2DArrayErr {
		data, err := r.processMultiFormResponse(resp)
		if err != nil {
			return err
		}
		*twoDByteArrayPtr = data
	} else {
		body, readErr := ioutil.ReadAll(resp.Body)
		if readErr != nil {
			return readErr
		}
		if strPtr, isStr := r.responseModel.(*string); isStr {
			*strPtr = string(body)
		} else if byteArrPtr, isByteArrPtr := r.responseModel.(*[]byte); isByteArrPtr {
			*byteArrPtr = body
		} else {
			var unmarsharErr error
			contentType := resp.Header.Get("Content-Type")
			if strings.Contains(contentType, "application/xml") {
				unmarsharErr = xml.Unmarshal(body, r.responseModel)
			}
			if strings.Contains(contentType, "application/json") {
				unmarsharErr = json.Unmarshal(body, r.responseModel)
			}
			if unmarsharErr != nil {
				return unmarsharErr
			}
			//validatorErr:=r.validateR
		}
	}
	return nil
}

func (r httpRequest) processMultiFormResponse(resp *http.Response) ([][]byte, error) {
	_, param, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
	if err != nil {
		return nil, err
	}
	data := make([][]byte, 0)
	reader := multipart.NewReader(resp.Body, param["boundary"])
	for part, err := reader.NextPart(); err == nil; part, err = reader.NextPart() {
		buf, err := ioutil.ReadAll(part)
		if err != nil {
			return nil, err
		}
		data = append(data, buf)

	}
	return data, nil
}

type HttpRequestBuilder interface {
	NewRequest() HTTPRequest
}

type httpRequestBuilder struct {
	httpClient client.HttpClient
}

func (builder httpRequestBuilder) NewRequest() HTTPRequest {
	return httpRequest{
		headers:    map[string]string{},
		cookies:    []*http.Cookie{},
		httpClient: builder.httpClient,
	}
}
func NewHttpRequestBuilder(httpClient client.HttpClient) HttpRequestBuilder {
	return httpRequestBuilder{httpClient: httpClient}
}
