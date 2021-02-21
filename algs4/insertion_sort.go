package algs4

// Insertion rearranges the array arr []int
func Insertion(arr []int) {
	arrLen := len(arr)
	for i := 1; i < arrLen; i++ {
		for j := i; j > 0; j-- {
			if arr[j-1] < arr[j] {
				break
			}
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
}
