package main

import (
	"log"
	"net/http"

	"liatdong/internal/config"
)

func main() {
	viperConfig := config.NewViperConfig()
	_ = config.NewLogger(viperConfig)
	_ = config.NewPgxConnection(viperConfig)

	log.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
