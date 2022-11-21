package flight

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"

	flightDomain "homework/internal/domain/flight"
)

type FlightStorage interface {
	GetFlightById(ctx context.Context, flightId uuid.UUID) (flightDomain.Flight, error)
	GetFlightList(ctx context.Context, filter flightDomain.Filter) ([]flightDomain.Flight, error)
}

type storage struct {
	db *pgxpool.Pool
}

func (s *storage) GetFlightList(ctx context.Context, filter flightDomain.Filter) ([]flightDomain.Flight, error) {
	sql := `
		SELECT
			f.id,
			f.number,
			f.arrival,
			f.departure,
			TO_CHAR(f.departure_date_time, 'hh:ii:ss dd-mm-yyyy'),
			TO_CHAR(f.arrival_date_time, 'hh:ii:ss dd-mm-yyyy'),
			BOOL_OR(t.can_return) AS can_return,
			COUNT(t.id) AS tickets_count,
			MAX(t.luggage) AS max_luggage,
			STRING_AGG(DISTINCT t.type, ', ') AS types
		FROM flight f
			LEFT JOIN ticket t on f.id = t.flight_id AND t.status = 'free'
		WHERE 1=1`

	if nil != filter.GetDeparture() {
		sql += fmt.Sprintf(" AND f.departure = '%s'", *filter.GetDeparture())
	}
	if nil != filter.GetArrival() {
		sql += fmt.Sprintf(" AND f.arrival = '%s'", *filter.GetArrival())
	}
	if nil != filter.GetDateStart() {
		sql += fmt.Sprintf(
			" AND f.departure_date_time > '%s 00:00:00'",
			(*filter.GetDateStart()).Format("2006-01-02"),
		)
	}
	if nil != filter.GetDateEnd() {
		sql += fmt.Sprintf(
			" AND f.departure_date_time < '%s 23:59:59'",
			(*filter.GetDateEnd()).Format("2006-01-02"),
		)
	}
	if nil != filter.GetType() {
		sql += fmt.Sprintf(" AND t.type = '%s'", *filter.GetType())
	}
	if nil != filter.GetHasLuggage() {
		if *filter.GetHasLuggage() {
			sql += " AND t.luggage > 0"
		} else {
			sql += " AND t.luggage = 0"
		}
	}

	sql += ` GROUP BY
			f.id,
			f.number,
			f.arrival,
			f.departure,
			f.departure_date_time,
			f.arrival_date_time;`

	rows, err := s.db.Query(ctx, sql)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var list []flightDomain.Flight
	for rows.Next() {
		var f flightDomain.Domain
		err := rows.Scan(
			&f.Id,
			&f.Number,
			&f.Arrival,
			&f.Departure,
			&f.DepartureDateTime,
			&f.ArrivalDateTime,
			&f.CanReturn,
			&f.TicketsCount,
			&f.MaxLuggage,
			&f.Types,
		)
		if err != nil {
			log.Println(err)
		}

		list = append(list, &f)
	}

	return list, err
}

func (s *storage) GetFlightById(ctx context.Context, flightId uuid.UUID) (flightDomain.Flight, error) {
	rows, err := s.db.Query(ctx, `
		SELECT
			f.id,
			f.number,
			f.arrival,
			f.departure,
			TO_CHAR(f.departure_date_time, 'hh:ii:ss dd-mm-yyyy'),
			TO_CHAR(f.arrival_date_time, 'hh:ii:ss dd-mm-yyyy'),
			BOOL_OR(t.can_return) AS can_return,
			COUNT(t.id) AS tickets_count,
			MAX(t.luggage) AS max_luggage,
			STRING_AGG(DISTINCT t.type, ', ') AS types
		FROM flight f
			LEFT JOIN ticket t on f.id = t.flight_id AND t.status = 'free'
		WHERE f.id = $1
		GROUP BY
			f.id,
			f.number,
			f.arrival,
			f.departure,
			f.departure_date_time,
			f.arrival_date_time;`,
		flightId,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var f flightDomain.Domain
	rows.Next()
	err = rows.Scan(
		&f.Id,
		&f.Number,
		&f.Arrival,
		&f.Departure,
		&f.DepartureDateTime,
		&f.ArrivalDateTime,
		&f.CanReturn,
		&f.TicketsCount,
		&f.MaxLuggage,
		&f.Types,
	)

	return &f, err
}

func NewFlightStorage(db *pgxpool.Pool) FlightStorage {
	return &storage{
		db: db,
	}
}
