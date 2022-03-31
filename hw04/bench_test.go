package array

import "testing"

func BenchmarkArray(b *testing.B) {
	arr := SingleArray[int]{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		arr.Put(i)
	}
}
