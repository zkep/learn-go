package sorts

import (
	"testing"
)

func TestQuick(t *testing.T) {
	array := []int{16, 25, 9, 15, 2, 82, 78, 31, 11, 36, -3}
	array = Quick(array, 0, len(array)-1)
	t.Log(array)
	t.FailNow()
}
