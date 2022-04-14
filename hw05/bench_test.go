package sort

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

var SliseSizes = []int{100, 1_000, 10_000, 100_000}

func randomSlice(size int) []int {
	arr := make([]int, size)

	for i := range arr {
		arr[i] = rand.Int()
	}

	return arr
}

// BenchmarkBubbleSort/slize-size-100-8              127389              9096 ns/op               0 B/op          0 allocs/op
// BenchmarkBubbleSort/slize-size-1000-8               3237            359130 ns/op               0 B/op          0 allocs/op
// BenchmarkBubbleSort/slize-size-10000-8                33          34960885 ns/op               0 B/op          0 allocs/op
// BenchmarkBubbleSort/slize-size-100000-8                1        9741657416 ns/op               0 B/op          0 allocs/op
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

// BenchmarkInsertionSort/slize-size-100-8           264141              4545 ns/op               0 B/op          0 allocs/op
// BenchmarkInsertionSort/slize-size-1000-8            3147            371806 ns/op               0 B/op          0 allocs/op
// BenchmarkInsertionSort/slize-size-10000-8             32          35469930 ns/op               0 B/op          0 allocs/op
// BenchmarkInsertionSort/slize-size-100000-8             1        3655870417 ns/op               0 B/op          0 allocs/op
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

// BenchmarkShellSort/slize-size-100-8               410570              2936 ns/op               0 B/op          0 allocs/op
// BenchmarkShellSort/slize-size-1000-8               22358             53534 ns/op               0 B/op          0 allocs/op
// BenchmarkShellSort/slize-size-10000-8               1666            716755 ns/op               0 B/op          0 allocs/op
// BenchmarkShellSort/slize-size-100000-8               122           9503250 ns/op               0 B/op          0 allocs/op
// BenchmarkShellSort/slize-size-1000000-8                9         125711342 ns/op               0 B/op          0 allocs/op
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
