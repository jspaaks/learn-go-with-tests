package stack

type Stack[T any] struct {
	values []T
}

// IsEmpty checks whether the stack contains any values.
func (s *Stack[T]) IsEmpty() bool {
	return len(s.values) == 0
}

// Pop removes the most recently added value elem from the stack and returns it along with the
// ok signal true. In case the stack was empty, Pop returns the zero value for type T, along with
// the ok signal false.
func (s *Stack[T]) Pop() (elem T, ok bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}
	index := len(s.values) - 1
	elem = s.values[index]
	s.values = s.values[:index]
	return elem, true
}

// Push 'value' onto the stack.
func (s *Stack[T]) Push(value T) {
	s.values = append(s.values, value)
}

// NewStack instantiates a new stack.
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}
