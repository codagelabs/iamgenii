package handlers

import (
	log "github.com/iamgenii/logs"
	"github.com/iamgenii/utils/http_utils"

	"net/http"

	"github.com/iamgenii/models"
	service "github.com/iamgenii/svc/cities/pkg/v1/services"
)

// CitiesHandlers for handler Functions
type CitiesHandlers struct {
	citiesSvc  service.CitiesService
	httpReader http_utils.HTTPReader
	httpWriter http_utils.HTTPWriter
}

// NewCitiesHandler inits dependencies for graphQL and Handlers
func NewCitiesHandler(citiesService service.CitiesService, httpReader http_utils.HTTPReader,
	httpWriter http_utils.HTTPWriter) *CitiesHandlers {
	return &CitiesHandlers{
		citiesSvc:  citiesService,
		httpWriter: httpWriter,
		httpReader: httpReader,
	}
}

// CreateCities godoc
// @Summary      Create a new city
// @Description  Create a new city entry
// @Tags         cities
// @Accept       json
// @Produce      json
// @Param        city  body      models.Cities  true  "City object"
// @Success      201   {object}  models.HTTPResp
// @Failure      400   {object}  models.HTTPErrResp
// @Failure      500   {object}  models.HTTPErrResp
// @Router       /cities/ [post]
func (citiesHandlers CitiesHandlers) CreateCities(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	categories := models.Cities{}
	readerErr := citiesHandlers.httpReader.ReadInput(&categories, req.Body)
	if readerErr != nil {
		log.Logger(ctx).Error("CitiesHandlers.CreateCities: Error in read request body. Error: ", readerErr)
		citiesHandlers.httpWriter.WriteHTTPError(w, readerErr)
		return
	}
	resp, err := citiesHandlers.citiesSvc.CreateCities(ctx, categories)
	if err != nil {
		log.Logger(ctx).Error("CitiesHandlers.CreateCities: Error raised by service: Error: ", err)
		citiesHandlers.httpWriter.WriteHTTPError(w, err)
		return
	}
	citiesHandlers.httpWriter.WriteOKResponse(w, http.StatusCreated, resp)

}

// GetCities godoc
// @Summary      Get all cities
// @Description  Retrieve all cities
// @Tags         cities
// @Accept       json
// @Produce      json
// @Success      200  {array}   models.Cities
// @Failure      500  {object}  models.HTTPErrResp
// @Router       /cities/ [get]
func (citiesHandlers CitiesHandlers) GetCities(w http.ResponseWriter,
	req *http.Request) {
	ctx := req.Context()
	resp, err := citiesHandlers.citiesSvc.GetCities(ctx)
	if err != nil {
		log.Logger(ctx).Error("CitiesHandlers.CreateCities: Error raised by service: Error: ", err)
		citiesHandlers.httpWriter.WriteHTTPError(w, err)
		return
	}
	citiesHandlers.httpWriter.WriteOKResponse(w, http.StatusOK, resp)
}
