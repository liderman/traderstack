package variable

import "github.com/shopspring/decimal"

type StringValidator struct {
	Regexp string
	MinLen int32
	MaxLen int32
}

type IntegerValidator struct {
	Min  int64
	Max  int64
	Step int64
}

type DecimalValidator struct {
	Min  *decimal.Decimal
	Max  *decimal.Decimal
	Step *decimal.Decimal
}

type TimeValidator struct {
}
