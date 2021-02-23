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

// InsertionX rearranges the array that using half exchanges instead of full exchanges to reduce data movement.
func InsertionX(arr []int) {
	arrLen := len(arr)

	// put smallest element in position to serve as sentinel
	exchanges := 0
	for i := arrLen - 1; i > 0; i-- {
		if arr[i] < arr[i-1] {
			arr[i], arr[i-1] = arr[i-1], arr[i]
			exchanges++
		}
	}

	if exchanges == 0 {
		return
	}
	// insertion sort with half-exchanges
	for i := 2; i < arrLen; i++ {
		temp := arr[i]
		j := i
		for temp < arr[j-1] {
			arr[j] = arr[j-1]
			j--
		}
		arr[j] = temp
	}
}
