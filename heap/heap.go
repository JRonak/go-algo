package heap

import "errors"

var HERR = errors.New("Operation failed due to empty heap")

type Heap struct {
	heap    []interface{}
	lenght  int64
	compare func(a, b interface{}) bool
}

func New(compare func(a, b interface{}) bool) *Heap {
	h := new(Heap)
	h.compare = compare
	return h
}

func NewHeapData(data []interface{}, compare func(a, b interface{}) bool) *Heap {
	h := new(Heap)
	h.compare = compare
	h.heap = data
	h.lenght = int64(len(h.heap))
	for i := int(h.Length()) / 2; i >= 0; i-- {
		h.heapify(i)
	}
	return h
}

func (this *Heap) Push(item interface{}) {
	this.heap = append(this.heap, item)
	child := this.lenght
	parent := (child - 1) / 2
	this.lenght++
	for parent != child && this.compare(this.heap[parent], this.heap[child]) {
		this.heap[parent], this.heap[child] = this.heap[child], this.heap[parent]
		child = parent
		parent = (parent - 1) / 2
	}
}

func (this *Heap) Peek() interface{} {
	if this.lenght > 0 {
		return this.heap[0]
	} else {
		panic(HERR)
	}
}

func (this *Heap) Pop() interface{} {
	if this.lenght > 1 {
		item := this.heap[0]
		this.heap[0] = this.heap[this.lenght-1]
		this.heap = this.heap[:this.lenght-1]
		this.lenght--
		this.heapify(0)
		return item
	} else if this.lenght == 1 {
		item := this.heap[0]
		this.heap = this.heap[:0]
		this.lenght--
		return item
	} else {
		panic(HERR)
	}
}

func (this *Heap) heapify(index int) {
	child1 := index*2 + 1
	child2 := index*2 + 2
	originalIndex := index
	if child1 < int(this.lenght) && this.compare(this.heap[index], this.heap[child1]) {
		originalIndex = child1
	}
	if child2 < int(this.lenght) && this.compare(this.heap[originalIndex], this.heap[child2]) {
		originalIndex = child2
	}
	if originalIndex != index {
		this.heap[index], this.heap[originalIndex] = this.heap[originalIndex], this.heap[index]
		this.heapify(originalIndex)
	}

}

func (this *Heap) Length() int64 {
	return this.lenght
}

func (this *Heap) Data() []interface{} {
	return this.heap
}

func (this *Heap) Empty() bool {
	if this.lenght > 0 {
		return false
	} else {
		return true
	}
}
