package sorts

// 计数排序
func count(array []int) []int {
	var max, min = Imax(array), Imin(array)
	numbers := make([]int, max-min+1)
	for _, value := range array {
		numbers[value-min]++
	}
	z := 0
	for i, c := range numbers {
		for c > 0 {
			array[z] = i + min
			z++
			c--
		}
	}
	return array
}

func Imax(array []int) int {
	max := array[0]
	for _, v := range array {
		if v > max {
			max = v
		}
	}
	return max
}

func Imin(array []int) int {
	min := array[0]
	for _, v := range array {
		if min > v {
			min = v
		}
	}
	return min
}
