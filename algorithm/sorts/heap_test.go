package sorts

import "testing"

func TestHeap(t *testing.T) {
	array := []int{16, 25, 9, 15, 2, 82, 78, 31, 11, 36, -3}
	array = HeapSort(array)
	t.Log(array)
	t.FailNow()
}
