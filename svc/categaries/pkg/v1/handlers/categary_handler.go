package handlers

import (
	"fmt"

	"github.com/iamgenii/utils/http_utils"

	log "github.com/iamgenii/logs"

	"net/http"

	"github.com/iamgenii/models"
	service "github.com/iamgenii/svc/categaries/pkg/v1/services"
)

var (
	//CategoriesID updates all the details
	CategoriesID = "categoriesId"
)

// CategoriesHandlers for handler Functions
type CategoriesHandlers struct {
	categoriesSvc service.CategoriesService
	httpReader    http_utils.HTTPReader
	httpWriter    http_utils.HTTPWriter
}

// NewCategoriesHandlerImpl inits dependencies for graphQL and Handlers
func NewCategoriesHandlerImpl(categoriesSvc service.CategoriesService, httpReader http_utils.HTTPReader, httpWriter http_utils.HTTPWriter) *CategoriesHandlers {
	return &CategoriesHandlers{
		categoriesSvc: categoriesSvc,
		httpWriter:    httpWriter,
		httpReader:    httpReader,
	}
}

// CreateCategories handler Function
func (categoriesHandlers CategoriesHandlers) CreateCategories(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	categories := models.Categories{}

	readerErr := categoriesHandlers.httpReader.ReadInput(&categories, req.Body)
	if readerErr != nil {
		log.Logger(ctx).Error("CategoriesHandlers.CreateCategories: Error in read input request body. Error :", readerErr)
		categoriesHandlers.httpWriter.WriteHTTPError(w, readerErr)
		return
	}

	resp, err := categoriesHandlers.categoriesSvc.CreateCategories(ctx, categories)
	if err != nil {
		log.Logger(ctx).Error("CategoriesHandlers.CreateCategories: Error got from service . Error :", err)
		categoriesHandlers.httpWriter.WriteHTTPError(w, err)
		return
	}

	//write an http json resp
	categoriesHandlers.httpWriter.WriteOKResponse(w, http.StatusCreated, resp)

}

// GetCategoriesByID handler Function
func (categoriesHandlers CategoriesHandlers) GetCategoriesByID(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	id, readerErr := categoriesHandlers.httpReader.GetURLParam(req, CategoriesID)
	if readerErr != nil {
		log.Logger(ctx).Error("CategoriesHandlers.GetCategoriesByID: Error in get Url parameter. Error: ", readerErr)
		categoriesHandlers.httpWriter.WriteHTTPError(w, readerErr)
		return
	}

	strID := fmt.Sprint(*id)
	resp, err := categoriesHandlers.categoriesSvc.GetCategoriesByID(ctx, strID)
	if err != nil {
		log.Logger(ctx).Error("CategoriesHandlers.GetCategoriesByID: Error got from service . Error :", err)
		categoriesHandlers.httpWriter.WriteHTTPError(w, err)
		return
	}
	categoriesHandlers.httpWriter.WriteOKResponse(w, http.StatusOK, resp)
}

// GetSubCategories handler Function
func (categoriesHandlers CategoriesHandlers) GetSubCategories(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	id, readerErr := categoriesHandlers.httpReader.GetURLParam(req, CategoriesID)
	if readerErr != nil {
		log.Logger(ctx).Error("CategoriesHandlers.GetSubCategories: Error in get Url parameter. Error: ", readerErr)
		categoriesHandlers.httpWriter.WriteHTTPError(w, readerErr)
		return
	}

	strID := fmt.Sprint(*id)
	resp, err := categoriesHandlers.categoriesSvc.GetSubCategoriesByID(ctx, strID)
	if err != nil {
		log.Logger(ctx).Error("CategoriesHandlers.GetSubCategories: Error got from service . Error :", err)
		categoriesHandlers.httpWriter.WriteHTTPError(w, err)
		return
	}

	categoriesHandlers.httpWriter.WriteOKResponse(w, http.StatusOK, resp)
}

// DeleteCategories handler Function
func (categoriesHandlers CategoriesHandlers) DeleteCategories(w http.ResponseWriter,
	req *http.Request) {
	ctx := req.Context()
	id, readerErr := categoriesHandlers.httpReader.GetURLParam(req, CategoriesID)
	if readerErr != nil {
		log.Logger(ctx).Error("CategoriesHandlers.DeleteCategories: Error in get Url parameter. Error: parameter name : ", readerErr)
		categoriesHandlers.httpWriter.WriteHTTPError(w, readerErr)
		return
	}
	strID := fmt.Sprint(*id)
	resp, err := categoriesHandlers.categoriesSvc.DeleteCategories(ctx, strID)
	if err != nil {
		log.Logger(ctx).Error("CategoriesHandlers.DeleteCategories: Error got from service . Error :", err)
		categoriesHandlers.httpWriter.WriteHTTPError(w, err)
		return
	}
	categoriesHandlers.httpWriter.WriteOKResponse(w, http.StatusOK, resp)
}
