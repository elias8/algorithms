package queue

type Queue[T any] []T

func New[T any]() *Queue[T] {
	return &Queue[T]{}
}

// Enqueue adds a new element to the collection.
func (q *Queue[T]) Enqueue(value T) {
	*q = append(*q, value)
}

// Dequeue removes the first added element.
func (q *Queue[T]) Dequeue() T {
	value := (*q)[0]
	*q = (*q)[1:len(*q)]
	return value
}

// Peek returns the first added element. Unlike Dequeue, this doesn't remove the element from the collection.
func (q *Queue[T]) Peek() T {
	return (*q)[0]
}

// Size returns the size of the collection.
func (q *Queue[T]) Size() int {
	return len(*q)
}

// IsEmpty returns `true` if the length of the collection is zero.
func (q *Queue[T]) IsEmpty() bool {
	return len(*q) == 0
}

// Clear removes all elements from the collection.
func (q *Queue[T]) Clear() {
	*q = []T{}
}
