package sort

func HeapSort(arr []int) {
	for i := len(arr)/2 + 1; i >= 0; i-- {
		moveToRoot(arr, i, len(arr))
	}

	for i := len(arr) - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		moveToRoot(arr, 0, i)
	}
}

func moveToRoot(arr []int, root, size int) {
	l := root*2 + 1
	r := root*2 + 2
	x := root

	if l < size && arr[l] > arr[x] {
		x = l
	}
	if r < size && arr[r] > arr[x] {
		x = r
	}

	if x == root {
		return
	}

	arr[x], arr[root] = arr[root], arr[x]
	moveToRoot(arr, x, size)
}
