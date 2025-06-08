package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"performance-analysis-api/internal/metrics"
)

func GetSystemMetrics(res http.ResponseWriter, req *http.Request) {
	systemMetrics := metrics.GetSystemMetrics()

	res.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(res).Encode(systemMetrics); err != nil {
		http.Error(res, "Error al codificar la respuesta JSON", http.StatusInternalServerError)
		fmt.Printf("Error al codificar la respuesta JSON: %v\n", err)
	}
}
