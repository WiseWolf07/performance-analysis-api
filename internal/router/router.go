package router

import (
	"performance-analysis-api/internal/handlers"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/metrics/system", handlers.GetSystemMetrics).Methods(("GET"))

	return r
}
