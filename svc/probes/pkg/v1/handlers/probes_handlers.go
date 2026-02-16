package handlers

import (
	"net/http"

	log "github.com/iamgenii/logs"
	"github.com/iamgenii/svc/probes/pkg/v1/services"
	"github.com/iamgenii/utils/http_utils"
)

// ProbesHandlers for handler Functions
type ProbesHandlers struct {
	probesSvc  services.ProbesService
	httpWriter http_utils.HTTPWriter
}

// NewProbesHandlers inits dependencies for Handlers
func NewProbesHandlers(probesSvc services.ProbesService, httpWriter http_utils.HTTPWriter) *ProbesHandlers {
	return &ProbesHandlers{
		probesSvc:  probesSvc,
		httpWriter: httpWriter,
	}
}

// Liveness godoc
// @Summary      Liveness probe
// @Description  Check if the application is running
// @Tags         probes
// @Produce      json
// @Success      200  {string}  string  "OK"
// @Router       /live [get]
func (h *ProbesHandlers) Liveness(w http.ResponseWriter, r *http.Request) {
	h.httpWriter.WriteOKResponse(w, http.StatusOK, "OK")
}

// Readiness godoc
// @Summary      Readiness probe
// @Description  Check if the application is ready to serve traffic
// @Tags         probes
// @Produce      json
// @Success      200  {string}  string  "OK"
// @Failure      503  {object}  models.HTTPErrResp
// @Router       /ready [get]
func (h *ProbesHandlers) Readiness(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	if err := h.probesSvc.HealthCheck(ctx); err != nil {
		log.Logger(ctx).Error("ProbesHandlers.Readiness: Error raised by service. Error: ", err)
		h.httpWriter.WriteCustomHTTPError(w, http.StatusServiceUnavailable, "Service Unavailable")
		return
	}
	h.httpWriter.WriteOKResponse(w, http.StatusOK, "OK")
}

// Health godoc
// @Summary      Health check
// @Description  Check application health status
// @Tags         probes
// @Produce      json
// @Success      200  {string}  string  "OK"
// @Failure      503  {object}  models.HTTPErrResp
// @Router       /health [get]
func (h *ProbesHandlers) Health(w http.ResponseWriter, r *http.Request) {
	h.Readiness(w, r)
}
