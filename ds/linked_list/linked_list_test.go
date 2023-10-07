package linked_list

import (
	"slices"
	"testing"
)

type scenario[T comparable] struct {
	name           string
	values         []T
	expected       []T
	operationValue int
}

func (s scenario[T]) testSize(t *testing.T, list *LinkedList[T]) {
	if size := list.Size(); size != len(s.expected) {
		t.Fatalf("Expected list size to be %d, but found %d", len(s.expected), size)
	}
}

func (s scenario[T]) testSlice(t *testing.T, list *LinkedList[T]) {
	if slice := list.Slice(); !slices.Equal(slice, s.expected) {
		t.Fatalf("Expected list to be %v, but found %v", s.expected, slice)
	}
}

func TestNew(t *testing.T) {
	list := New[int]()

	if isEmpty := list.IsEmpty(); !isEmpty {
		t.Errorf("Expected list to be empty, but found not")
	}
	if size := list.Size(); size != 0 {
		t.Errorf("Expected list size to be 0, but found %d", size)
	}
}

func TestLinkedList_Add(t *testing.T) {
	scenarios := []scenario[int]{
		{
			name:           "add to empty list",
			values:         []int{},
			expected:       []int{1},
			operationValue: 1,
		},
		{
			name:           "add to list with one element",
			values:         []int{1},
			expected:       []int{1, 2},
			operationValue: 2,
		},
		{
			name:           "add to list with multiple elements",
			values:         []int{1, 2, 3},
			expected:       []int{1, 2, 3, 4},
			operationValue: 4,
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			list := New[int]()
			for _, value := range scenario.values {
				list.Add(value)
			}

			list.Add(scenario.operationValue)
			scenario.testSize(t, list)
			scenario.testSlice(t, list)
		})
	}
}

func TestLinkedList_AddFirst(t *testing.T) {
	scenarios := []scenario[int]{
		{
			name:           "add to empty list",
			values:         []int{},
			expected:       []int{1},
			operationValue: 1,
		},
		{
			name:           "add to list with one element",
			values:         []int{1},
			expected:       []int{2, 1},
			operationValue: 2,
		},
		{
			name:           "add to list with multiple elements",
			values:         []int{1, 2, 3},
			expected:       []int{4, 1, 2, 3},
			operationValue: 4,
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			list := New[int]()
			for _, value := range scenario.values {
				list.Add(value)
			}

			list.AddFirst(scenario.operationValue)
			scenario.testSize(t, list)
			scenario.testSlice(t, list)
		})
	}
}

func TestLinkedList_AddLast(t *testing.T) {
	scenarios := []scenario[int]{
		{
			name:           "add to empty list",
			values:         []int{},
			expected:       []int{1},
			operationValue: 1,
		},
		{
			name:           "add to list with one element",
			values:         []int{1},
			expected:       []int{1, 2},
			operationValue: 2,
		},
		{
			name:           "add to list with multiple elements",
			values:         []int{1, 2, 3},
			expected:       []int{1, 2, 3, 4},
			operationValue: 4,
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			list := New[int]()
			for _, value := range scenario.values {
				list.Add(value)
			}

			list.AddLast(scenario.operationValue)
			scenario.testSize(t, list)
			scenario.testSlice(t, list)
		})
	}
}

func TestLinkedList_Get(t *testing.T) {
	list := New[int]()

	if value, found := list.Get(0); found {
		t.Fatalf("Expected value element to be not found, but found %d", value)
	}

	for i := 1; i <= 10; i++ {
		list.Add(i)
	}
	for i := 0; i < list.Size(); i++ {
		if value, found := list.Get(i); !found || *value != i+1 {
			t.Fatalf("Expected element at operationValue %d to be %d, but found %d", i, i+1, value)
		}
	}
}

func TestLinkedList_GetFirst(t *testing.T) {
	list := New[int]()

	if value, found := list.GetFirst(); found {
		t.Fatalf("Expected value element to be not found, but found %d", value)
	}

	for i := 1; i <= 10; i++ {
		list.Add(i)
	}
	if value, found := list.GetFirst(); !found || *value != 1 {
		t.Fatalf("Expected value element to be 1, but found %d", value)
	}
}

func TestLinkedList_GetLast(t *testing.T) {
	list := New[int]()

	if value, found := list.GetLast(); found {
		t.Fatalf("Expected value element to be not found, but found %d", value)
	}

	for i := 1; i <= 10; i++ {
		list.Add(i)
	}
	if value, found := list.GetLast(); !found || *value != 10 {
		t.Fatalf("Expected value element to be 10, but found %d", value)
	}
}

func TestLinkedList_RemoveFirst(t *testing.T) {
	scenarios := []scenario[int]{
		{
			name:     "remove from empty list",
			values:   []int{},
			expected: []int{},
		},
		{
			name:     "remove from list with one element",
			values:   []int{1},
			expected: []int{},
		},
		{
			name:     "remove from list with two elements",
			values:   []int{1, 2},
			expected: []int{2},
		},
		{
			name:     "remove from list with multiple elements",
			values:   []int{1, 2, 3},
			expected: []int{2, 3},
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			list := New[int]()
			for _, value := range scenario.values {
				list.Add(value)
			}

			list.RemoveFirst()
			scenario.testSize(t, list)
			scenario.testSlice(t, list)
		})
	}
}

func TestLinkedList_RemoveLast(t *testing.T) {
	scenarios := []scenario[int]{
		{
			name:     "remove from empty list",
			values:   []int{},
			expected: []int{},
		},
		{
			name:     "remove from list with one element",
			values:   []int{1},
			expected: []int{},
		},
		{
			name:     "remove from list with two elements",
			values:   []int{1, 2},
			expected: []int{1},
		},
		{
			name:     "remove from list with multiple elements",
			values:   []int{1, 2, 3},
			expected: []int{1, 2},
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			list := New[int]()
			for _, value := range scenario.values {
				list.Add(value)
			}

			list.RemoveLast()
			scenario.testSize(t, list)
			scenario.testSlice(t, list)
		})
	}
}

func TestLinkedList_Remove(t *testing.T) {
	scenarios := []scenario[int]{
		{
			name:           "remove from empty list",
			values:         []int{},
			operationValue: 0,
		},
		{
			name:           "remove from list with one element",
			operationValue: 1,
			values:         []int{1},
			expected:       []int{},
		},
		{
			name:           "remove first element from list with two elements",
			operationValue: 1,
			values:         []int{1, 2},
			expected:       []int{2},
		},
		{
			name:           "remove last element from list with two elements",
			operationValue: 2,
			values:         []int{1, 2},
			expected:       []int{1},
		},
		{
			name:           "remove first element from list with multiple elements",
			operationValue: 1,
			values:         []int{1, 2, 3},
			expected:       []int{2, 3},
		},
		{
			name:           "remove middle element from list with multiple elements",
			operationValue: 2,
			values:         []int{1, 2, 3},
			expected:       []int{1, 3},
		},
		{
			name:           "remove last element from list with multiple elements",
			operationValue: 3,
			values:         []int{1, 2, 3},
			expected:       []int{1, 2},
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			list := New[int]()
			for _, value := range scenario.values {
				list.Add(value)
			}

			list.Remove(scenario.operationValue)
			scenario.testSize(t, list)
			scenario.testSlice(t, list)
		})
	}
}

func TestLinkedList_RemoveAt(t *testing.T) {
	scenarios := []scenario[int]{
		{
			name:           "remove from empty list",
			values:         []int{},
			operationValue: 0,
		},
		{
			name:           "remove from list with one element",
			operationValue: 0,
			values:         []int{1},
			expected:       []int{},
		},
		{
			name:           "remove first element from list with two elements",
			operationValue: 0,
			values:         []int{1, 2},
			expected:       []int{2},
		},
		{
			name:           "remove last element from list with two elements",
			operationValue: 1,
			values:         []int{1, 2},
			expected:       []int{1},
		},
		{
			name:           "remove first element from list with multiple elements",
			operationValue: 0,
			values:         []int{1, 2, 3},
			expected:       []int{2, 3},
		},
		{
			name:           "remove middle element from list with multiple elements",
			operationValue: 1,
			values:         []int{1, 2, 3},
			expected:       []int{1, 3},
		},
		{
			name:           "remove last element from list with multiple elements",
			operationValue: 2,
			values:         []int{1, 2, 3},
			expected:       []int{1, 2},
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			list := New[int]()
			for _, value := range scenario.values {
				list.Add(value)
			}

			list.RemoveAt(scenario.operationValue)
			scenario.testSize(t, list)
			scenario.testSlice(t, list)
		})
	}
}

func TestLinkedList_Slice(t *testing.T) {
	list := New[int]()

	for i := 1; i <= 10; i++ {
		list.Add(i)
		expected := make([]int, i)
		for j := 0; j < i; j++ {
			expected[j] = j + 1
		}

		if slice := list.Slice(); !slices.Equal(slice, expected) {
			t.Fatalf("Expected list to be %v, but found %v", expected, slice)
		}
	}
}

func TestLinkedList_Reverse(t *testing.T) {
	scenarios := []scenario[int]{
		{
			name:     "reverse empty list",
			values:   []int{},
			expected: []int{},
		},
		{
			name:     "reverse list with one element",
			values:   []int{1},
			expected: []int{1},
		},
		{
			name:     "reverse list with two elements",
			values:   []int{1, 2},
			expected: []int{2, 1},
		},
		{
			name:     "reverse list with multiple elements",
			values:   []int{1, 2, 3},
			expected: []int{3, 2, 1},
		},
		{
			name:     "reverse list with many elements",
			values:   []int{1, 2, 3, 4},
			expected: []int{4, 3, 2, 1},
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			list := New[int]()
			for _, value := range scenario.values {
				list.Add(value)
			}

			list.Reverse()
			scenario.testSize(t, list)
			scenario.testSlice(t, list)

			if len(scenario.expected) > 0 {
				if value, found := list.GetFirst(); !found || *value != scenario.expected[0] {
					t.Fatalf("Expected first element to be %d, but found %d", scenario.expected[0], *value)
				}

				if value, found := list.GetLast(); !found || *value != scenario.expected[len(scenario.expected)-1] {
					t.Fatalf("Expected last element to be %d, but found %d", scenario.expected[len(scenario.expected)-1], *value)
				}
			}
		})
	}
}

func TestLinkedList_Size(t *testing.T) {
	list := New[int]()

	for i := 1; i <= 10; i++ {
		list.Add(i)
		if size := list.Size(); size != i {
			t.Fatalf("Expected list size to be %d, but found %d", i, size)
		}
	}
}

func TestLinkedList_IsEmpty(t *testing.T) {
	list := New[int]()

	if isEmpty := list.IsEmpty(); !isEmpty {
		t.Errorf("Expected list to be empty, but found not")
	}

	list.Add(1)
	if isEmpty := list.IsEmpty(); isEmpty {
		t.Errorf("Expected list to be not empty, but found empty")
	}
}

func TestLinkedList_Contains(t *testing.T) {
	list := New[int]()

	if contains := list.Contains(1); contains {
		t.Errorf("Expected list to be not contain element, but found contains")
	}

	for i := 1; i <= 10; i++ {
		list.Add(i)
		for j := 1; j <= i; j++ {
			if contains := list.Contains(j); !contains {
				t.Errorf("Expected list to be contain element %d, but found not", j)
			}
		}
	}
}

func TestLinkedList_Clear(t *testing.T) {
	list := New[int]()

	for i := 1; i <= 10; i++ {
		list.Add(i)
	}

	list.Clear()
	if isEmpty := list.IsEmpty(); !isEmpty {
		t.Errorf("Expected list to be empty, but found elements")
	}
	if size := list.Size(); size != 0 {
		t.Errorf("Expected list size to be 0, but found %d", size)
	}
}
