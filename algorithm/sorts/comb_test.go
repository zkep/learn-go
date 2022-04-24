package sorts

import (
	"testing"
)

func TestComb(t *testing.T) {
	array := []int{16, 25, 9, 15, 2, 82, 78, 31, 11, 36, -3}
	array = comb(array)
	t.Log(array)
	t.FailNow()
}
