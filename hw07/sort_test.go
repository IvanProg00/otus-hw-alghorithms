package sort

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQuickSortOne(t *testing.T) {
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
			val:      []int{9, -3, 5, 2, 6, 8, -6, 1, 3},
			expected: []int{-6, -3, 1, 2, 3, 5, 6, 8, 9},
		},
		{
			val:      []int{8, 43, 43, 64, 102, 345, 346, 346, 494, 592, 648, 781, 782, 983, 1},
			expected: []int{1, 8, 43, 43, 64, 102, 345, 346, 346, 494, 592, 648, 781, 782, 983},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			QuickSortOne(tt.val)
			require.EqualValues(t, tt.expected, tt.val)
		})
	}
}

func TestQuickSortTwo(t *testing.T) {
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
			val:      []int{9, -3, 5, 2, 6, 8, -6, 1, 3},
			expected: []int{-6, -3, 1, 2, 3, 5, 6, 8, 9},
		},
		{
			val:      []int{8, 43, 43, 64, 102, 345, 346, 346, 494, 592, 648, 781, 782, 983, 1},
			expected: []int{1, 8, 43, 43, 64, 102, 345, 346, 346, 494, 592, 648, 781, 782, 983},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			QuickSortTwo(tt.val)
			require.EqualValues(t, tt.expected, tt.val)
		})
	}
}

func TestMergeSort(t *testing.T) {
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
			val:      []int{9, -3, 5, 2, 6, 8, -6, 1, 3},
			expected: []int{-6, -3, 1, 2, 3, 5, 6, 8, 9},
		},
		{
			val:      []int{8, 43, 43, 64, 102, 345, 346, 346, 494, 592, 648, 781, 782, 983, 1},
			expected: []int{1, 8, 43, 43, 64, 102, 345, 346, 346, 494, 592, 648, 781, 782, 983},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			MergeSort(tt.val)
			require.EqualValues(t, tt.expected, tt.val)
		})
	}
}
