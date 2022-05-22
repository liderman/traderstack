package grpcsrv

import (
	"github.com/liderman/traderstack/internal/engine"
	"github.com/liderman/traderstack/internal/engine/baseoption"
	stackv1 "github.com/liderman/traderstack/internal/grpcsrv/gen/go/liderman/traderstack/stack/v1"
	"github.com/shopspring/decimal"
)

type ToDomainMapper struct {
}

func (t *ToDomainMapper) MapSetItems(in []*stackv1.SetItem) []*engine.SetStackItem {
	var ret []*engine.SetStackItem
	for _, v := range in {
		ret = append(ret, t.MapSetItem(v))
	}

	return ret
}

func (t *ToDomainMapper) MapSetItem(in *stackv1.SetItem) *engine.SetStackItem {
	return &engine.SetStackItem{
		Variable:  in.Variable,
		StackFunc: t.MapSetStackFunc(in.StackFunc),
	}
}

func (t *ToDomainMapper) MapSetStackFunc(in *stackv1.SetStackFunc) *engine.SetStackFunc {
	if in == nil {
		return nil
	}

	fnc := &engine.SetStackFunc{
		Name: in.Name,
	}

	for _, v := range in.Arguments {
		fnc.Arguments = append(fnc.Arguments, t.MapSetArgument(v))
	}

	return fnc
}

func (t *ToDomainMapper) MapSetArgument(in *stackv1.SetArgument) *engine.SetArgument {
	ret := &engine.SetArgument{
		Id: in.Id,
	}

	switch v := in.Value.(type) {
	case *stackv1.SetArgument_Variable:
		ret.Value = t.MapVariable(v.Variable)
	case *stackv1.SetArgument_Input:
		ret.Value = t.MapValue(v.Input)
	}

	return ret
}

func (t *ToDomainMapper) MapValue(in *stackv1.Value) interface{} {
	if in == nil {
		return nil
	}

	switch v := in.Val.(type) {
	case *stackv1.Value_Integer:
		return v.Integer
	case *stackv1.Value_String_:
		return v.String_
	case *stackv1.Value_Boolean:
		return v.Boolean
	case *stackv1.Value_Decimal:
		return t.MapDecimal(v.Decimal)
	case *stackv1.Value_Time:
		return v.Time.AsTime()
	}
	return nil
}

func (t *ToDomainMapper) MapDecimal(in string) *decimal.Decimal {
	if in == "" {
		return &decimal.Zero
	}

	ret, err := decimal.NewFromString(in)
	if err != nil {
		return &decimal.Zero
	}

	return &ret
}

func (t *ToDomainMapper) MapVariable(in *stackv1.Variable) *baseoption.Variable {
	if in == nil {
		return nil
	}

	return &baseoption.Variable{
		Name: in.Name,
	}
}
