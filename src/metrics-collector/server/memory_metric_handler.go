package server

import (
	"encoding/json"
	"metrics-collector/cf"
	"metrics-collector/metrics"
	"net/http"

	"github.com/cloudfoundry-incubator/cf_http/handlers"
	"github.com/cloudfoundry/sonde-go/events"
	"github.com/pivotal-golang/lager"
)

const TOKEN_TYPE_BEARER = "bearer"

type MemoryMetricHandler struct {
	cfClient cf.CfClient
	logger   lager.Logger
	noaa     NoaaConsumer
}

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type NoaaConsumer interface {
	ContainerMetrics(appGuid string, authToken string) ([]*events.ContainerMetric, error)
}

func NewMemoryMetricHandler(logger lager.Logger, cfc cf.CfClient, consumer NoaaConsumer) *MemoryMetricHandler {
	return &MemoryMetricHandler{
		cfClient: cfc,
		noaa:     consumer,
		logger:   logger,
	}
}

func (h *MemoryMetricHandler) GetMemoryMetric(w http.ResponseWriter, r *http.Request, vars map[string]string) {
	appId := vars["appid"]

	w.Header().Set("Content-Type", "application/json")

	containerMetrics, err := h.noaa.ContainerMetrics(appId, TOKEN_TYPE_BEARER+" "+h.cfClient.GetTokens().AccessToken)
	if err != nil {
		h.logger.Error("noaa-get-container-metrics", err)

		handlers.WriteJSONResponse(w, http.StatusInternalServerError, ErrorResponse{
			Code:    "Interal-Server-Error",
			Message: "Error getting memory metrics from doppler"})
		return
	}

	h.logger.Debug("get-container-metrics-from-doppler", lager.Data{"container-metrics": containerMetrics})
	metric := metrics.GetMemoryMetricFromContainerMetrics(appId, containerMetrics)

	body, err := json.Marshal(metric)
	if err != nil {
		h.logger.Error("marshal-memory-metrics", err, lager.Data{"metric": metric})

		handlers.WriteJSONResponse(w, http.StatusInternalServerError, ErrorResponse{
			Code:    "Interal-Server-Error",
			Message: "Error getting memory metrics from doppler"})
		return
	}

	w.Write(body)
}
