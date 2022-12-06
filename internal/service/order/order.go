package order

import (
	"context"

	"github.com/google/uuid"

	orderDomainPgk "homework/internal/domain/order"
	ticketDomainPgk "homework/internal/domain/ticket"
	orderStoragePkg "homework/internal/storage/postgresql/order"
	ticketStoragePkg "homework/internal/storage/postgresql/ticket"
)

type OrderService interface {
	OrderTicket(ctx context.Context, flightId uuid.UUID, firstName string, lastName string) (orderDomainPgk.Order, ticketDomainPgk.Ticket, error)
}

type service struct {
	orderStorage  orderStoragePkg.OrderStorage
	ticketStorage ticketStoragePkg.TicketStorage
}

func (s *service) OrderTicket(ctx context.Context, flightId uuid.UUID, firstName string, lastName string) (orderDomainPgk.Order, ticketDomainPgk.Ticket, error) {
	ticket, err := s.ticketStorage.GetTicketByFlightId(ctx, flightId)
	if err != nil {
		return nil, nil, err
	}

	var order orderDomainPgk.Order
	order = &orderDomainPgk.Domain{
		Status:      "booked",
		PaymentLink: "https://please_pay.me?transactionId=123321",
	}
	ticket.SetStatus("taken")
	ticket.SetPassFirstName(firstName)
	ticket.SetPassLastName(lastName)
	order, err = s.orderStorage.CreateOrder(ctx, order, ticket)

	return order, ticket, err
}

func NewOrderService(orderStorage orderStoragePkg.OrderStorage, ticketStorage ticketStoragePkg.TicketStorage) OrderService {
	return &service{
		orderStorage:  orderStorage,
		ticketStorage: ticketStorage,
	}
}
