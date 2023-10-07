package queue

import "testing"

func TestNew(t *testing.T) {
	queue := New[int]()
	if queue.Size() != 0 {
		t.Error("New() should return an empty queue")
	}
}

func TestQueue_Enqueue(t *testing.T) {
	queue := New[int]()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	if size := queue.Size(); size != 3 {
		t.Errorf("Expected 3 elements in the Queue but found %v", size)
	}
}

func TestQueue_Dequeue(t *testing.T) {
	queue := New[int]()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	if value := queue.Dequeue(); value != 1 {
		t.Errorf("Expected 3 but found %v", value)
	}
	if size := queue.Size(); size != 2 {
		t.Errorf("Expected 2 elements in the Queue after dequeue but found %v", size)
	}
}

func TestQueue_Peek(t *testing.T) {
	queue := New[int]()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	if value := queue.Peek(); value != 1 {
		t.Errorf("Expected 3 but found %v", value)
	}
	if size := queue.Size(); size != 3 {
		t.Errorf("Expected 3 elements in the Queue after peek but found %v", size)
	}
}

func TestQueue_Size(t *testing.T) {
	queue := New[int]()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	queue.Enqueue(5)

	if size := queue.Size(); size != 5 {
		t.Errorf("Expected 5 elements in the Queue but found %v", size)
	}
}

func TestQueue_IsEmpty(t *testing.T) {
	queue := New[int]()

	if isEmpty := queue.IsEmpty(); !isEmpty {
		t.Errorf("Expcted a new Queue to be empty but got non-empty Queue")
	}
	queue.Enqueue(1)
	if isEmpty := queue.IsEmpty(); isEmpty {
		t.Errorf("Expected the Queue to be non-empty after enqueue an element but got empty Queue")
	}
}

func TestQueue_Clear(t *testing.T) {
	queue := New[string]()
	queue.Enqueue("a")
	queue.Enqueue("b")
	queue.Enqueue("c")
	queue.Enqueue("d")
	queue.Clear()

	if size := queue.Size(); size != 0 {
		t.Error("Expected an empty Queue after clear but got non-empty Queue")
	}

	queue.Enqueue("a")
	if size := queue.Size(); size != 1 {
		t.Errorf("Expected 1 element in the Queue after enqueue but found %v", size)
	}
}
