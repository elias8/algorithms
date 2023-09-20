package stack

type Stack[T any] []T

func New[T any]() *Stack[T] {
	return &Stack[T]{}
}

// Push adds a new element to the collection.
func (s *Stack[T]) Push(value T) {
	*s = append(*s, value)
}

// Pop removes the most recently added element.
func (s *Stack[T]) Pop() T {
	value := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return value
}

// Peek returns the last added element. Unlike Pop, this doesn't remove the element from the collection.
func (s *Stack[T]) Peek() T {
	return (*s)[len(*s)-1]
}

// Size returns the size of the collection.
func (s *Stack[T]) Size() int {
	return len(*s)
}

// IsEmpty returns `true` if the length of the collection is zero.
func (s *Stack[T]) IsEmpty() bool {
	return len(*s) == 0
}

// Clear removes all elements from the collection.
func (s *Stack[T]) Clear() {
	*s = []T{}
}
