package queue

import (
	"testing"
)

func TestQueue(t *testing.T) {
	q := New()
	q.Enqueue(5)
	q.Enqueue(4)
	q.Enqueue(3)
	q.Enqueue(2)

	if q.Size() != 4 {
		t.Error("Invalid Queue size")
	}
	if q.Peek() != 5 {
		t.Error("Invalid peek value")
	}

	q.Dequeue()
	if q.Peek() != 4 {
		t.Error("Invalid peek value")
	}

	q.Enqueue(1)
	q.Dequeue()
	q.Dequeue()
	q.Dequeue()

	if q.Size() != 1 {
		t.Error("Invalid size value")
	}

	if q.Peek() != 1 {
		t.Error("Invalid peek value")
	}
}
