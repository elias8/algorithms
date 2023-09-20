package stack

import "testing"

func TestNew(t *testing.T) {
	stack := New[int]()
	if stack.Size() != 0 {
		t.Error("New() should return an empty stack")
	}
}

func TestStack_Push(t *testing.T) {
	stack := New[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	if size := stack.Size(); size != 3 {
		t.Errorf("Expected 3 elements in the Stack but found %v", size)
	}
}

func TestStack_Pop(t *testing.T) {
	stack := New[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	if value := stack.Pop(); value != 3 {
		t.Errorf("Expected 3 but found %v", value)
	}
	if size := stack.Size(); size != 2 {
		t.Errorf("Expected 2 elements in the Stack after pop but found %v", size)
	}
}

func TestStack_Peek(t *testing.T) {
	stack := New[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	if value := stack.Peek(); value != 3 {
		t.Errorf("Expected 3 but found %v", value)
	}
	if size := stack.Size(); size != 3 {
		t.Errorf("Expected 3 elements in the Stack after peek but found %v", size)
	}
}

func TestStack_Size(t *testing.T) {
	stack := New[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)

	if size := stack.Size(); size != 5 {
		t.Errorf("Expected 5 elements in the Stack but found %v", size)
	}
}

func TestStack_IsEmpty(t *testing.T) {
	stack := New[int]()

	if isEmpty := stack.IsEmpty(); !isEmpty {
		t.Errorf("Expcted a new Stack to be empty but got non-empty Stack")
	}
	stack.Push(1)
	if isEmpty := stack.IsEmpty(); isEmpty {
		t.Errorf("Expected the Stack to be non-empty after push an element but got empty Stack")
	}
}

func TestStack_Clear(t *testing.T) {
	stack := New[string]()
	stack.Push("a")
	stack.Push("b")
	stack.Push("c")
	stack.Push("d")
	stack.Clear()

	if size := stack.Size(); size != 0 {
		t.Error("Expected an empty Stack after clear but got non-empty Stack")
	}

	stack.Push("a")
	if size := stack.Size(); size != 1 {
		t.Errorf("Expected 1 element in the Stack after push but found %v", size)
	}
}
