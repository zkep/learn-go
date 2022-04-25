package sorts

func Insertion(array []int) []int {
	for currentIndex := 1; currentIndex < len(array); currentIndex++ {
		temp := array[currentIndex]
		iterator := currentIndex

		for ; iterator > 0 && array[iterator-1] >= temp; iterator-- {
			array[iterator] = array[iterator-1]
		}
		array[iterator] = temp
	}
	return array
}
