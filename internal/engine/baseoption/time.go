package baseoption

import "time"

type Time struct {
	Id           string    `json:"id"`
	Name         string    `json:"name"`
	Desc         string    `json:"desc"`
	BaseType     string    `json:"baseType"`
	ExtendedType string    `json:"extendedType"`
	Default      time.Time `json:"default"`
	Required     bool      `json:"required"`
	Min          *int64    `json:"min,omitempty"`
	Max          *int64    `json:"max,omitempty"`
	Step         *int64    `json:"step,omitempty"`
}

func NewTime(id string) *Time {
	return &Time{
		Id:       id,
		BaseType: "time",
	}
}
