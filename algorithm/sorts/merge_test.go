package sorts

import (
	"testing"
)

func TestMerge(t *testing.T) {
	array := []int{16, 25, 9, 15, 2, 82, 78, 31, 11, 36, -3}
	array = Merge(array)
	t.Log(array)
	t.FailNow()
}
