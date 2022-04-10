package array

type Array[T any] interface {
	Size() int
	IsEmpty() bool
	resize()
	Get(index int) T
	Put(item T)
}
