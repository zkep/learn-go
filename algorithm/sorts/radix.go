package sorts

func Radix(array []int) []int {
	if len(array) < 1 {
		return array
	}
	var negatives, nonNegatives []int
	for _, element := range array {
		if element < 0 {
			negatives = append(negatives, -element)
		} else {
			nonNegatives = append(nonNegatives, element)
		}
	}
	negatives = unsignedRadix(negatives)

	for i, j := 0, len(negatives)-1; i <= j; i, j = i+1, j-1 {
		negatives[i], negatives[j] = -negatives[j], -negatives[i]
	}

	return append(negatives, unsignedRadix(nonNegatives)...)
}

func unsignedRadix(array []int) []int {
	if len(array) == 0 {
		return array
	}
	maxElement := Imax(array)
	for exp := 1; maxElement/exp > 0; exp *= 10 {
		array = countRadix(array, exp)
	}
	return array
}

func countRadix(array []int, exp int) []int {
	var digits [10]int
	var output = make([]int, len(array))

	for _, item := range array {
		digits[(item/exp)%10]++
	}

	for i := 1; i < 10; i++ {
		digits[i] += digits[i-1]
	}

	for i := len(array) - 1; i >= 0; i-- {
		output[digits[(array[i]/exp)%10]-1] = array[i]
		digits[(array[i]/exp)%10]--
	}

	return output
}
