package sort

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

var SliseSizes = []int{100, 1_000, 10_000}

func randomSlice(size int) []int {
	arr := make([]int, size)

	for i := range arr {
		arr[i] = rand.Int()
	}

	return arr
}

// BenchmarkBubbleSort/slize-size-100-8              123969              9084 ns/op               0 B/op          0 allocs/op
// BenchmarkBubbleSort/slize-size-1000-8               3310            360824 ns/op               0 B/op          0 allocs/op
// BenchmarkBubbleSort/slize-size-10000-8                33          34879076 ns/op               0 B/op          0 allocs/op
func BenchmarkBubbleSort(b *testing.B) {
	rand.Seed(time.Now().Unix())
	b.ResetTimer()

	for _, v := range SliseSizes {
		b.Run(fmt.Sprintf("slize-size-%d", v), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				arr := randomSlice(v)
				b.StartTimer()

				BubbleSort(arr)
			}
		})
	}
}

// BenchmarkInsertionSort/slize-size-100-8           263697              4553 ns/op               0 B/op          0 allocs/op
// BenchmarkInsertionSort/slize-size-1000-8            3189            373437 ns/op               0 B/op          0 allocs/op
// BenchmarkInsertionSort/slize-size-10000-8             33          35351353 ns/op               0 B/op          0 allocs/op
func BenchmarkInsertionSort(b *testing.B) {
	rand.Seed(time.Now().Unix())
	b.ResetTimer()

	for _, v := range SliseSizes {
		b.Run(fmt.Sprintf("slize-size-%d", v), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				arr := randomSlice(v)
				b.StartTimer()

				InsertionSort(arr)
			}
		})
	}
}

// BenchmarkShellSort/slize-size-100-8               412507              2903 ns/op               0 B/op          0 allocs/op
// BenchmarkShellSort/slize-size-1000-8               22502             53263 ns/op               0 B/op          0 allocs/op
// BenchmarkShellSort/slize-size-10000-8               1672            718508 ns/op               0 B/op          0 allocs/op
func BenchmarkShellSort(b *testing.B) {
	rand.Seed(time.Now().Unix())
	b.ResetTimer()

	for _, v := range SliseSizes {
		b.Run(fmt.Sprintf("slize-size-%d", v), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				arr := randomSlice(v)
				b.StartTimer()

				ShellSort(arr)
			}
		})
	}
}
