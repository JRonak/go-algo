package sort

import (
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

func QuickSortParallel(data []interface{}, compare func(a, b interface{}) bool, procs int) []interface{} {
	q := new(quickSort)
	q.data = data
	q.compare = compare
	q.numProcs = int32(procs)
	q.count = 1
	q.wg.Add(1)
	q.sortParallel(0, len(data)-1, 1)
	q.wg.Wait()
	return q.data
}

func (this *quickSort) partition(low, high int) int {
	index := low + 1
	for i := low + 1; i <= high; i++ {
		if this.compare(this.data[low], this.data[i]) {
			this.data[index], this.data[i] = this.data[i], this.data[index]
			index++
		}
	}
	this.data[index-1], this.data[low] = this.data[low], this.data[index-1]
	return index - 1
}

func (this *quickSort) sort(low, high int) {
	for low < high {
		parition := this.partition(low, high)
		this.sort(low, parition-1)
		low = parition + 1
	}
}

func (this *quickSort) sortRandom(low, high int) {
	for low < high {
		r := this.rand.Int()%(high-low+1) + low
		this.data[low], this.data[r] = this.data[r], this.data[low]
		parition := this.partition(low, high)
		this.sortRandom(low, parition-1)
		low = parition + 1
	}
}

func (this *quickSort) sortParallel(low, high, status int) {

	for low < high {
		parition := this.partition(low, high)
		if atomic.LoadInt32(&this.count) < this.numProcs {
			atomic.AddInt32(&this.count, 1)
			go this.sortParallel(low, parition-1, 1)
		} else {
			this.sortParallel(low, parition-1, 0)
		}
		low = parition + 1
	}

	if status == 1 {
		atomic.AddInt32(&this.count, -1)
		if atomic.LoadInt32(&this.count) == 0 {
			this.wg.Done()
		}
	}
}
