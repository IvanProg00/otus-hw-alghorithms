package tree

import (
	"errors"
	"fmt"

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
	value       T
}

var ErrNotFound = errors.New("not found")

func NewTree[T Number]() *Tree[T] {
	return &Tree[T]{}
}

func (t *Tree[T]) Insert(val T) {
	t.root = insert(t.root, val)
}

func insert[T Number](curr *Node[T], val T) *Node[T] {
	if curr != nil {
		switch {
		case val == curr.value:
			curr.value = val
		case val > curr.value:
			curr.right = insert(curr.right, val)
		default:
			curr.left = insert(curr.left, val)
		}

		return curr
	}

	return &Node[T]{
		value: val,
	}
}

func (t *Tree[T]) Search(val T) (*Node[T], error) {
	return search(t.root, val)
}

func search[T Number](n *Node[T], val T) (*Node[T], error) {
	if n == nil {
		return nil, ErrNotFound
	}

	switch {
	case n.value == val:
		return n, nil
	case val > n.value:
		return search(n.right, val)
	default:
		return search(n.left, val)
	}
}

func (t *Tree[T]) Remove(val T) error {
	n, err := remove(t.root, val)
	if err != nil {
		return fmt.Errorf("remove: %w", err)
	}

	t.root = n

	return nil
}

func remove[T Number](n *Node[T], val T) (*Node[T], error) {
	if n == nil {
		return n, ErrNotFound
	}

	switch {
	case val > n.value:
		node, err := remove(n.right, val)
		if err != nil {
			return nil, err
		}

		n.right = node
	case val < n.value:
		node, err := remove(n.left, val)
		if err != nil {
			return nil, err
		}

		n.left = node
	default:
		if n.left == nil {
			temp := n.right
			n = nil

			return temp, nil
		}

		if n.right == nil {
			temp := n.left
			n = nil

			return temp, nil
		}

		temp := minValueNode(n.right)
		n.value = temp.value

		rightNode, err := remove(n.right, temp.value)
		if err != nil {
			return nil, err
		}

		n.right = rightNode
	}

	return n, nil
}

func minValueNode[T Number](n *Node[T]) *Node[T] {
	for n.left != nil {
		n = n.left
	}

	return n
}
