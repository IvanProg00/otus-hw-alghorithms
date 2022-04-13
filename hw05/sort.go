package sort

func Swap[T any](arr []T, i, k int) {
	arr[i], arr[k] = arr[k], arr[i]
	// As same as below
	//
	// temp := arr[i]
	// arr[i] = arr[k]
	// arr[k] = temp
}
