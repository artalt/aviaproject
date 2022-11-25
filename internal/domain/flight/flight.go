package flight

import "github.com/google/uuid"

type Flight interface {
	GetId() uuid.UUID
	GetNumber() string
	GetDeparture() string
	GetArrival() string
	GetDepartureDateTime() string
	GetArrivalDateTime() string
	GetCanReturn() bool
	GetTicketsCount() int
	GetMaxLuggage() int
	GetTypes() string
}

type Domain struct {
	Id                uuid.UUID
	Number            string
	Departure         string
	Arrival           string
	DepartureDateTime string
	ArrivalDateTime   string
	CanReturn         bool
	TicketsCount      int
	MaxLuggage        int
	Types             string
}

func (d *Domain) GetId() uuid.UUID {
	return d.Id
}
func (d *Domain) GetNumber() string {
	return d.Number
}
func (d *Domain) GetDeparture() string {
	return d.Departure
}
func (d *Domain) GetArrival() string {
	return d.Arrival
}
func (d *Domain) GetDepartureDateTime() string {
	return d.DepartureDateTime
}
func (d *Domain) GetArrivalDateTime() string {
	return d.ArrivalDateTime
}
func (d *Domain) GetCanReturn() bool {
	return d.CanReturn
}
func (d *Domain) GetTicketsCount() int {
	return d.TicketsCount
}
func (d *Domain) GetMaxLuggage() int {
	return d.MaxLuggage
}
func (d *Domain) GetTypes() string {
	return d.Types
}
