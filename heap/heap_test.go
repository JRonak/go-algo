package heap

import (
	"testing"
)

func compare(a, b interface{}) bool {
	if a.(int) > b.(int) {
		return true
	} else {
		return false
	}
}

func TestHeap(t *testing.T) {
	h := New(compare)
	h.Push(5)
	h.Push(100)
	h.Push(3)

	if h.Length() != 3 {
		t.Error("Heap size : invalid")
	}
	if h.Pop() != 3 {
		t.Error("Heap pop : invalid")
	}
	if h.Length() != 2 {
		t.Error("Heap size : invalid")
	}
	if h.Pop() != 5 {
		t.Error("Heap pop: invalid")
	}

	h.Push(1)
	if h.Peek() != 1 {
		t.Error("Heap peek: invalid")
	}

}
