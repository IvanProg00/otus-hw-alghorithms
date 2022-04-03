package array

import "testing"

// BenchmarkSingleArray   	   93619	    112805 ns/op	  378443 B/op	       1 allocs/op
func BenchmarkSingleArray(b *testing.B) {
	arr := SingleArray[int]{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		arr.Put(i)
	}
}

// BenchmarkVectorArray   	  708400	     41122 ns/op	  141887 B/op	       0 allocs/op
func BenchmarkVectorArray(b *testing.B) {
	arr := NewVectorArray[int](20)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		arr.Put(i)
	}
}
