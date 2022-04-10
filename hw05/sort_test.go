package sort

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBubbleSort(t *testing.T) {
	tests := []struct {
		val      []int
		expected []int
	}{
		{
			val:      []int{},
			expected: []int{},
		},
		{
			val:      []int{5, 0},
			expected: []int{0, 5},
		},
		{
			val:      []int{1, 2, 4},
			expected: []int{1, 2, 4},
		},
		{
			val:      []int{7, 3, 8},
			expected: []int{3, 7, 8},
		},
		{
			val:      []int{8, 43, 984, 1, 43, 64},
			expected: []int{1, 8, 43, 43, 64, 984},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			BubbleSort(tt.val)
			require.EqualValues(t, tt.expected, tt.val)
		})
	}
}

func TestInsertionSort(t *testing.T) {
	tests := []struct {
		val      []int
		expected []int
	}{
		{
			val:      []int{},
			expected: []int{},
		},
		{
			val:      []int{5, 0},
			expected: []int{0, 5},
		},
		{
			val:      []int{1, 2, 4},
			expected: []int{1, 2, 4},
		},
		{
			val:      []int{7, 3, 8},
			expected: []int{3, 7, 8},
		},
		{
			val:      []int{8, 43, 984, 1, 43, 64},
			expected: []int{1, 8, 43, 43, 64, 984},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			InsertionSort(tt.val)
			require.EqualValues(t, tt.expected, tt.val)
		})
	}
}
