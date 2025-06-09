// profile and cpu files are binary files that must be analyzed with "go tool pprof file_path"
package profiles

import (
	"bytes"
	"net/http"
	"runtime/pprof"
	"time"

	"github.com/gorilla/mux"
)

// func HeapProfileHandler(w http.ResponseWriter, r *http.Request) {
// 	var buf bytes.Buffer
// 	if err := pprof.Lookup("allocs").WriteTo(&buf, 1); err != nil {
// 		http.Error(w, "Error obteniendo perfil", http.StatusInternalServerError)
// 		return
// 	}

// 	result := map[string]string{
// 		"profile": buf.String(),
// 	}

//		w.Header().Set("Content-Type", "application/json")
//		json.NewEncoder(w).Encode(result)
//	}
// func ProfileHandler(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	profileName := vars["type"]

// 	p := pprof.Lookup(profileName)
// 	if p == nil {
// 		http.Error(w, "Perfil no encontrado", http.StatusNotFound)
// 		return
// 	}

// 	var buf bytes.Buffer
// 	_ = p.WriteTo(&buf, 1)

// 	json.NewEncoder(w).Encode(map[string]string{
// 		"profile": buf.String(),
// 	})
// }

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	profileName := vars["type"]

	// CPU and profile requires a special treatment
	//CPU
	if profileName == "cpu" {
		var buf bytes.Buffer

		// Begin the profile
		if err := pprof.StartCPUProfile(&buf); err != nil {
			http.Error(w, "No se pudo iniciar el perfil CPU", http.StatusInternalServerError)
			return
		}

		// Profiling 10 seconds
		time.Sleep(10 * time.Second)

		// stop profiling
		pprof.StopCPUProfile()

		// download binary file cpu.pprof
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", `attachment; filename="cpu.pprof"`)
		w.Write(buf.Bytes())
		return
	}

	//Profile
	if profileName == "profile" {
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=prof.pprof")

		pprof.StartCPUProfile(w)
		time.Sleep(10 * time.Second)
		pprof.StopCPUProfile()
		return
	}

	// Another profiles (heap, goroutine, etc.)
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
