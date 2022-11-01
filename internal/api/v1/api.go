package v1

import (
	"homework/specs"
)

// Быстрая проверка актуальности текущего интерфейса сервера.
var _ specs.ServerInterface = &apiServer{}

type apiServer struct {
}

func NewAPIServer() specs.ServerInterface {
	return &apiServer{}
}
