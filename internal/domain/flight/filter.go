package flight

import "time"

type Filter interface {
	GetDeparture() *string
	GetArrival() *string
	GetDateStart() *time.Time
	GetDateEnd() *time.Time
	GetType() *string
	GetHasLuggage() *bool
}

type FilterDto struct {
	Departure  *string
	Arrival    *string
	DateStart  *time.Time
	DateEnd    *time.Time
	Type       *string
	HasLuggage *bool
}

func (f *FilterDto) GetDeparture() *string {
	return f.Departure
}
func (f *FilterDto) GetArrival() *string {
	return f.Arrival
}
func (f *FilterDto) GetDateStart() *time.Time {
	return f.DateStart
}
func (f *FilterDto) GetDateEnd() *time.Time {
	return f.DateEnd
}
func (f *FilterDto) GetType() *string {
	return f.Type
}
func (f *FilterDto) GetHasLuggage() *bool {
	return f.HasLuggage
}
func (f *FilterDto) IsValid() bool {
	if nil != f.Departure && "" == *f.Departure {
		return false
	}
	if nil != f.Arrival && "" == *f.Arrival {
		return false
	}
	if nil != f.Type && "" == *f.Type {
		return false
	}

	return true
}
