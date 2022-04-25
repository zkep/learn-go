package sorts

import (
	"testing"
)

func TestRadix(t *testing.T) {
	array := []int{16, 25, 9, 15, 2, 82, 78, 31, 11, 36, -3}
	array = Radix(array)
	t.Log(array)
	t.FailNow()
}
