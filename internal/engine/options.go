package engine

import (
	"errors"
	"github.com/liderman/traderstack/internal/engine/baseoption"
	"github.com/shopspring/decimal"
)

var ErrOptionArgNotExist = errors.New("argument is not exist")
var ErrOptionVarNotExist = errors.New("variable is not exist")
var ErrOptionBadType = errors.New("value is bad type")

type Options struct {
	argIdx map[string]interface{}
	varIdx map[string]interface{}
}

func NewOptions(arguments []*Argument, varIdx map[string]interface{}) *Options {
	argIdx := map[string]interface{}{}
	for _, v := range arguments {
		argIdx[v.Id] = v.Value
	}

	return &Options{
		argIdx: argIdx,
		varIdx: varIdx,
	}
}

func (o *Options) get(key string) (interface{}, error) {
	v := o.argIdx[key]
	if v == nil {
		return v, ErrOptionArgNotExist
	}

	variable, varOk := v.(*baseoption.Variable)
	if varOk {
		ret, ok := o.varIdx[variable.Name]
		if !ok {
			return ret, ErrOptionVarNotExist
		}
		return ret, nil
	}

	return v, nil
}

func (o *Options) GetInteger(key string) (int64, error) {
	v, err := o.get(key)
	if err != nil {
		return 0, err
	}

	ret, ok := v.(int64)
	if !ok {
		return 0, ErrOptionBadType
	}
	return ret, nil
}

func (o *Options) GetString(key string) (string, error) {
	v, err := o.get(key)
	if err != nil {
		return "", err
	}

	ret, ok := v.(string)
	if !ok {
		return "", ErrOptionBadType
	}
	return ret, nil
}

func (o *Options) GetBoolean(key string) (bool, error) {
	v, err := o.get(key)
	if err != nil {
		return false, err
	}

	ret, ok := v.(bool)
	if !ok {
		return false, ErrOptionBadType
	}
	return ret, nil
}

func (o *Options) GetDecimal(key string) (decimal.Decimal, error) {
	v, err := o.get(key)
	if err != nil {
		return decimal.Zero, err
	}

	ret, ok := v.(decimal.Decimal)
	if !ok {
		return decimal.Zero, ErrOptionBadType
	}
	return ret, nil
}

func (o *Options) GetNumericDecimal(key string) (decimal.Decimal, error) {
	v, err := o.get(key)
	if err != nil {
		return decimal.Zero, err
	}

	switch val := v.(type) {
	case decimal.Decimal:
		return val, nil
	case int64:
		return decimal.NewFromInt(val), nil
	}

	return decimal.Zero, ErrOptionBadType
}
