package engine

import (
	"github.com/google/uuid"
)

type Stack struct {
	Id    string
	Name  string
	Items []*StackItem
}

func NewStack(name string) *Stack {
	return &Stack{
		Id:    uuid.New().String(),
		Name:  name,
		Items: []*StackItem{},
	}
}
