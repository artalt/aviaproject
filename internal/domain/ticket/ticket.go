package ticket

import "github.com/google/uuid"

type Ticket interface {
	GetId() uuid.UUID
	GetFlightId() uuid.UUID
	GetOrderId() uuid.UUID
	GetStatus() string
	GetType() string
	GetPassFirstName() *string
	GetPassLastName() *string
	GetSeat() string
	GetLuggage() int
	GetPrice() int
	GetCanReturn() bool
	SetStatus(status string)
	SetPassFirstName(firstName string)
	SetPassLastName(lastName string)
}

type Domain struct {
	Id            uuid.UUID
	FlightId      uuid.UUID
	OrderId       uuid.UUID
	Status        string
	Type          string
	PassFirstName *string
	PassLastName  *string
	Seat          string
	Luggage       int
	Price         int
	CanReturn     bool
}

func (d *Domain) GetId() uuid.UUID {
	return d.Id
}
func (d *Domain) GetFlightId() uuid.UUID {
	return d.FlightId
}
func (d *Domain) GetOrderId() uuid.UUID {
	return d.OrderId
}
func (d *Domain) GetStatus() string {
	return d.Status
}
func (d *Domain) GetType() string {
	return d.Type
}
func (d *Domain) GetPassFirstName() *string {
	return d.PassFirstName
}
func (d *Domain) GetPassLastName() *string {
	return d.PassLastName
}
func (d *Domain) GetSeat() string {
	return d.Seat
}
func (d *Domain) GetLuggage() int {
	return d.Luggage
}
func (d *Domain) GetPrice() int {
	return d.Price
}
func (d *Domain) GetCanReturn() bool {
	return d.CanReturn
}
func (d *Domain) SetStatus(status string) {
	d.Status = status
}
func (d *Domain) SetPassFirstName(firstName string) {
	d.PassFirstName = &firstName
}
func (d *Domain) SetPassLastName(lastName string) {
	d.PassLastName = &lastName
}
