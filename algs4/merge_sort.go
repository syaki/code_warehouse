package algs4

// merge a[lo .. mid] with a[mid+1 ..hi] using aux[lo .. hi]
func merge(a, aux []int, lo, mid, hi int) {
	i := lo
	j := mid + 1

	// copy to aux[]
	for k := lo; k <= hi; k++ {
		aux[k] = a[k]
	}

	// merge back to a[]
	for k := lo; k <= hi; k++ {
		if i > mid {
			a[k] = aux[j]
			j++
		} else if j > hi {
			a[k] = aux[i]
			i++
		} else if aux[i] < aux[j] {
			a[k] = aux[i]
			i++
		} else {
			a[k] = aux[j]
			j++
		}
	}
}

// mergeSort a[lo..hi] using auxiliary array aux[lo..hi]
func mergeSort(a, aux []int, lo, hi int) {
	if hi <= lo {
		return
	}
	mid := lo + (hi-lo)/2
	mergeSort(a, aux, lo, mid)
	mergeSort(a, aux, mid+1, hi)
	merge(a, aux, lo, mid, hi)
}

// Sort rearranges the array
func Sort(a []int) {
	aux := make([]int, len(a))
	mergeSort(a, aux, 0, len(a)-1)
}
