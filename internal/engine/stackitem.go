package engine

type StackItem struct {
	Variable  string
	StackFunc *StackFunc
}

type SetStackItem struct {
	Variable  string
	StackFunc *SetStackFunc
}
