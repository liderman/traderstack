package engine

import (
	"sync"
)

type StackFuncRepository struct {
	repo map[string]StackFuncRun

	mu sync.RWMutex
}

func NewStackFuncRepository() *StackFuncRepository {
	return &StackFuncRepository{
		repo: map[string]StackFuncRun{},
	}
}

func (s *StackFuncRepository) Register(sf StackFuncRun) {
	s.mu.Lock()
	s.repo[sf.Name()] = sf
	s.mu.Unlock()
}

func (s *StackFuncRepository) Get(name string) StackFuncRun {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.repo[name]
}

func (s *StackFuncRepository) GetAll() []StackFuncRun {
	s.mu.RLock()
	defer s.mu.RUnlock()

	ret := make([]StackFuncRun, 0, len(s.repo))
	for _, v := range s.repo {
		ret = append(ret, v)
	}

	return ret
}

func (s *StackFuncRepository) GetDeclaration(name string) *StackFunc {
	return s.funcToDeclaration(s.Get(name))
}

func (s *StackFuncRepository) GetAllDeclaration() []*StackFunc {
	funcs := s.GetAll()
	ret := make([]*StackFunc, 0, len(funcs))

	for _, v := range funcs {
		ret = append(ret, s.funcToDeclaration(v))
	}

	return ret
}

func (s *StackFuncRepository) funcToDeclaration(fnc StackFuncRun) *StackFunc {
	if fnc == nil {
		return nil
	}

	return &StackFunc{
		Name:      fnc.Name(),
		Arguments: fnc.Arguments(),
		BaseType:  fnc.BaseType(),
	}
}
