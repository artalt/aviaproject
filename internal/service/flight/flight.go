package flight

import (
	"context"

	"github.com/google/uuid"

	flightDomain "homework/internal/domain/flight"
	flightStorage "homework/internal/storage/postgresql/flight"
)

type FlightService interface {
	GetFlightById(ctx context.Context, flightId uuid.UUID) (flightDomain.Flight, error)
	GetFlightList(ctx context.Context, filter flightDomain.Filter) ([]flightDomain.Flight, error)
}

type service struct {
	flightStorage flightStorage.FlightStorage
}

func (s *service) GetFlightList(ctx context.Context, filter flightDomain.Filter) ([]flightDomain.Flight, error) {
	return s.flightStorage.GetFlightList(ctx, filter)
}
func (s *service) GetFlightById(ctx context.Context, flightId uuid.UUID) (flightDomain.Flight, error) {
	return s.flightStorage.GetFlightById(ctx, flightId)
}

func NewFlightService(flightStorage flightStorage.FlightStorage) FlightService {
	return &service{
		flightStorage: flightStorage,
	}
}
