package sorts

func Merge(array []int) []int {
	if len(array) < 2 {
		return array
	}
	var middle = len(array) / 2
	var a = Merge(array[:middle])
	var b = Merge(array[middle:])
	return merge(a, b)
}

func merge(a, b []int) []int {
	var r = make([]int, len(a)+len(b))
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			r[i+j] = a[i]
			i++
		} else {
			r[i+j] = b[j]
			j++
		}
	}

	for i < len(a) {
		r[i+j] = a[i]
		i++
	}

	for j < len(b) {
		r[i+j] = b[j]
		j++
	}

	return r
}
