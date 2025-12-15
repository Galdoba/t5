package rollresult

type RollResult []int

func New(r ...int) RollResult {
	return RollResult(r)
}

func (rr RollResult) Sum() int {
	s := 0
	for _, r := range rr {
		s += r
	}
	return s
}

func (rr RollResult) Results() []int {
	return rr
}
