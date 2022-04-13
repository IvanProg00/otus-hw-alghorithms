package sort

func ShellSort(arr []int) {
	for gap := len(arr) / 2; gap > 0; gap /= 2 {
		for i := 0; i+gap < len(arr); i++ {
			j := i + gap
			jTemp := arr[j]

			for j-gap >= 0 && arr[j-gap] > jTemp {
				Swap(arr, j, j-gap)
				j -= gap
			}

			arr[j] = jTemp
		}
	}
}
