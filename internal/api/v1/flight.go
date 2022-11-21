package v1

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"

	flightDomain "homework/internal/domain/flight"
	"homework/specs"
)

func (a apiServer) GetFlightList(w http.ResponseWriter, r *http.Request, params specs.GetFlightListParams) {
	var dateStartParam *time.Time
	if nil != params.DateStart {
		t1, err := time.Parse("02-01-2006", *params.DateStart)
		if err != nil {
			w.WriteHeader(400)
			return
		}
		dateStartParam = &t1
	}
	var dateEndParam *time.Time
	if nil != params.DateEnd {
		t2, err := time.Parse("02-01-2006", *params.DateEnd)
		if err != nil {
			w.WriteHeader(400)
			return
		}
		dateEndParam = &t2
	}
	filter := &flightDomain.FilterDto{
		Departure: params.Departure,
		Arrival:   params.Arrival,
		DateStart: dateStartParam,
		DateEnd:   dateEndParam,
		Type:      params.Type,
	}
	if nil != params.HasLuggage {
		hasLuggageFlag := "1" == *params.HasLuggage
		filter.HasLuggage = &hasLuggageFlag
	}
	if !filter.IsValid() {
		w.WriteHeader(400)
		return
	}
	ctx := context.Background()
	list, err := a.flightService.GetFlightList(ctx, filter)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	var results []*specs.Flight
	for _, flight := range list {
		flightData := &specs.Flight{
			Id:                flight.GetId().String(),
			Number:            flight.GetNumber(),
			Arrival:           flight.GetArrival(),
			Departure:         flight.GetDeparture(),
			DepartureDateTime: flight.GetDepartureDateTime(),
			ArrivalDateTime:   flight.GetArrivalDateTime(),
			CanReturn:         flight.GetCanReturn(),
			TicketsCount:      flight.GetTicketsCount(),
			HasLuggage:        flight.GetMaxLuggage() > 0,
			Types:             flight.GetTypes(),
		}

		results = append(results, flightData)
	}

	response, err := json.Marshal(results)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	_, err = w.Write(response)
	if err != nil {
		panic("can't write response")
	}
}

func (a apiServer) GetFlightById(w http.ResponseWriter, r *http.Request, id string) {
	flightId, err := uuid.Parse(id)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	ctx := context.Background()
	flight, err := a.flightService.GetFlightById(ctx, flightId)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	flightData := &specs.Flight{
		Id:                flight.GetId().String(),
		Number:            flight.GetNumber(),
		Arrival:           flight.GetArrival(),
		Departure:         flight.GetDeparture(),
		DepartureDateTime: flight.GetDepartureDateTime(),
		ArrivalDateTime:   flight.GetArrivalDateTime(),
		CanReturn:         flight.GetCanReturn(),
		TicketsCount:      flight.GetTicketsCount(),
		HasLuggage:        flight.GetMaxLuggage() > 0,
		Types:             flight.GetTypes(),
	}

	response, err := json.Marshal(flightData)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	_, err = w.Write(response)
	if err != nil {
		panic("can't write response")
	}
}
