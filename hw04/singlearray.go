package array

type SingleArray[T any] struct {
	arr []T
}

func (a *SingleArray[T]) Size() int {
	return len(a.arr)
}

func (a *SingleArray[T]) IsEmpty() bool {
	return a.Size() <= 0
}

func (a *SingleArray[T]) resize() {
	arr := make([]T, a.Size()+1)

	copy(arr, a.arr)

	a.arr = arr
}

func (a *SingleArray[T]) Get(index int) T {
	return a.arr[index]
}

func (a *SingleArray[T]) Put(item T) {
	a.resize()
	a.arr[a.Size()-1] = item
}
