package sorts

// 冒泡排序
func BubbleSort(array []int) []int {
	// 第一个循环代表排序次数
	// 第二个循环代表（总数-循环次数）之后的数据的排序
	// 每一次最外层的循环之后会把最大的数据往最后排
	length := len(array) - 1
	for i := 0; i < length; i++ {
		for j := 0; j < length-i; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array
}
