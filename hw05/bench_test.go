package sort

import (
	"fmt"
	"testing"

	"github.com/IvanProg00/otus-hw-algorithms/utils"
	"github.com/stretchr/testify/require"
)

var SliseSizes = []int{100, 1_000, 10_000, 100_000}

func BenchmarkBubbleSort(b *testing.B) {
	b.ResetTimer()

	for _, v := range SliseSizes {
		b.Run(fmt.Sprintf("slize-size-%d", v), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				arr, err := utils.RandomSlice(v)
				if err != nil {
					require.NoError(b, err)
				}
				b.StartTimer()

				BubbleSort(arr)
			}
		})
	}
}

func BenchmarkInsertionSort(b *testing.B) {
	b.ResetTimer()

	for _, v := range SliseSizes {
		b.Run(fmt.Sprintf("slize-size-%d", v), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				arr, err := utils.RandomSlice(v)
				require.NoError(b, err)
				b.StartTimer()

				InsertionSort(arr)
			}
		})
	}
}

func BenchmarkShellSort(b *testing.B) {
	b.ResetTimer()

	for _, v := range SliseSizes {
		b.Run(fmt.Sprintf("slize-size-%d", v), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				arr, err := utils.RandomSlice(v)
				require.NoError(b, err)
				b.StartTimer()

				ShellSort(arr)
			}
		})
	}
}
