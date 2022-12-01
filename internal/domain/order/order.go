package order

import "github.com/google/uuid"

type Order interface {
	GetId() uuid.UUID
	GetStatus() string
	GetPaymentLink() string
}

type Domain struct {
	Id          uuid.UUID
	Status      string
	PaymentLink string
}

func (d *Domain) GetId() uuid.UUID {
	return d.Id
}
func (d *Domain) GetStatus() string {
	return d.Status
}
func (d *Domain) GetPaymentLink() string {
	return d.PaymentLink
}
