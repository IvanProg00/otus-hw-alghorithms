package sort

func SelectionSort(arr []int) {
	for i := range arr {
		min := i
		for k := i + 1; k < len(arr); k++ {
			if arr[k] < arr[min] {
				min = k
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
}
