package sorts

func Pigeonhole(array []int) []int {
	if len(array) == 0 {
		return array
	}
	min := Imin(array)
	max := Imax(array)
	size := max - min + 1
	holes := make([]int, size)
	for _, element := range array {
		holes[element-min]++
	}
	i := 0
	for j := 0; j < size; j++ {
		for holes[j] > 0 {
			holes[j]--
			array[i] = j + min
			i++
		}
	}

	return array
}
