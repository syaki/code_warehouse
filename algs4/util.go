package algs4

// IsSorted return is the array arr []int sorted?
func IsSorted(arr []int) bool {
	arrLen := len(arr)
	for i := 1; i < arrLen; i++ {
		if arr[i] < arr[i-1] {
			return false
		}
	}
	return true
}
