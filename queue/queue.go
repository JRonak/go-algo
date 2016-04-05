package queue

import "errors"

var (
	QERR = errors.New("Operation cannot be carried out as the queue is empty")
)

type Queue struct {
	data   []interface{}
	length int64
}

func New() *Queue {
	return new(Queue)
}

func (this *Queue) Enqueue(item interface{}) {
	this.data = append(this.data, item)
	this.length++
}

func (this *Queue) Dequeue() interface{} {
	if this.length > 0 {
		item := this.data[0]
		this.data = this.data[1:this.length]
		this.length--
		return item
	} else {
		panic(QERR)
	}
}

func (this *Queue) Peek() interface{} {
	if this.length > 0 {
		return this.data[0]
	} else {
		panic(QERR)
	}
}

func (this *Queue) Size() int64 {
	return this.length
}

func (this *Queue) Data() []interface{} {
	return this.data
}

func (this *Queue) Empty() bool {
	if this.length == 0 {
		return true
	} else {
		return false
	}
}
