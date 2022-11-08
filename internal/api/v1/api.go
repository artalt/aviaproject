package v1

import (
	"homework/specs"
	"net/http"
)

// Быстрая проверка актуальности текущего интерфейса сервера.
var _ specs.ServerInterface = &apiServer{}

type apiServer struct {
}

func (a apiServer) GetFlightList(w http.ResponseWriter, r *http.Request, params specs.GetFlightListParams) {
	//TODO implement me
	panic("implement me")
}

func (a apiServer) GetFlightById(w http.ResponseWriter, r *http.Request, id string) {
	//TODO implement me
	panic("implement me")
}

func (a apiServer) OrderTicket(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (a apiServer) Registration(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func NewAPIServer() specs.ServerInterface {
	return &apiServer{}
}
