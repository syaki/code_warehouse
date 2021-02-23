package algs4

import "testing"

func TestShell(t *testing.T) {
	arr := []int{3, 4, 7, 1, 0, 5, 2, 6}
	Shell(arr)
	if !IsSorted(arr) {
		t.Errorf("failed")
	}
}
