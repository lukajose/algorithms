package stacks

type Stack struct {
	arr  []interface{}
	curr int
}

func NewStack(N int) *Stack {
	var stack = &Stack{
		arr: make([]interface{}, N),
	}
	return stack
}

func (s *Stack) IsEmpty() bool {
	return len(s.arr) == 0
}

func (s *Stack) Size() int {
	return len(s.arr)
}

func (s *Stack) Push(el interface{}) {
	s.arr = append(s.arr, el)
}

func (s *Stack) Pop() (interface{}, bool) {
	ok := false
	var el interface{}
	if s.IsEmpty() {
		return el, ok
	}
	index := len(s.arr) - 1
	el = s.arr[index]
	s.arr = s.arr[:index]
	return el, true
}

func (s *Stack) Next() (interface{}, bool) {

	if len(s.arr) == 0 {
		return nil, false
	}
	cIndex := s.curr
	s.curr = (s.curr + 1) % len(s.arr)
	return s.arr[cIndex], true
}
