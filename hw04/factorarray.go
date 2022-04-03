package array

const factor = 2

type FactorArray[T any] struct {
	arr  []T
	size int
}

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

	for i, v := range a.arr {
		arr[i] = v
	}

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
