package handlers

import (
	"encoding/json"
	"net/http"
)

func GetSystemMetrics(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"message": "Obteniendo métricas del sistema..."}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
