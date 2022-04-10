package array

type VectorArray[T any] struct {
	arr    []T
	size   int
	vector int
}

func NewVectorArray[T any](vector int) Array[T] {
	return &VectorArray[T]{
		vector: vector,
	}
}

func (a *VectorArray[T]) Size() int {
	return a.size
}

func (a *VectorArray[T]) IsEmpty() bool {
	return a.Size() <= 0
}

func (a *VectorArray[T]) resize() {
	arr := make([]T, a.Size()+a.vector)

	for i, v := range a.arr {
		arr[i] = v
	}

	a.arr = arr
}

func (a *VectorArray[T]) Get(index int) T {
	return a.arr[index]
}

func (a *VectorArray[T]) Put(item T) {
	if a.Size() == len(a.arr) {
		a.resize()
	}
	a.arr[a.Size()] = item
	a.size++
}
