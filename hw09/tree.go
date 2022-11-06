package tree

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Float | constraints.Integer
}

type Tree[T Number] struct {
	root *Node[T]
}

type Node[T Number] struct {
	left, right *Node[T]
	value       []T
}

func NewTree[T Number]() *Tree[T] {
	return &Tree[T]{}
}

func (t *Tree[T]) Insert(val T) {
	t.root = insert(t.root, val)
}

func insert[T Number](curr *Node[T], val T) *Node[T] {
	if curr != nil && len(curr.value) != 0 {
		currVal := curr.value[0]

		switch {
		case val == currVal:
			curr.value = append(curr.value, val)
		case val > currVal:
			curr.right = insert(curr.right, val)
		default:
			curr.left = insert(curr.left, val)
		}

		return curr
	}

	return &Node[T]{
		value: []T{val},
	}
}
