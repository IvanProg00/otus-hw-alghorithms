package sort

func QuickSortTwo(arr []int) {
	quickSortTwo(arr, 0, len(arr)-1)
}

func quickSortTwo(arr []int, first, last int) {
	if first >= last {
		return
	}

	v := partitionTwo(arr, first, last)

	quickSortTwo(arr, first, v-1)
	quickSortTwo(arr, v+1, last)
}

func partitionTwo(arr []int, first, last int) int {
	for i := first; i <= last; i++ {
		if arr[i] < arr[last] {
			arr[i], arr[first] = arr[first], arr[i]
			first++
		}
	}

	arr[first], arr[last] = arr[last], arr[first]

	return first
}
