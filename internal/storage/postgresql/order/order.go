package order

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"

	orderDomainPgk "homework/internal/domain/order"
	ticketDomainPgk "homework/internal/domain/ticket"
)

type OrderStorage interface {
	CreateOrder(ctx context.Context, order orderDomainPgk.Order, ticket ticketDomainPgk.Ticket) (orderDomainPgk.Order, error)
}

type storage struct {
	db *pgxpool.Pool
}

func (s *storage) CreateOrder(ctx context.Context, order orderDomainPgk.Order, ticket ticketDomainPgk.Ticket) (orderDomainPgk.Order, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	res, err := s.db.Exec(ctx, `
		INSERT INTO "order" VALUES ($1, $2, $3)`,
		id,
		order.GetStatus(),
		order.GetPaymentLink(),
	)
	if err != nil {
		return nil, err
	}
	if res.RowsAffected() != 1 {
		return nil, errors.New("insert order error")
	}

	res, err = s.db.Exec(ctx, `
		UPDATE ticket SET order_id = $1, pass_first_name = $2, pass_last_name = $3 , status = $4 WHERE id=$5`,
		id,
		ticket.GetPassFirstName(),
		ticket.GetPassLastName(),
		ticket.GetStatus(),
		ticket.GetId(),
	)
	if err != nil {
		return nil, err
	}
	if res.RowsAffected() != 1 {
		return nil, errors.New("update ticket error")
	}

	return &orderDomainPgk.Domain{
		Id:          id,
		Status:      order.GetStatus(),
		PaymentLink: order.GetPaymentLink(),
	}, nil
}

func NewOrderStorage(db *pgxpool.Pool) OrderStorage {
	return &storage{
		db: db,
	}
}
