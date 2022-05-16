package sort

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"
)

var SliseSizes = []int{100, 1_000, 10_000, 100_000}

func randomSlice(size int) ([]int, error) {
	arr := make([]int, size)

	for i := range arr {
		n, err := rand.Int(rand.Reader, big.NewInt(1_000_000))
		if err != nil {
			return []int{}, fmt.Errorf("error get random int: %w", err)
		}

		arr[i] = int(n.Int64())
	}

	return arr, nil
}

func BenchmarkBubbleSort(b *testing.B) {
	b.ResetTimer()

	for _, v := range SliseSizes {
		b.Run(fmt.Sprintf("slize-size-%d", v), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				arr, err := randomSlice(v)
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
				arr, err := randomSlice(v)
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
				arr, err := randomSlice(v)
				require.NoError(b, err)
				b.StartTimer()

				ShellSort(arr)
			}
		})
	}
}
