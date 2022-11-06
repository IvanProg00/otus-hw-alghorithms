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
	if t.root == nil {
		t.root = &Node[T]{
			value: []T{val},
		}

		return
	}

	t.root.insert(val)
}

func (n *Node[T]) insert(val T) {
	if len(n.value) == 0 || n.value[0] == val {
		n.value = append(n.value, val)
		return
	}

	nodeVal := n.value[0]

	if val > nodeVal {
		n.right = setNodeValue(n.right, val)
	} else {
		n.left = setNodeValue(n.left, val)
	}
}

func setNodeValue[T Number](curr *Node[T], val T) *Node[T] {
	if curr != nil && len(curr.value) != 0 {
		currVal := curr.value[0]

		switch {
		case val == currVal:
			curr.value = append(curr.value, val)
		case val > currVal:
			curr.right = setNodeValue(curr.right, val)
		default:
			curr.left = setNodeValue(curr.left, val)
		}

		return curr
	}

	return &Node[T]{
		value: []T{val},
	}
}
