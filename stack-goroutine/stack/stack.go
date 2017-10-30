package stack

// Stack is a stack struct.
type Stack struct {
	vals  []string
	limit int
}

// NewStack constructs Stack
func NewStack(limit int) Stack {
	return Stack{vals: make([]string, 0), limit: limit}
}

// Push pushes arg value.
func (s *Stack) Push(str string) {
	if (len(s.vals) + 1) > s.limit {
		pushed := make([]string, 0)
		for i := 1; i < len(s.vals); i++ {
			pushed = append(pushed, s.vals[i])
		}
		pushed = append(pushed, str)
		s.vals = pushed
	} else {
		s.vals = append(s.vals, str)
	}
}

// Pop pops from vals
func (s *Stack) Pop() string {
	removeLast := func() []string {
		removed := make([]string, 0)
		for i := 0; i < len(s.vals)-1; i++ {
			removed = append(removed, s.vals[i])
		}
		return removed
	}

	result := s.vals[len(s.vals)-1]
	s.vals = removeLast()
	return result
}
