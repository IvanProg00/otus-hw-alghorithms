package array

import "testing"

// BenchmarkSingleArray            202302             72474 ns/op          813246 B/op          1 allocs/op
func BenchmarkSingleArray(b *testing.B) {
	arr := SingleArray[int]{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		arr.Put(i)
	}
}

// BenchmarkVectorArray           1000000             13187 ns/op          200207 B/op          0 allocs/op
func BenchmarkVectorArray(b *testing.B) {
	arr := NewVectorArray[int](20)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		arr.Put(i)
	}
}

// BenchmarkFactorArray          288690402                5.769 ns/op          29 B/op          0 allocs/op
func BenchmarkFactorArray(b *testing.B) {
	arr := NewFactorArray[int]()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		arr.Put(i)
	}
}
