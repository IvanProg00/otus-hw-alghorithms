package tree

import (
	"errors"

	"golang.org/x/exp/constraints"
)

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

var ErrNotFound = errors.New("not found")

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

func (t *Tree[T]) Search(val T) (*Node[T], error) {
	return search(t.root, val)
}

func search[T Number](n *Node[T], val T) (*Node[T], error) {
	if n == nil || len(n.value) == 0 {
		return nil, ErrNotFound
	}

	nVal := n.value[0]

	switch {
	case nVal == val:
		return n, nil
	case val > nVal:
		return search(n.right, val)
	default:
		return search(n.left, val)
	}
}
