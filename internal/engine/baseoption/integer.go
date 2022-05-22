package baseoption

import (
	"errors"
	"fmt"
)

type Integer struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Desc         string `json:"desc"`
	BaseType     string `json:"baseType"`
	ExtendedType string `json:"extendedType"`
	Default      int64  `json:"default"`
	Required     bool   `json:"required"`
	Min          *int64 `json:"min,omitempty"`
	Max          *int64 `json:"max,omitempty"`
	Step         *int64 `json:"step,omitempty"`
}

func NewInteger(id string) *Integer {
	return &Integer{
		Id:       id,
		BaseType: "integer",
	}
}

func (s *Integer) Validate(value any) error {
	v, ok := value.(int64)
	if !ok {
		return errors.New("bad value type")
	}

	if s.Min != nil && v < *s.Min {
		return fmt.Errorf("minimum value %d", *s.Min)
	}

	if s.Max != nil && v > *s.Max {
		return fmt.Errorf("maximum value %d", *s.Max)
	}

	if s.Step != nil && v%*s.Step != 0 {
		return fmt.Errorf("value is not valid step %d", *s.Step)
	}

	return nil
}
