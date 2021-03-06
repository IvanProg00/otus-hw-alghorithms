package sort

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBubbleSort(t *testing.T) {
	t.Parallel()

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
		i, tt := i, tt
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			t.Parallel()

			BubbleSort(tt.val)
			require.EqualValues(t, tt.expected, tt.val)
		})
	}
}

func TestInsertionSort(t *testing.T) {
	t.Parallel()

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
		i, tt := i, tt
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			t.Parallel()

			InsertionSort(tt.val)
			require.EqualValues(t, tt.expected, tt.val)
		})
	}
}

func TestShellSort(t *testing.T) {
	t.Parallel()

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
		i, tt := i, tt
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			t.Parallel()

			ShellSort(tt.val)
			require.EqualValues(t, tt.expected, tt.val)
		})
	}
}
