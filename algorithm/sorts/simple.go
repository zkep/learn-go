package sorts

func Simple(array []int) []int {
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array); j++ {
			if array[i] < array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
	return array
}
