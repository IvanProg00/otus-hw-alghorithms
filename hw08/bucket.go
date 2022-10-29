package sort

func Bucket(arr []int) {
	if len(arr) == 0 {
		return
	}

	max := getMax(arr)
	buckets := make([][]int, len(arr))
	coefficientToGetPos := len(arr) / (max + 1)

	for i := range arr {
		pos := arr[i] * coefficientToGetPos
		subArr := buckets[pos]

		subArr = append(subArr, arr[i])

		insertionSort(subArr)
		buckets[pos] = subArr
	}

	copyArr(arr, buckets)
}

func getMax(arr []int) int {
	if len(arr) == 0 {
		panic("list is empty")
	}

	max := arr[0]

	for i := range arr {
		if max < arr[i] {
			max = arr[i]
		}
	}

	return max
}

func copyArr(dst []int, src [][]int) {
	i := 0

	for k := range src {
		for j := range src[k] {
			dst[i] = src[k][j]
			i++
		}
	}
}

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		j := i - 1
		for j >= 0 && arr[i] < arr[j] {
			arr[i], arr[j] = arr[j], arr[i]
			j--
			i--
		}
	}
}
