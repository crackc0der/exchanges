package exchange

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type ExchangeServiceInterface interface {
	ServiceExchange(exc *Exchange) *ExchangeResult
}

func NewEndpoint(log *slog.Logger, service ExchangeServiceInterface) *Endpoint {
	return &Endpoint{log: log, service: service}
}

type Endpoint struct {
	service ExchangeServiceInterface
	log     *slog.Logger
}

func (e *Endpoint) EndpointExchange(writer http.ResponseWriter, request *http.Request) {
	var exc Exchange

	if err := json.NewDecoder(request.Body).Decode(&exc); err != nil {
		e.log.Error("error in method Endpoint.AddUser:" + err.Error())
	}

	result := e.service.ServiceExchange(&exc)

	if err := json.NewEncoder(writer).Encode(&result); err != nil {
		e.log.Error(err.Error())
	}
}
