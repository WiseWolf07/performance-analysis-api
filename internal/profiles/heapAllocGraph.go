package profiles

import (
	"encoding/json"
	"net/http"
	"runtime"
	"sync"
	"time"
)

type MemStatsSample struct {
	Timestamp int64                  `json:"timestamp"`
	Data      map[string]interface{} `json:"data"`
}

var (
	Samples      []MemStatsSample
	SamplesMutex sync.Mutex
	maxSamples   = 100 // guarda solo las últimas 100 muestras
)

// Ejecuta esto en un goroutine al iniciar la app
func CollectMemStats() {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		var memStats runtime.MemStats
		runtime.ReadMemStats(&memStats)

		data := map[string]interface{}{
			"alloc":           memStats.Alloc,
			"total_alloc":     memStats.TotalAlloc,
			"sys":             memStats.Sys,
			"mallocs":         memStats.Mallocs,
			"frees":           memStats.Frees,
			"heap_alloc":      memStats.HeapAlloc,
			"heap_sys":        memStats.HeapSys,
			"heap_idle":       memStats.HeapIdle,
			"heap_inuse":      memStats.HeapInuse,
			"heap_objects":    memStats.HeapObjects,
			"num_gc":          memStats.NumGC,
			"next_gc":         memStats.NextGC,
			"gc_cpu_fraction": memStats.GCCPUFraction,
		}

		sample := MemStatsSample{
			Timestamp: time.Now().UnixMilli(),
			Data:      data,
		}

		SamplesMutex.Lock()
		Samples = append(Samples, sample)
		if len(Samples) > maxSamples {
			Samples = Samples[1:]
		}
		SamplesMutex.Unlock()
	}
}

// Handler para devolver el historial actual de métricas
func GetMemStatsHistory(w http.ResponseWriter, r *http.Request) {
	SamplesMutex.Lock()
	defer SamplesMutex.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Samples)
}
