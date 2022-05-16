package sort

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/IvanProg00/otus-hw-algorithms/utils"
	"github.com/stretchr/testify/require"
)

var SliseSizes = []int{100, 1_000, 10_000, 100_000, 1_000_000}

func BenchmarkQuickSortOne(b *testing.B) {
	rand.Seed(time.Now().Unix())
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

				QuickSortOne(arr)
			}
		})
	}
}

func BenchmarkQuickSortTwo(b *testing.B) {
	rand.Seed(time.Now().Unix())
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

				QuickSortTwo(arr)
			}
		})
	}
}

func BenchmarkMergeSort(b *testing.B) {
	rand.Seed(time.Now().Unix())
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

				MergeSort(arr)
			}
		})
	}
}
