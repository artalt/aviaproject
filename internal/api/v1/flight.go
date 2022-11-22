package v1

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"

	flightDomainPkg "homework/internal/domain/flight"
	"homework/specs"
)

func (a apiServer) GetFlightList(w http.ResponseWriter, r *http.Request, params specs.GetFlightListParams) {
	w.Header().Add("Content-Type", "application/json")
	var dateStartParam *time.Time
	if nil != params.DateStart {
		t1, err := time.Parse("02-01-2006", *params.DateStart)
		if err != nil {
			a.writeErrorResponse(w, 400, "Invalid param dateStart")
			return
		}
		dateStartParam = &t1
	}
	var dateEndParam *time.Time
	if nil != params.DateEnd {
		t2, err := time.Parse("02-01-2006", *params.DateEnd)
		if err != nil {
			a.writeErrorResponse(w, 400, "Invalid param dateEnd")
			return
		}
		dateEndParam = &t2
	}
	filter := &flightDomainPkg.FilterDto{
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
		a.writeErrorResponse(w, 400, "Invalid params")
		return
	}
	ctx := context.Background()
	list, err := a.flightService.GetFlightList(ctx, filter)
	if err != nil {
		a.writeErrorResponse(w, 500, err.Error())
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
		a.writeErrorResponse(w, 500, "Error create json response")
		return
	}

	_, err = w.Write(response)
	if err != nil {
		a.writeErrorResponse(w, 500, "Can't write response")
	}
}

func (a apiServer) GetFlightById(w http.ResponseWriter, r *http.Request, id string) {
	w.Header().Add("Content-Type", "application/json")
	flightId, err := uuid.Parse(id)
	if err != nil {
		a.writeErrorResponse(w, 400, "Invalid path param id")
		return
	}
	ctx := context.Background()
	flight, err := a.flightService.GetFlightById(ctx, flightId)
	if err != nil {
		a.writeErrorResponse(w, 500, err.Error())
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
		a.writeErrorResponse(w, 500, "Error create json response")
		return
	}

	_, err = w.Write(response)
	if err != nil {
		a.writeErrorResponse(w, 500, "Can't write response")
	}
}
