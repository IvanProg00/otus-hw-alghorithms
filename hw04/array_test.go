package array

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSize(t *testing.T) {
	t.Parallel()

	tests := []struct {
		arr      []int
		expected int
	}{
		{
			arr:      []int{},
			expected: 0,
		},
		{
			arr:      []int{5},
			expected: 1,
		},
		{
			arr:      []int{5, 8, 49},
			expected: 3,
		},
	}

	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			t.Parallel()

			singleArr := SingleArray[int]{tt.arr}
			singleRes := singleArr.Size()
			require.Equal(t, tt.expected, singleRes)

			vectorArr := NewVectorArray[int](5)
			for _, val := range tt.arr {
				vectorArr.Put(val)
			}
			vectorRes := vectorArr.Size()
			require.Equal(t, tt.expected, vectorRes)

			factorArr := NewFactorArray[int]()
			for _, val := range tt.arr {
				factorArr.Put(val)
			}
			factorRes := factorArr.Size()
			require.Equal(t, tt.expected, factorRes)
		})
	}
}

func TestIsEmpty(t *testing.T) {
	t.Parallel()

	tests := []struct {
		arr      []int
		expected bool
	}{
		{
			arr:      []int{},
			expected: true,
		},
		{
			arr:      []int{5},
			expected: false,
		},
		{
			arr:      []int{5, 8, 49},
			expected: false,
		},
	}

	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			t.Parallel()

			singleArr := SingleArray[int]{tt.arr}
			singleRes := singleArr.IsEmpty()
			require.Equal(t, tt.expected, singleRes)

			vectorArr := NewVectorArray[int](5)
			for _, val := range tt.arr {
				vectorArr.Put(val)
			}
			vectorRes := vectorArr.IsEmpty()
			require.Equal(t, tt.expected, vectorRes)

			factorArr := NewFactorArray[int]()
			for _, val := range tt.arr {
				factorArr.Put(val)
			}
			factorRes := factorArr.IsEmpty()
			require.Equal(t, tt.expected, factorRes)
		})
	}
}

func TestResize(t *testing.T) {
	t.Parallel()

	tests := []struct {
		arr      []int
		expected []int
	}{
		{
			arr:      []int{},
			expected: []int{0},
		},
		{
			arr:      []int{5},
			expected: []int{5, 0},
		},
		{
			arr:      []int{5, 8, 49},
			expected: []int{5, 8, 49, 0},
		},
	}

	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			t.Parallel()

			singleArr := SingleArray[int]{tt.arr}
			singleArr.resize()
			require.Equal(t, SingleArray[int]{tt.expected}, singleArr)
		})
	}
}

func TestGet(t *testing.T) {
	t.Parallel()

	tests := []struct {
		arr      []int
		index    int
		expected int
	}{
		{
			arr:      []int{89, 9, 3, 43, 54, 94},
			index:    4,
			expected: 54,
		},
		{
			arr:      []int{5},
			index:    0,
			expected: 5,
		},
		{
			arr:      []int{5, 8, 49},
			index:    2,
			expected: 49,
		},
	}

	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			t.Parallel()

			singleArr := SingleArray[int]{tt.arr}
			singleRes := singleArr.Get(tt.index)
			require.Equal(t, tt.expected, singleRes)

			vectorArr := NewVectorArray[int](5)
			for _, val := range tt.arr {
				vectorArr.Put(val)
			}
			vectorRes := vectorArr.Get(tt.index)
			require.Equal(t, tt.expected, vectorRes)

			factorArr := NewFactorArray[int]()
			for _, val := range tt.arr {
				factorArr.Put(val)
			}
			factorRes := factorArr.Get(tt.index)
			require.Equal(t, tt.expected, factorRes)
		})
	}
}

func TestPut(t *testing.T) {
	t.Parallel()

	tests := []struct {
		arr      []int
		expected []int
		val      int
	}{
		{
			arr:      []int{},
			val:      79,
			expected: []int{79},
		},
		{
			arr:      []int{89, 9, 3, 43, 54, 94},
			val:      4,
			expected: []int{89, 9, 3, 43, 54, 94, 4},
		},
		{
			arr:      []int{5},
			val:      4839,
			expected: []int{5, 4839},
		},
		{
			arr:      []int{5, 8, 49},
			val:      94,
			expected: []int{5, 8, 49, 94},
		},
	}

	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			t.Parallel()

			singleArr := SingleArray[int]{tt.arr}
			singleArr.Put(tt.val)
			require.EqualValues(t, tt.expected, singleArr.arr)

			vectorArr := VectorArray[int]{vector: 5}
			for _, val := range tt.arr {
				vectorArr.Put(val)
			}
			vectorArr.Put(tt.val)
			require.Equal(t, tt.expected, vectorArr.arr[:vectorArr.Size()])

			factorArr := FactorArray[int]{}
			for _, val := range tt.arr {
				factorArr.Put(val)
			}
			factorArr.Put(tt.val)
			require.Equal(t, tt.expected, factorArr.arr[:factorArr.Size()])
		})
	}
}
