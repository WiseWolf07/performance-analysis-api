package router

import (
	"performance-analysis-api/internal/handlers"
	"performance-analysis-api/internal/profiles"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/metrics/system", handlers.GetSystemMetrics).Methods(("GET"))
	// router.HandleFunc("/profile/heap", profiles.HeapProfileHandler).Methods("GET")
	router.HandleFunc("/profile/{type}", profiles.ProfileHandler)
	return router
}
