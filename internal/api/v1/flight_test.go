package v1

import (
	"context"
	"encoding/json"
	"homework/specs"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	flightDomainPkg "homework/internal/domain/flight"
	mockFlightServicePkg "homework/internal/service/flight/mock"
)

func TestGetFlightById(t *testing.T) {
	// Arrange
	flightData := &specs.Flight{
		Id:                "c6832bfd-39c2-4520-b41c-ce9b1d530588",
		Number:            "LOL123",
		Arrival:           "she",
		Departure:         "dme",
		DepartureDateTime: "08:22:52 24-11-2022",
		ArrivalDateTime:   "08:22:54 25-11-2022",
		CanReturn:         true,
		TicketsCount:      35,
		HasLuggage:        true,
		Types:             "economy",
	}
	flightId, _ := uuid.Parse(flightData.Id)
	flight := &flightDomainPkg.Domain{
		Id:                flightId,
		Number:            flightData.Number,
		Arrival:           flightData.Arrival,
		Departure:         flightData.Departure,
		DepartureDateTime: flightData.DepartureDateTime,
		ArrivalDateTime:   flightData.ArrivalDateTime,
		CanReturn:         flightData.CanReturn,
		TicketsCount:      flightData.TicketsCount,
		MaxLuggage:        10,
		Types:             flightData.Types,
	}
	data, _ := json.Marshal(flightData)
	want := string(data)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	flightService := mockFlightServicePkg.NewMockFlightService(ctrl)
	flightService.EXPECT().
		GetFlightById(context.Background(), flightId).
		Return(flight, nil)

	flightController := NewAPIServer(flightService, nil)

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/flight/"+flightData.Id, nil)

	// Act
	flightController.GetFlightById(rec, req, flightData.Id)
	got := rec.Body.String()

	// Assert
	assert.JSONEq(t, want, got)
}
