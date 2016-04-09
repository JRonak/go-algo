package sort

type mergeSort struct {
	data    []interface{}
	compare func(a, b interface{}) bool
}

func MergeSort(data []interface{}, compare func(a, b interface{}) bool) []interface{} {
	sort := new(mergeSort)
	sort.compare = compare
	sort.data = data
	sort.sort(0, len(data)-1)
	return sort.data
}

func (this *mergeSort) merge(low, mid, high int) {
	temp := []interface{}{}
	i, j := low, mid+1
	for i <= mid && j <= high {
		if this.compare(this.data[j], this.data[i]) {
			temp = append(temp, this.data[i])
			i++
		} else {
			temp = append(temp, this.data[j])
			j++
		}
	}
	for ; i <= mid; i++ {
		temp = append(temp, this.data[i])
	}
	for ; j <= high; j++ {
		temp = append(temp, this.data[j])
	}
	j = 0
	for i = low; i <= high; i++ {
		this.data[i] = temp[j]
		j++
	}

}

func (this *mergeSort) sort(low, high int) {
	if low < high {
		mid := low + (high-low)/2
		this.sort(low, mid)
		this.sort(mid+1, high)
		this.merge(low, mid, high)
	}
}
