package chess

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestKing(t *testing.T) {
	tests := []struct {
		pos         uint64
		expNumMoves int
		expMask     uint64
	}{
		{0, 3, 770},
		{1, 5, 1797},
		{7, 3, 49216},
		{8, 5, 197123},
		{10, 8, 920078},
		{15, 5, 12599488},
		{54, 8, 16186183351374184448},
		{55, 5, 13853283560024178688},
		{56, 3, 144959613005987840},
		{63, 3, 4665729213955833856},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			require := require.New(t)
			numMoves, mask := King(tt.pos)

			require.Equal(tt.expMask, mask)
			require.Equal(tt.expNumMoves, numMoves)
		})
	}
}
