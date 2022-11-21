package v1

import (
	"net/http"

	"homework/internal/service/flight"
	"homework/specs"
)

// Быстрая проверка актуальности текущего интерфейса сервера.
var _ specs.ServerInterface = &apiServer{}

type apiServer struct {
	flightService flight.FlightService
}

func (a apiServer) OrderTicket(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (a apiServer) Registration(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func NewAPIServer(flightService flight.FlightService) specs.ServerInterface {
	return &apiServer{
		flightService: flightService,
	}
}
