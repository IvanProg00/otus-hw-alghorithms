package sort

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSelectionSort(t *testing.T) {
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
		{
			val:      []int{8, 43, 43, 64, 102, 345, 346, 346, 494, 592, 648, 781, 782, 983, 1},
			expected: []int{1, 8, 43, 43, 64, 102, 345, 346, 346, 494, 592, 648, 781, 782, 983},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			SelectionSort(tt.val)
			require.EqualValues(t, tt.expected, tt.val)
		})
	}
}

func TestHeapSort(t *testing.T) {
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
		{
			val:      []int{8, 43, 43, 64, 102, 345, 346, 346, 494, 592, 648, 781, 782, 983, 1},
			expected: []int{1, 8, 43, 43, 64, 102, 345, 346, 346, 494, 592, 648, 781, 782, 983},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			HeapSort(tt.val)
			require.EqualValues(t, tt.expected, tt.val)
		})
	}
}