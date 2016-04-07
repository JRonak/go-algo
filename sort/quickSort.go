package sort

import (
	//	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

type quickSort struct {
	data     []interface{}
	compare  func(a, b interface{}) bool
	rand     *rand.Rand
	count    int32
	numProcs int32
	wg       sync.WaitGroup
}

func QuickSort(data []interface{}, compare func(a, b interface{}) bool) []interface{} {
	q := quickSort{}
	q.data = data
	q.compare = compare
	q.sort(0, len(data)-1)
	return q.data
}

func QuickSortRandomize(data []interface{}, compare func(a, b interface{}) bool) []interface{} {
	q := quickSort{}
	q.data = data
	q.compare = compare
	q.rand = rand.New(rand.NewSource(time.Now().Unix()))
	q.sortRandom(0, len(data)-1)
	return q.data
}

func QuickSortParallel(data []interface{}, compare func(a, b interface{}) bool) []interface{} {
	q := new(quickSort)
	q.data = data
	q.compare = compare
	q.numProcs = 8
	q.count = 1
	q.wg.Add(1)
	q.sortRandomP(0, len(data)-1, 1)
	q.wg.Wait()
	return q.data
}

func (this *quickSort) partition(low, high int) int {
	index := low - 1
	for i := low; i <= high; i++ {
		if this.compare(this.data[low], this.data[i]) {
			index++
			this.data[index], this.data[i] = this.data[i], this.data[index]
		}
	}
	this.data[index], this.data[low] = this.data[low], this.data[index]
	return index
}

func (this *quickSort) sort(low, high int) {
	if low < high {
		parition := this.partition(low, high)
		this.sort(low, parition-1)
		this.sort(parition+1, high)
	}
}

func (this *quickSort) sortRandom(low, high int) {
	if low < high {

		r := this.rand.Int()%(high-low+1) + low
		this.data[low], this.data[r] = this.data[r], this.data[low]
		parition := this.partition(low, high)
		this.sort(low, parition-1)
		this.sort(parition+1, high)
	}
}

func (this *quickSort) sortRandomP(low, high, status int) {

	if low < high {
		parition := this.partition(low, high)
		if atomic.LoadInt32(&this.count) < this.numProcs {
			atomic.AddInt32(&this.count, 1)
			go this.sortRandomP(low, parition-1, 1)
		} else {
			this.sortRandomP(low, parition-1, 0)
		}
		this.sortRandomP(parition+1, high, 0)
	}

	if status == 1 {
		atomic.AddInt32(&this.count, -1)
		if atomic.LoadInt32(&this.count) == 0 {
			this.wg.Done()
		}
	}
}
