package sort

func InsertionSort(data []interface{}, compare func(a, b interface{}) bool) []interface{} {
	length := len(data)
	for i := 1; i < length; i++ {
		for j := i - 1; j >= 0 && compare(data[j], data[j+1]); j-- {
			data[j+1], data[j] = data[j], data[j+1]
		}
	}
	return data
}
