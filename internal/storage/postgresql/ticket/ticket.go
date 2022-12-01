package ticket

import (
	"context"
	"errors"
	"log"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"

	ticketDomainPkg "homework/internal/domain/ticket"
)

type TicketStorage interface {
	GetTicketByFlightId(ctx context.Context, flightId uuid.UUID) (ticketDomainPkg.Ticket, error)
}

type storage struct {
	db *pgxpool.Pool
}

func (s *storage) GetTicketByFlightId(ctx context.Context, flightId uuid.UUID) (ticketDomainPkg.Ticket, error) {
	row := s.db.QueryRow(ctx, `
		SELECT
			id,
			flight_id,
			order_id,
			status,
			type,
			pass_first_name,
			pass_last_name,
			seat,
			luggage,
			price,
			can_return
		FROM ticket
		WHERE 
			flight_id = $1
			AND status = 'free'
		`,
		flightId,
	)

	var t ticketDomainPkg.Domain
	err := row.Scan(
		&t.Id,
		&t.FlightId,
		&t.OrderId,
		&t.Status,
		&t.Type,
		&t.PassFirstName,
		&t.PassLastName,
		&t.Seat,
		&t.Luggage,
		&t.Price,
		&t.CanReturn,
	)
	if err != nil {
		log.Println(err)
		return nil, errors.New("free ticket not found")
	}

	return &t, err
}

func NewTicketStorage(db *pgxpool.Pool) TicketStorage {
	return &storage{
		db: db,
	}
}
