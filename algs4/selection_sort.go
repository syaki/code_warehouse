package algs4

// Selection rearranges array arr []int
func Selection(arr []int) {
	arrLen := len(arr)
	for i := 0; i < arrLen; i++ {
		minIdx := i
		for j := i + 1; j < arrLen; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}
