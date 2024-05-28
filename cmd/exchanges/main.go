package main

import (
	"log"
	"net/http"
	"time"

	"github.com/crackc0der/exchanges/config"
	"github.com/crackc0der/exchanges/internal/exchange"
	"github.com/crackc0der/exchanges/logger"
)

func main() {
	config, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	addr := config.Host + config.Port
	logger := logger.NewLogger()
	service := exchange.NewService()
	endpoint := exchange.NewEndpoint(logger, service, config)
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/exchange", endpoint.EndpointExchange)
	srv := http.Server{
		Addr:                         addr,
		Handler:                      mux,
		ReadHeaderTimeout:            time.Second * time.Duration(config.Timeout),
		TLSConfig:                    nil,
		ReadTimeout:                  0,
		WriteTimeout:                 0,
		IdleTimeout:                  0,
		MaxHeaderBytes:               0,
		TLSNextProto:                 nil,
		ConnState:                    nil,
		ErrorLog:                     nil,
		BaseContext:                  nil,
		ConnContext:                  nil,
		DisableGeneralOptionsHandler: false,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
