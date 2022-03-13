package primes

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPrimes1(t *testing.T) {
	tests := []struct {
		n        int
		expected int
	}{
		{10, 4},        // 0.00s
		{1, 0},         // 0.00s
		{2, 1},         // 0.00s
		{3, 2},         // 0.00s
		{4, 2},         // 0.00s
		{5, 3},         // 0.00s
		{100, 25},      // 0.00s
		{1000, 168},    // 0.00s
		{10000, 1229},  // 0.06s
		{100000, 9592}, // 3.59s
		// {1000000, 78498},
		// {10000000, 664579},
		// {100000000, 5761455},
		// {1000000000, 50847534},
		// {123456789, 7027260},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			require := require.New(t)

			require.Equal(tt.expected, Primes1(tt.n))
		})
	}
}
