package linked_list

import "reflect"

type node[T any] struct {
	Value T
	Next  *node[T]
}

func newNode[T any](value T) *node[T] {
	return &node[T]{Value: value}
}

// LinkedList represents a singly linked list.
type LinkedList[T any] struct {
	head *node[T]
	tail *node[T]
}

// New returns an empty LinkedList.
func New[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

// Add appends the specified element to the end of this list.
func (l *LinkedList[T]) Add(value T) {
	newNode := newNode(value)
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else if l.head == l.tail {
		l.head.Next = newNode
		l.tail = newNode
	} else {
		l.tail.Next = newNode
		l.tail = newNode
	}
}

// AddFirst inserts the specified element at the beginning of this list.
func (l *LinkedList[T]) AddFirst(value T) {
	newNode := newNode(value)
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		newNode.Next = l.head
		l.head = newNode
	}
}

// AddLast appends the specified element to the end of this list.
// This method is equivalent to Add.
func (l *LinkedList[T]) AddLast(value T) {
	l.Add(value)
}

// GetFirst returns the first element in this list.
func (l *LinkedList[T]) GetFirst() (*T, bool) {
	if l.head == nil {
		return nil, false
	}
	return &l.head.Value, true
}

// GetLast returns the last element in this list.
func (l *LinkedList[T]) GetLast() (*T, bool) {
	if l.tail == nil {
		return nil, false
	}
	return &l.tail.Value, true
}

// Get returns the element at the specified position in this list.
func (l *LinkedList[T]) Get(index int) (*T, bool) {
	position := 0
	current := l.head
	for current != nil && position <= index {
		if position == index {
			return &current.Value, true
		}
		current = current.Next
		position++
	}
	return nil, false
}

// RemoveFirst removes the first element (the head) of the collection.
func (l *LinkedList[T]) RemoveFirst() {
	if l.head != nil {
		l.head = l.head.Next
	}
}

// RemoveLast removes the last element (the tail) of the collection.
func (l *LinkedList[T]) RemoveLast() {
	if l.head == nil {
		return
	}
	if l.head == l.tail {
		l.head = nil
		l.tail = nil
		return
	}

	current := l.head
	for current.Next != l.tail {
		current = current.Next
	}
	current.Next = nil
}

// Remove removes the first occurrence of the operationValue from the list, if it is present.
// Returns true if the element was found and removed, false otherwise.
func (l *LinkedList[T]) Remove(value T) {
	if l.head == nil {
		return
	} else if reflect.DeepEqual(l.head.Value, value) {
		l.head = l.head.Next
		return
	}

	current := l.head
	for current.Next != nil {
		if reflect.DeepEqual(current.Next.Value, value) {
			current.Next = current.Next.Next
			return
		}
		current = current.Next
	}
}

// RemoveAt removes the element at the specified position in this list.
func (l *LinkedList[T]) RemoveAt(index int) {
	if index == 0 {
		if l.head != nil {
			l.head = l.head.Next
		}
		return
	}
	position := 1
	current := l.head
	for current.Next != nil && position <= index {
		if position == index {
			current.Next = current.Next.Next
			return
		}
		current = current.Next
		position++
	}
}

// Slice returns a slice containing all the elements in this list in proper sequence (from first to last element).
func (l *LinkedList[T]) Slice() []T {
	var slice []T
	current := l.head
	for current != nil {
		slice = append(slice, current.Value)
		current = current.Next
	}
	return slice
}

// Reverse reverses the order of the elements in the list.
func (l *LinkedList[T]) Reverse() {
	var previous *node[T]
	current := l.head
	for current != nil {
		next := current.Next
		current.Next = previous
		previous = current
		current = next
	}
	temp := l.head
	l.head = previous
	l.tail = temp
}

// Size returns the number of elements in this list.
func (l *LinkedList[T]) Size() int {
	size := 0
	current := l.head
	for current != nil {
		current = current.Next
		size++
	}
	return size
}

// IsEmpty returns true if this list contains no elements.
func (l *LinkedList[T]) IsEmpty() bool {
	return l.head == nil
}

// Contains return true if the specified value is present in the list, false otherwise.
func (l *LinkedList[T]) Contains(value T) bool {
	current := l.head
	for current != nil {
		if reflect.DeepEqual(current.Value, value) {
			return true
		}
		current = current.Next
	}
	return false
}

// Clear removes all the elements from this list.
func (l *LinkedList[T]) Clear() {
	l.head = nil
	l.tail = nil
}
