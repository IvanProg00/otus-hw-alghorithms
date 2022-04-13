package sort

func InsertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		k := i

		for k > 0 && arr[k] < arr[k-1] {
			Swap(arr, k, k-1)
			k--
		}
	}
}
