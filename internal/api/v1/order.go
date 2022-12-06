package v1

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"homework/specs"
	"net/http"
)

func (a apiServer) OrderTicket(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)
	var request specs.CreateOrderRequest
	err := decoder.Decode(&request)
	if err != nil || "" == request.FlightId || "" == request.FirstName || "" == request.LastName {
		a.writeErrorResponse(w, 400, "invalid request params")
		return
	}
	flightId, err := uuid.Parse(request.FlightId)
	if err != nil {
		a.writeErrorResponse(w, 400, "invalid param `flightId`")
		return
	}

	ctx := context.Background()
	order, ticket, err := a.orderService.OrderTicket(ctx, flightId, request.FirstName, request.LastName)
	if err != nil {
		a.writeErrorResponse(w, 500, err.Error())
		return
	}

	orderData := &specs.Order{
		Id:          order.GetId().String(),
		Status:      order.GetStatus(),
		PaymentLink: order.GetPaymentLink(),
		Ticket: specs.Ticket{
			Id:            ticket.GetId().String(),
			FlightId:      ticket.GetFlightId().String(),
			Status:        ticket.GetStatus(),
			Type:          ticket.GetType(),
			PassFirstName: ticket.GetPassFirstName(),
			PassLastName:  ticket.GetPassLastName(),
			Seat:          ticket.GetSeat(),
			Luggage:       ticket.GetLuggage(),
			Price:         ticket.GetPrice(),
			CanReturn:     ticket.GetCanReturn(),
		},
	}

	response, err := json.Marshal(orderData)
	if err != nil {
		a.writeErrorResponse(w, 500, "Error create json response")
		return
	}

	_, err = w.Write(response)
	if err != nil {
		a.writeErrorResponse(w, 500, "Can't write response")
	}
}
