package sort

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func randomSlice(size int) []int {
	arr := make([]int, size)

	for i := range arr {
		arr[i] = rand.Int()
	}

	return arr
}

var SliseSizes = []int{100, 1_000, 10_000, 100_000, 1_000_000}

// BenchmarkQuickSortOne/slize-size-100-8            361191              3117 ns/op               0 B/op          0 allocs/op
// BenchmarkQuickSortOne/slize-size-1000-8            28899             41383 ns/op               0 B/op          0 allocs/op
// BenchmarkQuickSortOne/slize-size-10000-8            2438            488720 ns/op               0 B/op          0 allocs/op
// BenchmarkQuickSortOne/slize-size-100000-8            206           5885733 ns/op               0 B/op          0 allocs/op
// BenchmarkQuickSortOne/slize-size-1000000-8            16          68065091 ns/op               0 B/op          0 allocs/op
func BenchmarkQuickSortOne(b *testing.B) {
	rand.Seed(time.Now().Unix())
	b.ResetTimer()

	for _, v := range SliseSizes {
		b.Run(fmt.Sprintf("slize-size-%d", v), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				arr := randomSlice(v)
				b.StartTimer()

				QuickSortOne(arr)
			}
		})
	}
}

// BenchmarkQuickSortTwo/slize-size-100-8            483141              2485 ns/op               0 B/op          0 allocs/op
// BenchmarkQuickSortTwo/slize-size-1000-8            34228             35058 ns/op               0 B/op          0 allocs/op
// BenchmarkQuickSortTwo/slize-size-10000-8            2718            433702 ns/op               0 B/op          0 allocs/op
// BenchmarkQuickSortTwo/slize-size-100000-8            223           5357234 ns/op               0 B/op          0 allocs/op
// BenchmarkQuickSortTwo/slize-size-1000000-8            18          63624998 ns/op               0 B/op          0 allocs/op
func BenchmarkQuickSortTwo(b *testing.B) {
	rand.Seed(time.Now().Unix())
	b.ResetTimer()

	for _, v := range SliseSizes {
		b.Run(fmt.Sprintf("slize-size-%d", v), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				arr := randomSlice(v)
				b.StartTimer()

				QuickSortTwo(arr)
			}
		})
	}
}

// BenchmarkMergeSort/slize-size-100-8               172479              7008 ns/op            5504 B/op        198 allocs/op
// BenchmarkMergeSort/slize-size-1000-8               15373             77514 ns/op           81152 B/op       1998 allocs/op
// BenchmarkMergeSort/slize-size-10000-8               1309            907311 ns/op         1110790 B/op      19998 allocs/op
// BenchmarkMergeSort/slize-size-100000-8               100          11168807 ns/op        14055027 B/op     199998 allocs/op
// BenchmarkMergeSort/slize-size-1000000-8                9         121635495 ns/op        164375989 B/op   2000001 allocs/op
func BenchmarkMergeSort(b *testing.B) {
	rand.Seed(time.Now().Unix())
	b.ResetTimer()

	for _, v := range SliseSizes {
		b.Run(fmt.Sprintf("slize-size-%d", v), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				arr := randomSlice(v)
				b.StartTimer()

				MergeSort(arr)
			}
		})
	}
}
