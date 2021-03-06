package server

import (
	"fmt"
	"metrics-collector/cf"
	"metrics-collector/config"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pivotal-golang/lager"
	"github.com/tedsuo/ifrit"
	"github.com/tedsuo/ifrit/http_server"
)

const PATH_MEMORY_METRIC = "/v1/apps/{appid}/metrics/memory"
const ROUTE_NAME_MEMORY_METRIC = "memory-metric"

type VarsFunc func(w http.ResponseWriter, r *http.Request, vars map[string]string)

func (vh VarsFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	vh(w, r, vars)
}

func NewServer(logger lager.Logger, conf config.ServerConfig, cfc cf.CfClient, consumer NoaaConsumer) ifrit.Runner {
	mmh := NewMemoryMetricHandler(logger, cfc, consumer)

	r := mux.NewRouter()
	r.Methods("GET").Path(PATH_MEMORY_METRIC).Handler(VarsFunc(mmh.GetMemoryMetric)).Name(ROUTE_NAME_MEMORY_METRIC)

	addr := fmt.Sprintf("0.0.0.0:%d", conf.Port)
	return http_server.New(addr, r)
}
