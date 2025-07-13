package profiles

import (
	"bytes"
	"encoding/json"
	"net/http"
	"runtime"
	"runtime/pprof"
	"time"

	"github.com/gorilla/mux"
)

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	profileName := vars["type"]

	switch profileName {
	case "cpu":
		var buf bytes.Buffer
		if err := pprof.StartCPUProfile(&buf); err != nil {
			http.Error(w, "No se pudo iniciar el perfil CPU", http.StatusInternalServerError)
			return
		}

		time.Sleep(10 * time.Second)
		pprof.StopCPUProfile()

		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", `attachment; filename="cpu.pprof"`)
		w.Write(buf.Bytes())
		return

	case "profile":
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=prof.pprof")

		pprof.StartCPUProfile(w)
		time.Sleep(10 * time.Second)
		pprof.StopCPUProfile()
		return

	case "heapjson":
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

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
		return
	}

	// Otro perfil (heap, goroutine, threadcreate, block, mutex, etc.)
	p := pprof.Lookup(profileName)
	if p == nil {
		http.Error(w, "Perfil no encontrado", http.StatusNotFound)
		return
	}

	var buf bytes.Buffer
	if err := p.WriteTo(&buf, 1); err != nil {
		http.Error(w, "Error al generar el perfil", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Write(buf.Bytes())
}
