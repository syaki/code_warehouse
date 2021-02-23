package algs4

import "testing"

func TestInsertion(t *testing.T) {
	t.Run("Test Insertion", func(t *testing.T) {
		arr := []int{3, 4, 7, 1, 0, 5, 2, 6}
		Insertion(arr)
		if !IsSorted(arr) {
			t.Errorf("failed")
		}
	})
	t.Run("Test InsertionX", func(t *testing.T) {
		arr := []int{3, 4, 7, 1, 0, 5, 2, 6}
		InsertionX(arr)
		if !IsSorted(arr) {
			t.Errorf("failed")
		}
	})
}
