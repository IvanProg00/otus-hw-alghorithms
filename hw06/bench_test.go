package sort

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

var SliseSizes = []int{100, 1_000, 10_000, 100_000, 1_000_000}

func randomSlice(size int) []int {
	arr := make([]int, size)

	for i := range arr {
		arr[i] = rand.Int()
	}

	return arr
}

// BenchmarkSelectionSort/slize-size-100-8                   180285              6242 ns/op               0 B/op          0 allocs/op
// BenchmarkSelectionSort/slize-size-1000-8                    3187            368832 ns/op               0 B/op          0 allocs/op
// BenchmarkSelectionSort/slize-size-10000-8                     36          31995226 ns/op               0 B/op          0 allocs/op
// BenchmarkSelectionSort/slize-size-100000-8                     1        3162073917 ns/op               0 B/op          0 allocs/op
func BenchmarkSelectionSort(b *testing.B) {
	rand.Seed(time.Now().Unix())
	b.ResetTimer()

	for _, v := range SliseSizes {
		b.Run(fmt.Sprintf("slize-size-%d", v), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				arr := randomSlice(v)
				b.StartTimer()

				SelectionSort(arr)
			}
		})
	}
}

// BenchmarkHeapSort/slize-size-100-8                        303670              3679 ns/op               0 B/op          0 allocs/op
// BenchmarkHeapSort/slize-size-1000-8                        21429             55958 ns/op               0 B/op          0 allocs/op
// BenchmarkHeapSort/slize-size-10000-8                        1706            695655 ns/op               0 B/op          0 allocs/op
// BenchmarkHeapSort/slize-size-100000-8                        133           8954361 ns/op               0 B/op          0 allocs/op
// BenchmarkHeapSort/slize-size-1000000-8                         9         113165870 ns/op               0 B/op          0 allocs/op
func BenchmarkHeapSort(b *testing.B) {
	rand.Seed(time.Now().Unix())
	b.ResetTimer()

	for _, v := range SliseSizes {
		b.Run(fmt.Sprintf("slize-size-%d", v), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				arr := randomSlice(v)
				b.StartTimer()

				HeapSort(arr)
			}
		})
	}
}
