package array

type T interface {
	any
}

type Array struct {
	arr []T
}

func (a *Array) Size() int {
	return 0
}

func (a *Array) IsEmpty() bool {
	return false
}

func (a *Array) resize() {
}

func (a *Array) Get(index int) T {
	return a.arr[index]
}

func (a *Array) Put(item T) {
}
