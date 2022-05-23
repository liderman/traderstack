package engine

import (
	"errors"
	"sync"
)

// StackManager Менеджер управления стеками.
type StackManager struct {
	sfr  *StackFuncRepository
	data map[string]*Stack
	mu   sync.Mutex
}

func NewStackManager(sfr *StackFuncRepository) *StackManager {
	return &StackManager{
		sfr:  sfr,
		data: map[string]*Stack{},
	}
}

func (s *StackManager) Create(name string) *Stack {
	stack := NewStack(name)

	s.mu.Lock()
	s.data[stack.Id] = stack
	s.mu.Unlock()

	return stack
}

func (s *StackManager) Update(id, name string, setItems []*SetStackItem) (*Stack, error) {
	s.mu.Lock()
	stack, ok := s.data[id]
	s.mu.Unlock()
	if !ok {
		return nil, errors.New("stack is not exists")
	}

	stack.Name = name
	items := make([]*StackItem, 0, len(setItems))
	for _, v := range setItems {
		fnc, err := s.mapStackFunc(v.StackFunc)
		if err != nil {
			return nil, err
		}
		items = append(items, &StackItem{
			Variable:  v.Variable,
			StackFunc: fnc,
		})
	}
	stack.Items = items

	s.mu.Lock()
	s.data[id] = stack
	s.mu.Unlock()

	return stack, nil
}

func (s *StackManager) Get(id string) *Stack {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.data[id]
}

func (s *StackManager) GetAll() []*Stack {
	s.mu.Lock()
	defer s.mu.Unlock()
	var ret []*Stack
	for _, v := range s.data {
		ret = append(ret, v)
	}
	return ret
}

func (s *StackManager) Delete(id string) {
	s.mu.Lock()
	delete(s.data, id)
	s.mu.Unlock()
}

func (s *StackManager) FuncArgumentVarList(stackId, itemVariable, argumentId string) []string {
	s.mu.Lock()
	stack, ok := s.data[stackId]
	s.mu.Unlock()

	if !ok {
		return nil
	}

	idx := -1
	var arg *Argument
	for k, v := range stack.Items {
		if v.Variable != itemVariable || v.StackFunc == nil {
			continue
		}

		arg = v.StackFunc.GetArgument(argumentId)
		if arg != nil {
			idx = k
			break
		}
	}

	if arg == nil {
		return nil
	}

	var ret []string
	for k := idx - 1; k > 0; k-- {
		item := stack.Items[k]
		if item.StackFunc == nil {
			continue
		}

		if arg.CheckInputBaseType(item.StackFunc.BaseType) {
			ret = append(ret, item.Variable)
		}
	}

	return ret
}

func (s *StackManager) mapStackFunc(in *SetStackFunc) (*StackFunc, error) {
	if in == nil {
		return nil, nil
	}
	fnc := s.sfr.GetDeclaration(in.Name)
	if fnc == nil {
		return nil, errors.New("function is not exist")
	}

	for _, v := range in.Arguments {
		err := fnc.SetArgument(v.Id, v.Value)
		if err != nil {
			return nil, err
		}
	}

	return fnc, nil
}
