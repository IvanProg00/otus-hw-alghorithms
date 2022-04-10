package sort

func BubbleSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for k := i + 1; k < len(arr); k++ {
			if arr[i] > arr[k] {
				Swap(arr, i, k)
			}
		}
	}

	return arr
}
