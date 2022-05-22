package baseoption

import "github.com/shopspring/decimal"

type Decimal struct {
	Id           string           `json:"id"`
	Name         string           `json:"name"`
	Desc         string           `json:"desc"`
	BaseType     string           `json:"baseType"`
	ExtendedType string           `json:"extendedType"`
	Default      decimal.Decimal  `json:"default"`
	Required     bool             `json:"required"`
	Min          *decimal.Decimal `json:"min,omitempty"`
	Max          *decimal.Decimal `json:"max,omitempty"`
	Step         *decimal.Decimal `json:"step,omitempty"`
}

func NewDecimal(id string) *Decimal {
	return &Decimal{
		Id:       id,
		BaseType: "decimal",
	}
}
