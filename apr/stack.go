package apr

type Stack struct {
	elements []interface{}
}

func NewStack(n int) *Stack {
	return &Stack{elements: make([]interface{}, 0, n)}
}

func (s *Stack) Push(e interface{}) {
	s.elements = append(s.elements, e)
}

func (s *Stack) Append(stack *Stack) {
	s.elements = append(s.elements, stack.elements...)
}

func (s *Stack) Pop() interface{} {
	size := s.Size()
	if size == 0 {
		return nil
	}
	lastElement := s.elements[size-1]
	s.elements[size-1] = nil
	s.elements = s.elements[:size-1]
	return lastElement
}

func (s *Stack) Top() interface{} {
	size := s.Size()
	if size == 0 {
		return nil
	}
	return s.elements[size-1]
}

func (s *Stack) Index(index int) interface{} {
	size := s.Size()
	if index < size {
		return s.elements[index]
	}
	return nil
}

func (s *Stack) Clear() bool {
	if s.IsEmpty() {
		return false
	}
	for i := 0; i < s.Size(); i++ {
		s.elements[i] = nil
	}
	s.elements = make([]interface{}, 0)
	return true
}

func (s *Stack) Size() int {
	return len(s.elements)
}

func (s *Stack) IsEmpty() bool {
	if len(s.elements) == 0 {
		return true
	}
	return false
}

func (s *Stack) Clone() *Stack {
	stack := Stack{
		elements: make([]interface{}, s.Size()),
	}
	copy(stack.elements, s.elements)
	return &stack
}
