package sort

func QuickSortOne(arr []int) {
	quickSortOne(arr, 0, len(arr)-1)
}

func quickSortOne(arr []int, first, last int) {
	if first >= last {
		return
	}

	pos := partitionOne(arr, first, last)
	quickSortOne(arr, first, pos)
	quickSortOne(arr, pos+1, last)
}

func partitionOne(arr []int, first, last int) int {
	pivot := arr[first]
	i := first - 1
	j := last + 1

	for {
		for {
			i++
			if arr[i] >= pivot {
				break
			}
		}

		for {
			j--
			if arr[j] <= pivot {
				break
			}
		}

		if i >= j {
			return j
		}

		arr[i], arr[j] = arr[j], arr[i]
	}
}
