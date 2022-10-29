package array

type FactorArray[T any] struct {
	arr  []T
	size int
}

const factor = 2

func NewFactorArray[T any]() Array[T] {
	return &FactorArray[T]{}
}

func (a *FactorArray[T]) Size() int {
	return a.size
}

func (a *FactorArray[T]) IsEmpty() bool {
	return a.Size() <= 0
}

func (a *FactorArray[T]) resize() {
	arr := make([]T, a.Size()*factor+1)

	copy(arr, a.arr)

	a.arr = arr
}

func (a *FactorArray[T]) Get(index int) T {
	return a.arr[index]
}

func (a *FactorArray[T]) Put(item T) {
	if a.Size() == len(a.arr) {
		a.resize()
	}

	a.arr[a.Size()] = item
	a.size++
}
