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
					value: []int{5},
				},
			},
		},
		{
			val:  8,
			tree: &Tree[int]{},
			expTree: Tree[int]{
				root: &Node[int]{
					value: []int{8},
				},
			},
		},
		{
			val: 13,
			tree: &Tree[int]{
				root: &Node[int]{
					value: []int{13},
				},
			},
			expTree: Tree[int]{
				root: &Node[int]{
					value: []int{13, 13},
				},
			},
		},
		{
			val: -8,
			tree: &Tree[int]{
				root: &Node[int]{
					value: []int{1},
				},
			},
			expTree: Tree[int]{
				root: &Node[int]{
					value: []int{1},
					left: &Node[int]{
						value: []int{-8},
					},
				},
			},
		},
		{
			val: 58,
			tree: &Tree[int]{
				root: &Node[int]{
					value: []int{44},
					left:  &Node[int]{value: []int{36}},
				},
			},
			expTree: Tree[int]{
				root: &Node[int]{
					value: []int{44},
					left:  &Node[int]{value: []int{36}},
					right: &Node[int]{value: []int{58}},
				},
			},
		},
		{
			val: 13,
			tree: &Tree[int]{
				root: &Node[int]{
					value: []int{18},
					left:  &Node[int]{value: []int{12}},
				},
			},
			expTree: Tree[int]{
				root: &Node[int]{
					value: []int{18},
					left: &Node[int]{
						value: []int{12},
						right: &Node[int]{
							value: []int{13},
						},
					},
				},
			},
		},
		{
			val: 36,
			tree: &Tree[int]{
				root: &Node[int]{
					value: []int{4, 4},
					left:  &Node[int]{value: []int{2}},
					right: &Node[int]{
						value: []int{18, 18},
						right: &Node[int]{
							value: []int{28},
							right: &Node[int]{
								value: []int{45},
								left:  &Node[int]{value: []int{34}},
							},
						},
					},
				},
			},
			expTree: Tree[int]{
				root: &Node[int]{
					value: []int{4, 4},
					left:  &Node[int]{value: []int{2}},
					right: &Node[int]{
						value: []int{18, 18},
						right: &Node[int]{
							value: []int{28},
							right: &Node[int]{
								value: []int{45},
								left: &Node[int]{
									value: []int{34},
									right: &Node[int]{
										value: []int{36},
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
		exp  []int
	}{
		{
			val: 5,
			tree: &Tree[int]{
				root: &Node[int]{
					value: []int{5},
				},
			},
			exp: []int{5},
		},
		{
			val: -43,
			tree: &Tree[int]{
				root: &Node[int]{
					value: []int{-15},
					left: &Node[int]{
						value: []int{-43},
					},
				},
			},
			exp: []int{-43},
		},
		{
			val: 68,
			tree: &Tree[int]{
				root: &Node[int]{
					value: []int{44},
					left:  &Node[int]{value: []int{21}},
					right: &Node[int]{value: []int{68}},
				},
			},
			exp: []int{68},
		},
		{
			val: 18,
			tree: &Tree[int]{
				root: &Node[int]{
					value: []int{18},
					left:  &Node[int]{value: []int{12}},
					right: &Node[int]{value: []int{13}},
				},
			},
			exp: []int{18},
		},
		{
			val: 34,
			tree: &Tree[int]{
				root: &Node[int]{
					value: []int{4, 4},
					left:  &Node[int]{value: []int{2}},
					right: &Node[int]{
						value: []int{18, 18},
						right: &Node[int]{
							value: []int{28},
							right: &Node[int]{
								value: []int{45},
								left: &Node[int]{
									value: []int{34},
									right: &Node[int]{value: []int{36}},
								},
							},
						},
					},
				},
			},
			exp: []int{34},
		},
	}

	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			t.Parallel()

			res, err := tt.tree.Search(tt.val)
			require.NoError(t, err)
			require.NotNil(t, res)
			require.EqualValues(t, tt.exp, res)
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
					value: []int{14},
					left: &Node[int]{
						value: []int{-5},
						right: &Node[int]{
							value: []int{8},
							left: &Node[int]{
								value: []int{4},
								left: &Node[int]{
									value: []int{0},
									right: &Node[int]{value: []int{2}},
								},
							},
						},
					},
					right: &Node[int]{value: []int{18}},
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
