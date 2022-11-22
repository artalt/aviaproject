package flight

import (
	"context"

	"github.com/google/uuid"

	flightDomainPkg "homework/internal/domain/flight"
	flightStoragePkg "homework/internal/storage/postgresql/flight"
)

type FlightService interface {
	GetFlightById(ctx context.Context, flightId uuid.UUID) (flightDomainPkg.Flight, error)
	GetFlightList(ctx context.Context, filter flightDomainPkg.Filter) ([]flightDomainPkg.Flight, error)
}

type service struct {
	flightStorage flightStoragePkg.FlightStorage
}

func (s *service) GetFlightList(ctx context.Context, filter flightDomainPkg.Filter) ([]flightDomainPkg.Flight, error) {
	return s.flightStorage.GetFlightList(ctx, filter)
}
func (s *service) GetFlightById(ctx context.Context, flightId uuid.UUID) (flightDomainPkg.Flight, error) {
	return s.flightStorage.GetFlightById(ctx, flightId)
}

func NewFlightService(flightStorage flightStoragePkg.FlightStorage) FlightService {
	return &service{
		flightStorage: flightStorage,
	}
}
