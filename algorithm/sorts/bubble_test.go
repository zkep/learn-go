package sorts

import (
	"testing"
)

func TestBubbleSort(t *testing.T) {
	array := []int{16, 25, 9, 15, 2, 82, 78, 31, 11, 36, -3}
	array = BubbleSort(array)
	t.Log(array)
	t.FailNow()
}
