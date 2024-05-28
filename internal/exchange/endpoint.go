package exchange

import (
	"encoding/json"
	"net/http"

	"github.com/crackc0der/exchanges/config"
	"github.com/crackc0der/exchanges/logger"
)

type ExchangeServiceInterface interface {
	ServiceExchange(exc *Exchange) *ExchangeResult
}

func NewEndpoint(log *logger.Logger, service ExchangeServiceInterface, config *config.Config) *Endpoint {
	return &Endpoint{logger: log, service: service, config: config}
}

type Endpoint struct {
	service ExchangeServiceInterface
	logger  *logger.Logger
	config  *config.Config
}

func (e *Endpoint) EndpointExchange(writer http.ResponseWriter, request *http.Request) {
	var exc Exchange

	if err := json.NewDecoder(request.Body).Decode(&exc); err != nil {
		e.logger.Log(e.config.LogLevel, "error in method Endpoint.AddUser:"+err.Error())
	}

	result := e.service.ServiceExchange(&exc)

	if err := json.NewEncoder(writer).Encode(&result); err != nil {
		e.logger.Log(e.config.LogLevel, "error in method Endpoint.AddUser:"+err.Error())
	}
}
