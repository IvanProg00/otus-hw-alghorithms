package tree

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInsert(t *testing.T) {
	t.Parallel()

	tests := []struct {
		val     int
		tree    *Tree[int]
		expTree Tree[int]
	}{
		{
			val:  5,
			tree: NewTree[int](),
			expTree: Tree[int]{
				root: &Node[int]{
					value: 5,
				},
			},
		},
		{
			val:  8,
			tree: &Tree[int]{},
			expTree: Tree[int]{
				root: &Node[int]{
					value: 8,
				},
			},
		},
		{
			val: 13,
			tree: &Tree[int]{
				root: &Node[int]{
					value: 13,
				},
			},
			expTree: Tree[int]{
				root: &Node[int]{
					value: 13,
				},
			},
		},
		{
			val: -8,
			tree: &Tree[int]{
				root: &Node[int]{
					value: 1,
				},
			},
			expTree: Tree[int]{
				root: &Node[int]{
					value: 1,
					left: &Node[int]{
						value: -8,
					},
				},
			},
		},
		{
			val: 58,
			tree: &Tree[int]{
				root: &Node[int]{
					value: 44,
					left:  &Node[int]{value: 36},
				},
			},
			expTree: Tree[int]{
				root: &Node[int]{
					value: 44,
					left:  &Node[int]{value: 36},
					right: &Node[int]{value: 58},
				},
			},
		},
		{
			val: 13,
			tree: &Tree[int]{
				root: &Node[int]{
					value: 18,
					left:  &Node[int]{value: 12},
				},
			},
			expTree: Tree[int]{
				root: &Node[int]{
					value: 18,
					left: &Node[int]{
						value: 12,
						right: &Node[int]{
							value: 13,
						},
					},
				},
			},
		},
		{
			val: 36,
			tree: &Tree[int]{
				root: &Node[int]{
					value: 4,
					left:  &Node[int]{value: 2},
					right: &Node[int]{
						value: 18,
						right: &Node[int]{
							value: 28,
							right: &Node[int]{
								value: 45,
								left:  &Node[int]{value: 34},
							},
						},
					},
				},
			},
			expTree: Tree[int]{
				root: &Node[int]{
					value: 4,
					left:  &Node[int]{value: 2},
					right: &Node[int]{
						value: 18,
						right: &Node[int]{
							value: 28,
							right: &Node[int]{
								value: 45,
								left: &Node[int]{
									value: 34,
									right: &Node[int]{
										value: 36,
									},
								},
							},
						},
					},
				},
			},
		},
	}

	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			t.Parallel()

			tt.tree.Insert(tt.val)
			require.NotNil(t, tt.tree)
			require.EqualValues(t, tt.expTree, *tt.tree)
		})
	}
}

func TestSearch(t *testing.T) {
	t.Parallel()

	tests := []struct {
		val  int
		tree *Tree[int]
		exp  Node[int]
	}{
		{
			val: 5,
			tree: &Tree[int]{
				root: &Node[int]{
					value: 5,
				},
			},
			exp: Node[int]{
				value: 5,
			},
		},
		{
			val: -43,
			tree: &Tree[int]{
				root: &Node[int]{
					value: -15,
					left: &Node[int]{
						value: -43,
					},
				},
			},
			exp: Node[int]{
				value: -43,
			},
		},
		{
			val: 68,
			tree: &Tree[int]{
				root: &Node[int]{
					value: 44,
					left:  &Node[int]{value: 21},
					right: &Node[int]{value: 68},
				},
			},
			exp: Node[int]{value: 68},
		},
		{
			val: 18,
			tree: &Tree[int]{
				root: &Node[int]{
					value: 18,
					left:  &Node[int]{value: 12},
					right: &Node[int]{value: 13},
				},
			},
			exp: Node[int]{
				value: 18,
				left:  &Node[int]{value: 12},
				right: &Node[int]{value: 13},
			},
		},
		{
			val: 34,
			tree: &Tree[int]{
				root: &Node[int]{
					value: 4,
					left:  &Node[int]{value: 2},
					right: &Node[int]{
						value: 18,
						right: &Node[int]{
							value: 28,
							right: &Node[int]{
								value: 45,
								left: &Node[int]{
									value: 34,
									right: &Node[int]{value: 36},
								},
							},
						},
					},
				},
			},
			exp: Node[int]{
				value: 34,
				right: &Node[int]{value: 36},
			},
		},
	}

	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			t.Parallel()

			res, err := tt.tree.Search(tt.val)
			require.NoError(t, err)
			require.NotNil(t, res)
			require.EqualValues(t, tt.exp, *res)
		})
	}
}

func TestSearch_error(t *testing.T) {
	t.Parallel()

	tests := []struct {
		val  int
		tree *Tree[int]
	}{
		{
			val:  5,
			tree: NewTree[int](),
		},
		{
			val:  8,
			tree: &Tree[int]{},
		},
		{
			val: 3,
			tree: &Tree[int]{
				root: &Node[int]{
					value: 14,
					left: &Node[int]{
						value: -5,
						right: &Node[int]{
							value: 8,
							left: &Node[int]{
								value: 4,
								left: &Node[int]{
									value: 0,
									right: &Node[int]{value: 2},
								},
							},
						},
					},
					right: &Node[int]{value: 18},
				},
			},
		},
	}

	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			t.Parallel()

			res, err := tt.tree.Search(tt.val)
			require.ErrorIs(t, err, ErrNotFound)
			require.Nil(t, res)
		})
	}
}

func TestRemove(t *testing.T) {
	t.Parallel()

	tests := []struct {
		val     int
		tree    *Tree[int]
		expTree Tree[int]
	}{
		{
			val: 79,
			tree: &Tree[int]{
				root: &Node[int]{
					value: 79,
				},
			},
			expTree: Tree[int]{},
		},
		{
			val: 94,
			tree: &Tree[int]{
				root: &Node[int]{
					value: 45,
					right: &Node[int]{
						value: 94,
					},
				},
			},
			expTree: Tree[int]{
				root: &Node[int]{
					value: 45,
				},
			},
		},
		{
			val: -78,
			tree: &Tree[int]{
				root: &Node[int]{
					value: 4,
					left: &Node[int]{
						value: -78,
					},
				},
			},
			expTree: Tree[int]{
				root: &Node[int]{
					value: 4,
				},
			},
		},
		{
			val: 36,
			tree: &Tree[int]{
				root: &Node[int]{
					value: 1,
					left:  &Node[int]{value: -89},
					right: &Node[int]{value: 36},
				},
			},
			expTree: Tree[int]{
				root: &Node[int]{
					value: 1,
					left:  &Node[int]{value: -89},
				},
			},
		},
		{
			val: -43,
			tree: &Tree[int]{
				root: &Node[int]{
					value: 54,
					left: &Node[int]{
						value: -43,
						right: &Node[int]{
							value: 0,
						},
					},
				},
			},
			expTree: Tree[int]{
				root: &Node[int]{
					value: 54,
					left:  &Node[int]{value: 0},
				},
			},
		},
		{
			val: 23,
			tree: &Tree[int]{
				root: &Node[int]{
					value: 94,
					left: &Node[int]{
						value: 23,
						left: &Node[int]{
							value: 5,
						},
					},
				},
			},
			expTree: Tree[int]{
				root: &Node[int]{
					value: 94,
					left:  &Node[int]{value: 5},
				},
			},
		},
		{
			val: 50,
			tree: &Tree[int]{
				root: &Node[int]{
					value: 50,
					left:  &Node[int]{value: 40},
					right: &Node[int]{
						value: 70,
						left:  &Node[int]{value: 60},
						right: &Node[int]{value: 80},
					},
				},
			},
			expTree: Tree[int]{
				root: &Node[int]{
					value: 60,
					left:  &Node[int]{value: 40},
					right: &Node[int]{
						value: 70,
						right: &Node[int]{value: 80},
					},
				},
			},
		},
		{
			val: 50,
			tree: &Tree[int]{
				root: &Node[int]{
					value: 50,
					left:  &Node[int]{value: 40},
					right: &Node[int]{
						value: 70,
						left:  &Node[int]{value: 60},
						right: &Node[int]{value: 80},
					},
				},
			},
			expTree: Tree[int]{
				root: &Node[int]{
					value: 60,
					left:  &Node[int]{value: 40},
					right: &Node[int]{
						value: 70,
						right: &Node[int]{value: 80},
					},
				},
			},
		},
	}

	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			t.Parallel()

			err := tt.tree.Remove(tt.val)
			require.NoError(t, err)
			require.NotNil(t, tt.tree)
			require.EqualValues(t, tt.expTree, *tt.tree)
		})
	}
}

func TestRemove_error(t *testing.T) {
	t.Parallel()

	tests := []struct {
		val    int
		tree   *Tree[int]
		expErr error
	}{
		{
			val: 27,
			tree: &Tree[int]{
				root: &Node[int]{
					value: 29,
				},
			},
			expErr: ErrNotFound,
		},
		{
			val: 49,
			tree: &Tree[int]{
				root: &Node[int]{
					value: 24,
					right: &Node[int]{
						value: 50,
					},
				},
			},
			expErr: ErrNotFound,
		},
		{
			val: 39,
			tree: &Tree[int]{
				root: &Node[int]{
					value: 105,
					left: &Node[int]{
						value: 31,
					},
				},
			},
			expErr: ErrNotFound,
		},
		{
			val: 11,
			tree: &Tree[int]{
				root: &Node[int]{
					value: 10,
					left:  &Node[int]{value: 11},
				},
			},
			expErr: ErrNotFound,
		},
	}

	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			t.Parallel()

			err := tt.tree.Remove(tt.val)
			require.ErrorIs(t, err, tt.expErr)
		})
	}
}
