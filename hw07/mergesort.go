package sort

func MergeSort(arr []int) {
	sort(arr, 0, len(arr)-1)
}

func sort(arr []int, first, last int) {
	if first < last {
		mid := first + (last-first)/2

		sort(arr, first, mid)
		sort(arr, mid+1, last)

		merge(arr, first, last, mid)
	}
}

func merge(arr []int, first, last, mid int) {
	n1 := mid - first + 1
	n2 := last - mid

	l := make([]int, 0, n1)
	r := make([]int, 0, n2)

	for i := 0; i < n1; i++ {
		l = append(l, arr[first+i])
	}

	for i := 0; i < n2; i++ {
		r = append(r, arr[mid+1+i])
	}

	var (
		i int
		j int
	)

	k := first

	for i < n1 && j < n2 {
		if l[i] <= r[j] {
			arr[k] = l[i]
			i++
		} else {
			arr[k] = r[j]
			j++
		}
		k++
	}

	for i < n1 {
		arr[k] = l[i]
		i++
		k++
	}

	for j < n2 {
		arr[k] = r[j]
		j++
		k++
	}
}
