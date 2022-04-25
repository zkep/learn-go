package sorts

import (
	"testing"
)

func TestShell(t *testing.T) {
	array := []int{16, 25, 9, 15, 2, 82, 78, 31, 11, 36, -3}
	array = Shell(array)
	t.Log(array)
	t.FailNow()
}
