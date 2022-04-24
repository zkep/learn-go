package sorts

func getNextGap(gap int) int {
	gap = (gap * 10) / 13
	if gap < 1 {
		return 1
	}
	return gap
}

func comb(array []int) []int {
	length := len(array)
	gap := length
	swapped := true
	for gap != 1 || swapped {
		gap = getNextGap(gap)
		swapped = false
		for i := 0; i < length-gap; i++ {
			if array[i] > array[i+gap] {
				array[i], array[i+gap] = array[i+gap], array[i]
			}
		}
	}
	return array
}
