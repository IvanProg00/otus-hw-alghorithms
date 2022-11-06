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
			require.EqualValues(t, tt.expTree, *tt.tree)
		})
	}
}
