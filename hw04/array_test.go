package array

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSize(t *testing.T) {
	tests := []struct {
		arr      Array[int]
		expected int
	}{
		{
			arr:      &SingleArray[int]{[]int{}},
			expected: 0,
		},
		{
			arr:      &SingleArray[int]{[]int{5}},
			expected: 1,
		},
		{
			arr:      &SingleArray[int]{[]int{5, 8, 49}},
			expected: 3,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			res := tt.arr.Size()

			require.Equal(t, tt.expected, res)
		})
	}
}

func TestIsEmpty(t *testing.T) {
	tests := []struct {
		arr      Array[int]
		expected bool
	}{
		{
			arr:      &SingleArray[int]{[]int{}},
			expected: true,
		},
		{
			arr:      &SingleArray[int]{[]int{5}},
			expected: false,
		},
		{
			arr:      &SingleArray[int]{[]int{5, 8, 49}},
			expected: false,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			res := tt.arr.IsEmpty()

			require.Equal(t, tt.expected, res)
		})
	}
}

func TestResize(t *testing.T) {
	tests := []struct {
		arr      Array[int]
		expected Array[int]
	}{
		{
			arr:      &SingleArray[int]{[]int{}},
			expected: &SingleArray[int]{[]int{0}},
		},
		{
			arr:      &SingleArray[int]{[]int{5}},
			expected: &SingleArray[int]{[]int{5, 0}},
		},
		{
			arr:      &SingleArray[int]{[]int{5, 8, 49}},
			expected: &SingleArray[int]{[]int{5, 8, 49, 0}},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			tt.arr.resize()

			require.Equal(t, tt.expected, tt.arr)
		})
	}
}

func TestGet(t *testing.T) {
	tests := []struct {
		arr      Array[int]
		index    int
		expected int
	}{
		{
			arr:      &SingleArray[int]{[]int{89, 9, 3, 43, 54, 94}},
			index:    4,
			expected: 54,
		},
		{
			arr:      &SingleArray[int]{[]int{5}},
			index:    0,
			expected: 5,
		},
		{
			arr:      &SingleArray[int]{[]int{5, 8, 49}},
			index:    2,
			expected: 49,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			res := tt.arr.Get(tt.index)

			require.Equal(t, tt.expected, res)
		})
	}
}

func TestPut(t *testing.T) {
	tests := []struct {
		arr      SingleArray[int]
		val      int
		expected SingleArray[int]
	}{
		{
			arr:      SingleArray[int]{[]int{89, 9, 3, 43, 54, 94}},
			val:      4,
			expected: SingleArray[int]{[]int{89, 9, 3, 43, 54, 94, 4}},
		},
		{
			arr:      SingleArray[int]{[]int{5}},
			val:      4839,
			expected: SingleArray[int]{[]int{5, 4839}},
		},
		{
			arr:      SingleArray[int]{[]int{5, 8, 49}},
			val:      94,
			expected: SingleArray[int]{[]int{5, 8, 49, 94}},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			tt.arr.Put(tt.val)

			require.Equal(t, tt.expected, tt.arr)
		})
	}
}
