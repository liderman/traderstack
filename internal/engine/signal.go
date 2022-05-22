package engine

type Signal struct {
	// buy, sell
	Action string
	// short или long
	IsShort bool
}

type SignalInt interface {
	Calc() *Signal
}

type SignalAndWrap struct {
	data []SignalInt
}

func NewSignalAndWrap(data []SignalInt) *SignalAndWrap {
	return &SignalAndWrap{
		data: data,
	}
}

func (s *SignalAndWrap) Calc() *Signal {
	var first *Signal
	for _, v := range s.data {
		res := v.Calc()
		if res == nil {
			return nil
		}

		if first == nil {
			first = res
			continue
		}

		if res.IsShort != first.IsShort || res.Action != first.Action {
			return nil
		}
	}

	return first
}

type SignalAndOr struct {
	data []SignalInt
}

func NewSignalAndOr(data []SignalInt) *SignalAndOr {
	return &SignalAndOr{
		data: data,
	}
}

func (s *SignalAndOr) Calc() *Signal {
	for _, v := range s.data {
		res := v.Calc()
		if res != nil {
			return res
		}
	}

	return nil
}
