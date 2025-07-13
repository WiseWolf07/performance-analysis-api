package profiles

import (
	"encoding/json"
	"net/http"
	"runtime"
	"sync"
	"time"
)

type Sample struct {
	Timestamp     int64   `json:"timestamp"`
	Fragmentation float64 `json:"fragmentation"`
	HeapSys       uint64  `json:"heap_sys"`
	HeapInuse     uint64  `json:"heap_inuse"`
}

var (
	fragData       []Sample
	fragMutex      sync.Mutex
	maxFragSamples = 50
)

func StartCollector() {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		var mem runtime.MemStats
		runtime.ReadMemStats(&mem)

		frag := 0.0
		if mem.HeapSys > 0 {
			frag = float64(mem.HeapIdle) / float64(mem.HeapSys) * 100
		}

		s := Sample{
			Timestamp:     time.Now().UnixMilli(),
			Fragmentation: frag,
			HeapSys:       mem.HeapSys,
			HeapInuse:     mem.HeapInuse,
		}

		fragMutex.Lock()
		fragData = append(fragData, s)
		if len(fragData) > maxFragSamples {
			fragData = fragData[1:]
		}
		fragMutex.Unlock()
	}
}

func GetHistory(w http.ResponseWriter, r *http.Request) {
	fragMutex.Lock()
	defer fragMutex.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(fragData)
}
