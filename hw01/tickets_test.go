package tickets

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCountNumSum(t *testing.T) {
	t.Parallel()

	tests := []struct {
		decimals int
		expected int
	}{
		{decimals: 1, expected: 10},
		{decimals: 2, expected: 670},
		{decimals: 3, expected: 55252},
		{decimals: 4, expected: 4816030},
		{decimals: 5, expected: 432457640},
		// {decimals: 6, expected: 39581170420},
		// {decimals: 7, expected: 3671331273480},
		// {decimals: 8, expected: 343900019857310},
		// {decimals: 9, expected: 32458256583753952},
		// {decimals: 10, expected: 3081918923741896840},
	}

	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			t.Parallel()

			res := CountNumSum(tt.decimals)
			require.Equal(t, tt.expected, res)
		})
	}
}
