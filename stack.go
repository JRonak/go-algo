package containers

import (
	"errors"
)

var (
	EMPTYERROR = errors.New("Operation failed: Stack is empty")
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

func (this *Stack) Peek() interface{} {
	if this.index > 0 {
		return this.list[this.index-1]
	} else {
		panic(EMPTYERROR)
	}
}

func (this *Stack) Pop() interface{} {
	if this.index > 0 {
		this.index--
		data := this.list[this.index]
		this.list = this.list[:this.index]
		return data
	} else {
		panic(EMPTYERROR)
	}
}

func (this *Stack) Size() int64 {
	return this.index
}

func (this *Stack) Data() []interface{} {
	return this.list
}
