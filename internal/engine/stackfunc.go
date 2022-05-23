package engine

import (
	"errors"
	"time"
)

// StackFuncRun Функция элемента стека для выполнения.
type StackFuncRun interface {
	Name() string
	BaseType() string
	Run(options *Options, now time.Time, accountId string, isTest bool) (interface{}, error)
	Arguments() []*Argument
}

type StackFunc struct {
	Name      string
	Arguments []*Argument
	BaseType  string
}

func (s *StackFunc) SetArgument(id string, value interface{}) error {
	for _, v := range s.Arguments {
		if v.Id == id {
			v.Value = value
			return nil
		}
	}

	return errors.New("argument is not exist")
}

func (s *StackFunc) GetArgument(id string) *Argument {
	for _, v := range s.Arguments {
		if v.Id == id {
			return v
		}
	}

	return nil
}

type SetStackFunc struct {
	Name      string
	Arguments []*SetArgument
}
