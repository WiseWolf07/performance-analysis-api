package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"performance-analysis-api/internal/router"
)

func main() {

	// func () {
	// 	fmt.Println("pprof disponible en http://localhost:6060/debug/pprof/")
	// 	log.Println(http.ListenAndServe("localhost:6060", nil))
	// }()
	// profiles.Profiler()
	r := router.NewRouter()

	fmt.Println("Servidor escuchando en el puerto 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))

}
