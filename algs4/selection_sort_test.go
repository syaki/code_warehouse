package algs4

import (
	"testing"
)

func TestSelection(t *testing.T) {
	arr := []int{3, 4, 7, 1, 0, 5, 2, 6}
	Selection(arr)
	if !IsSorted(arr) {
		t.Errorf("failed")
	}
}
