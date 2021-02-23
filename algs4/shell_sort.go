package algs4

// Shell rearranges the array
func Shell(a []int) {
	n := len(a)

	// 3x+1 increment sequence:  1, 4, 13, 40, 121, 364, 1093, ...
	h := 1
	for h < n/3 {
		h = 3*h + 1
	}

	for h >= 1 {
		// h-sort the array
		for i := h; i < n; i++ {
			for j := i; j >= h && a[j] < a[j-h]; j -= h {
				a[j], a[j-h] = a[j-h], a[j]
			}
		}
		h = (h - 1) / 3
	}
}
