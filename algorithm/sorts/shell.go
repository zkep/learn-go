package sorts

func Shell(array []int) []int {
	for d := int(len(array) / 2); d > 0; d /= 2 {
		for i := d; i < len(array); i++ {
			for j := i; j >= d && array[j-d] > array[j]; j -= d {
				array[j], array[j-d] = array[j-d], array[j]
			}
		}
	}
	return array
}
