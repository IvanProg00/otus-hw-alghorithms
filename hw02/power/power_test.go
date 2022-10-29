package power

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIter(t *testing.T) {
	t.Parallel()

	const roundUntil = 100000000000

	tests := []struct {
		x        float64
		y        int
		expected float64
	}{
		{2, 10, 1024},                              // 0.00s
		{123456789, 0, 1},                          // 0.00s
		{1.001, 1000, 2.71692393224},               // 0.00s
		{1.0001, 10000, 2.71814592682},             // 0.00s
		{1.00001, 100000, 2.71826823719},           // 0.00s
		{1.000001, 1000000, 2.7182804691},          // 0.00s
		{1.0000001, 10000000, 2.71828169413},       // 0.03s
		{1.00000001, 100000000, 2.71828179835},     // 0.14s
		{1.000000001, 1000000000, 2.71828205201},   // 1.25s
		{1.0000000001, 10000000000, 2.71828205323}, // 12.53s
	}

	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			t.Parallel()

			res := Iter(tt.x, tt.y)
			res = math.Round(res*roundUntil) / roundUntil
			require.Equal(t, tt.expected, res)
		})
	}
}
