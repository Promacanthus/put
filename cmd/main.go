package main

import (
	"net/http"

	"github.com/promacanthus/put/pkg/log"
	"github.com/promacanthus/put/pkg/put"
	"go.uber.org/zap"
)

func main() {
	defer log.Logger.Sync()

	http.Handle("/webhook", put.NewServer())
	log.Logger.Info("Webhook Server started on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Logger.Fatal("Running Webhook Server failed, err: %v", zap.Error(err))
	}
}
