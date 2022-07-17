package engine

import (
	"fmt"
	"time"
)

// TestStack Движок тестирования стека.
type TestStack struct {
	sm  *StackManager
	sfr *StackFuncRepository
}

func NewTestStack(sm *StackManager, sfr *StackFuncRepository) *TestStack {
	return &TestStack{
		sm:  sm,
		sfr: sfr,
	}
}

type TestItemResult struct {
	Variable string
	Result   interface{}
	BaseType string
	Error    string
}

func (t *TestStack) Test(stackId string, now time.Time, accountId string) ([]*TestItemResult, error) {
	stack := t.sm.Get(stackId)
	if stack == nil {
		return nil, fmt.Errorf("stack `%s` is not exist", stackId)
	}

	varIdx := map[string]interface{}{}
	ret := make([]*TestItemResult, 0, len(stack.Items))
	for _, v := range stack.Items {
		res := t.testItem(v, varIdx, now, accountId)
		if res.Error == "" {
			varIdx[res.Variable] = res.Result
		}

		ret = append(ret, res)
	}

	return ret, nil
}

func (t *TestStack) testItem(item *StackItem, varIdx map[string]interface{}, now time.Time, accountId string) *TestItemResult {
	ret := &TestItemResult{
		Variable: item.Variable,
	}

	if item.StackFunc == nil {
		ret.Error = "Функция не задана"
		return ret
	}

	ret.BaseType = item.StackFunc.BaseType

	fnc := t.sfr.Get(item.StackFunc.Name)
	if fnc == nil {
		ret.Error = "Функция не существует"
		return ret
	}

	options := NewOptions(item.StackFunc.Arguments, varIdx)
	value, err := fnc.Run(options, now, accountId, true)
	if err != nil {
		ret.Error = fmt.Sprintf("Ошибка выполнения: %s", err)
		return ret
	}

	ret.Result = value
	return ret
}
