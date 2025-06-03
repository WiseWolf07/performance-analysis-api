package main

import (
	"fmt"
	"log"
	"net/http"

	"performance-analysis-api/internal/router"
)

func main() {
	r := router.NewRouter()

	fmt.Println("Servidor escuchando en el puerto 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
