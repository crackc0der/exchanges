package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/crackc0der/exchanges/internal/exchange"
)

func main() {
	httpPort := "127.0.0.1:8080"
	timeout := 10
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	service := exchange.NewService()
	endpoint := exchange.NewEndpoint(logger, service)
	mux := http.NewServeMux()

	mux.HandleFunc("/v1/exchange", endpoint.EndpointExchange)

	srv := http.Server{
		Addr:              httpPort,
		Handler:           mux,
		ReadHeaderTimeout: time.Second * time.Duration(timeout),
	}
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
