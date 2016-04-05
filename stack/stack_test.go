package stack

import (
	"testing"
)

func TestStack(t *testing.T) {
	tstack := NewStack()

	tstack.Push(5)
	tstack.Push(4)
	tstack.Push(3)

	if tstack.Size() != 3 {
		t.Error("Size value wrong")
	}

	if tstack.Peek() != 3 {
		t.Error("Wrong peek value")
	}

	tstack.Pop()

	if tstack.Size() != 2 {
		t.Error("Size value wrong")
	}

	if tstack.Peek() != 4 {
		t.Error("Wrong peek value")
	}
}
