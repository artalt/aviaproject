package flight

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	flightDomainPkg "homework/internal/domain/flight"
	mockFlightStoragePkg "homework/internal/storage/postgresql/flight/mock"
)

func TestGetFlightById(t *testing.T) {
	// Arrange
	flightId, _ := uuid.Parse("c6832bfd-39c2-4520-b41c-ce9b1d530588")
	flight := &flightDomainPkg.Domain{
		Id:                flightId,
		Number:            "LOL123",
		Arrival:           "she",
		Departure:         "dme",
		DepartureDateTime: "08:22:52 24-11-2022",
		ArrivalDateTime:   "08:22:54 25-11-2022",
		CanReturn:         true,
		TicketsCount:      35,
		MaxLuggage:        10,
		Types:             "economy",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	flightStorage := mockFlightStoragePkg.NewMockFlightStorage(ctrl)
	flightStorage.EXPECT().
		GetFlightById(context.Background(), flightId).
		Return(flight, nil)

	flightService := NewFlightService(flightStorage)

	// Act
	got, err := flightService.GetFlightById(context.Background(), flightId)

	// Assert
	require.Equal(t, nil, err)
	assert.Equal(t, flight, got)
}
