package main

import (
	"log"
	"net/http"

	"github.com/promacanthus/put/pkg/put"
)

func main() {
	http.Handle("/webhook", put.NewServer())
	log.Println("Webhook Server started on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Running Webhook Server failed, err: %v", err)
	}
}
