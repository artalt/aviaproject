package v1

import (
	"fmt"
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

func (a apiServer) writeErrorResponse(w http.ResponseWriter, statusCode int, msg string) {
	w.WriteHeader(statusCode)
	_, err := w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, msg)))
	if err != nil {
		w.WriteHeader(500)
	}
}
