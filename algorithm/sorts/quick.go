package sorts

// 快速排序
func Quick(array []int, low, high int) []int {
	if len(array) <= 1 {
		return array
	}
	if low < high {
		pivot := Partition(array, low, high)
		Quick(array, low, pivot-1)
		Quick(array, pivot+1, high)
	}
	return array
}

func Partition(array []int, low, high int) int {
	index := low - 1
	pivotElement := array[high]
	for i := low; i < high; i++ {
		if array[i] <= pivotElement {
			index += 1
			array[index], array[i] = array[i], array[index]
		}
	}
	array[index+1], array[high] = array[high], array[index+1]
	return index + 1
}
