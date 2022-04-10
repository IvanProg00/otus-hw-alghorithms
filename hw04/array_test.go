package array

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSize(t *testing.T) {
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
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			require := require.New(t)

			singleArr := SingleArray[int]{tt.arr}
			singleRes := singleArr.Size()
			require.Equal(tt.expected, singleRes)

			vectorArr := NewVectorArray[int](5)
			for _, val := range tt.arr {
				vectorArr.Put(val)
			}
			vectorRes := vectorArr.Size()
			require.Equal(tt.expected, vectorRes)

			factorArr := NewFactorArray[int]()
			for _, val := range tt.arr {
				factorArr.Put(val)
			}
			factorRes := factorArr.Size()
			require.Equal(tt.expected, factorRes)
		})
	}
}

func TestIsEmpty(t *testing.T) {
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
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			require := require.New(t)

			singleArr := SingleArray[int]{tt.arr}
			singleRes := singleArr.IsEmpty()
			require.Equal(tt.expected, singleRes)

			vectorArr := NewVectorArray[int](5)
			for _, val := range tt.arr {
				vectorArr.Put(val)
			}
			vectorRes := vectorArr.IsEmpty()
			require.Equal(tt.expected, vectorRes)

			factorArr := NewFactorArray[int]()
			for _, val := range tt.arr {
				factorArr.Put(val)
			}
			factorRes := factorArr.IsEmpty()
			require.Equal(tt.expected, factorRes)
		})
	}
}

func TestResize(t *testing.T) {
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
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			require := require.New(t)

			singleArr := SingleArray[int]{tt.arr}
			singleArr.resize()
			require.Equal(SingleArray[int]{tt.expected}, singleArr)
		})
	}
}

func TestGet(t *testing.T) {
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
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			require := require.New(t)

			singleArr := SingleArray[int]{tt.arr}
			singleRes := singleArr.Get(tt.index)
			require.Equal(tt.expected, singleRes)

			vectorArr := NewVectorArray[int](5)
			for _, val := range tt.arr {
				vectorArr.Put(val)
			}
			vectorRes := vectorArr.Get(tt.index)
			require.Equal(tt.expected, vectorRes)

			factorArr := NewFactorArray[int]()
			for _, val := range tt.arr {
				factorArr.Put(val)
			}
			factorRes := factorArr.Get(tt.index)
			require.Equal(tt.expected, factorRes)
		})
	}
}

func TestPut(t *testing.T) {
	tests := []struct {
		arr      []int
		val      int
		expected []int
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
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			require := require.New(t)

			singleArr := SingleArray[int]{tt.arr}
			singleArr.Put(tt.val)
			require.EqualValues(tt.expected, singleArr.arr)

			vectorArr := VectorArray[int]{vector: 5}
			for _, val := range tt.arr {
				vectorArr.Put(val)
			}
			vectorArr.Put(tt.val)
			require.Equal(tt.expected, vectorArr.arr[:vectorArr.Size()])

			factorArr := FactorArray[int]{}
			for _, val := range tt.arr {
				factorArr.Put(val)
			}
			factorArr.Put(tt.val)
			require.Equal(tt.expected, factorArr.arr[:factorArr.Size()])
		})
	}
}
