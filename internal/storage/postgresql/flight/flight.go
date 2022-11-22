package flight

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"

	flightDomainPkg "homework/internal/domain/flight"
)

type FlightStorage interface {
	GetFlightById(ctx context.Context, flightId uuid.UUID) (flightDomainPkg.Flight, error)
	GetFlightList(ctx context.Context, filter flightDomainPkg.Filter) ([]flightDomainPkg.Flight, error)
}

type storage struct {
	db *pgxpool.Pool
}

func (s *storage) GetFlightList(ctx context.Context, filter flightDomainPkg.Filter) ([]flightDomainPkg.Flight, error) {
	var where []string
	if nil != filter.GetDeparture() {
		where = append(where, fmt.Sprintf("f.departure = '%s'", *filter.GetDeparture()))
	}
	if nil != filter.GetArrival() {
		where = append(where, fmt.Sprintf("f.arrival = '%s'", *filter.GetArrival()))
	}
	if nil != filter.GetDateStart() {
		where = append(where, fmt.Sprintf(
			"f.departure_date_time > '%s 00:00:00'",
			(*filter.GetDateStart()).Format("2006-01-02"),
		))
	}
	if nil != filter.GetDateEnd() {
		where = append(where, fmt.Sprintf(
			"f.departure_date_time < '%s 23:59:59'",
			(*filter.GetDateEnd()).Format("2006-01-02"),
		))
	}
	if nil != filter.GetType() {
		where = append(where, fmt.Sprintf("t.type = '%s'", *filter.GetType()))
	}
	if nil != filter.GetHasLuggage() {
		if *filter.GetHasLuggage() {
			where = append(where, "t.luggage > 0")
		} else {
			where = append(where, "t.luggage = 0")
		}
	}
	var sqlWhere string
	if len(where) > 0 {
		sqlWhere = " WHERE " + strings.Join(where, " AND ")
	}

	sql := fmt.Sprintf(`
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
		%s GROUP BY
			f.id,
			f.number,
			f.arrival,
			f.departure,
			f.departure_date_time,
			f.arrival_date_time;`,
		sqlWhere,
	)

	rows, err := s.db.Query(ctx, sql)
	if err != nil {
		log.Println(err)
		return nil, errors.New("can't get flight list from db")
	}
	defer rows.Close()

	var list []flightDomainPkg.Flight
	for rows.Next() {
		var f flightDomainPkg.Domain
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
			return nil, errors.New("error scan results from db")
		}

		list = append(list, &f)
	}

	return list, err
}

func (s *storage) GetFlightById(ctx context.Context, flightId uuid.UUID) (flightDomainPkg.Flight, error) {
	row := s.db.QueryRow(ctx, `
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

	var f flightDomainPkg.Domain
	err := row.Scan(
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
		return nil, errors.New("error scan result from db")
	}

	return &f, err
}

func NewFlightStorage(db *pgxpool.Pool) FlightStorage {
	return &storage{
		db: db,
	}
}
