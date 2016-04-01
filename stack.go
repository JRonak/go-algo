package containers

import (
	"errors"
)

type Stack struct {
	list  []interface{}
	index int64
}

func NewStack() *Stack {
	return new(Stack)
}

func (this *Stack) Push(data interface{}) {
	this.list = append(this.list, data)
	this.index++
}

func (this *Stack) Empty() bool {
	if this.index > 0 {
		return false
	}
	return true
}

func (this *Stack) Pop() (interface{}, error) {
	if !this.Empty() {
		this.index--
		return this.list[this.index], nil
	} else {
		return 1, errors.New("Stack is empty")
	}
}

func (this *Stack) Size() int64 {
	return this.index
}

func (this *Stack) Data() []interface{} {
	return this.list
}
