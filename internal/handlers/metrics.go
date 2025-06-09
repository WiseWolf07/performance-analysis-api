package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"performance-analysis-api/internal/metrics"
)

func GetSystemMetrics(res http.ResponseWriter, req *http.Request) {
	startTime := time.Now()

	systemMetrics := metrics.GetSystemMetrics()

	latency := time.Since(startTime)
	systemMetrics.LatencyMs = float64(latency.Nanoseconds()) / float64(time.Millisecond)

	res.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(res).Encode(systemMetrics); err != nil {
		http.Error(res, "Error al codificar la respuesta JSON", http.StatusInternalServerError)
		fmt.Printf("Error al codificar la respuesta JSON: %v\n", err)
	}
}
