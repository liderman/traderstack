package grpcsrv

import (
	"github.com/liderman/traderstack/internal/engine"
	"github.com/liderman/traderstack/internal/engine/baseoption"
	stackv1 "github.com/liderman/traderstack/internal/grpcsrv/gen/go/liderman/traderstack/stack/v1"
	"github.com/shopspring/decimal"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type ToProtobufMapper struct {
}

func (t *ToProtobufMapper) MapStacks(in []*engine.Stack) []*stackv1.Stack {
	ret := make([]*stackv1.Stack, 0, len(in))
	for _, v := range in {
		ret = append(ret, t.MapStack(v))
	}

	return ret
}

func (t *ToProtobufMapper) MapStack(in *engine.Stack) *stackv1.Stack {
	if in == nil {
		return nil
	}
	ret := &stackv1.Stack{
		Id:    in.Id,
		Name:  in.Name,
		Items: []*stackv1.Item{},
	}

	for _, v := range in.Items {
		ret.Items = append(ret.Items, t.MapItem(v))
	}

	return ret
}

func (t *ToProtobufMapper) MapItem(in *engine.StackItem) *stackv1.Item {
	return &stackv1.Item{
		Variable:  in.Variable,
		StackFunc: t.MapStackFunc(in.StackFunc),
	}
}

func (t *ToProtobufMapper) MapStackFunc(in *engine.StackFunc) *stackv1.StackFunc {
	if in == nil {
		return nil
	}
	return &stackv1.StackFunc{
		Name:      in.Name,
		Arguments: t.MapArguments(in.Arguments),
		BaseType:  in.BaseType,
	}
}

func (t *ToProtobufMapper) MapArguments(in []*engine.Argument) []*stackv1.Argument {
	ret := make([]*stackv1.Argument, 0, len(in))
	for _, v := range in {
		ret = append(ret, t.MapArgument(v))
	}
	return ret
}

func (t *ToProtobufMapper) MapArgument(in *engine.Argument) *stackv1.Argument {
	if in == nil {
		return nil
	}

	ret := &stackv1.Argument{
		Id:           in.Id,
		Name:         in.Name,
		Desc:         in.Desc,
		BaseType:     in.BaseType,
		ExtendedType: in.ExtendedType,
		Required:     in.Required,
	}

	switch v := in.Value.(type) {
	case *baseoption.Variable:
		ret.Value = &stackv1.Argument_Variable{
			Variable: t.MapVariable(v),
		}
	default:
		ret.Value = &stackv1.Argument_Input{
			Input: t.MapValue(v),
		}
	}

	return ret
}

func (t *ToProtobufMapper) MapValue(in interface{}) *stackv1.Value {
	if in == nil {
		return nil
	}

	ret := &stackv1.Value{}

	switch v := in.(type) {
	case int64:
		ret.Val = &stackv1.Value_Integer{
			Integer: v,
		}
	case string:
		ret.Val = &stackv1.Value_String_{
			String_: v,
		}
	case bool:
		ret.Val = &stackv1.Value_Boolean{
			Boolean: v,
		}
	case *decimal.Decimal:
		ret.Val = &stackv1.Value_Decimal{
			Decimal: t.MapDecimal(v),
		}
	case time.Time:
		ret.Val = &stackv1.Value_Time{
			Time: timestamppb.New(v),
		}
	default:
		return nil
	}

	return ret
}

func (t *ToProtobufMapper) MapDecimal(in *decimal.Decimal) string {
	return in.String()
}

func (t *ToProtobufMapper) MapVariable(in *baseoption.Variable) *stackv1.Variable {
	return &stackv1.Variable{
		Name: in.Name,
	}
}

func (t *ToProtobufMapper) MapFuncs(in []*engine.StackFunc) []*stackv1.StackFunc {
	ret := make([]*stackv1.StackFunc, 0, len(in))

	for _, v := range in {
		ret = append(ret, t.MapStackFunc(v))
	}

	return ret
}

func (t *ToProtobufMapper) MapTestItemResults(in []*engine.TestItemResult) []*stackv1.TestItemResult {
	ret := make([]*stackv1.TestItemResult, 0, len(in))

	for _, v := range in {
		ret = append(ret, t.MapTestItemResult(v))
	}

	return ret
}

func (t *ToProtobufMapper) MapTestItemResult(in *engine.TestItemResult) *stackv1.TestItemResult {
	return &stackv1.TestItemResult{
		Variable: in.Variable,
		Result:   t.MapValue(in.Result),
		BaseType: in.BaseType,
		Error:    in.Error,
	}
}
