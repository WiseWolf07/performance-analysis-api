package router

import (
	"net/http"
	"performance-analysis-api/internal/handlers"
	"performance-analysis-api/internal/profiles"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/metrics/system", handlers.GetSystemMetrics).Methods("GET")
	router.HandleFunc("/profile/{type}", profiles.ProfileHandler).Methods("GET")
	router.HandleFunc("/profiles/graph", profiles.GetMemStatsHistory).Methods("GET")
	router.HandleFunc("/profiles/fragmentation", profiles.GetHistory).Methods("GET")
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	go profiles.StartCollector()
	go profiles.CollectMemStats()

	return router

}
