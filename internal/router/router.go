package router

import (
	"performance-analysis-api/internal/handlers"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/metrics/system", handlers.GetSystemMetrics).Methods(("GET"))

	return router
}
