package chess

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestKing(t *testing.T) {
	t.Parallel()

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
		i, tt := i, tt
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			t.Parallel()
			require := require.New(t)
			numMoves, mask := King(tt.pos)

			require.Equal(tt.expMask, mask)
			require.Equal(tt.expNumMoves, numMoves)
		})
	}
}

func TestHorse(t *testing.T) {
	t.Parallel()

	tests := []struct {
		pos         uint64
		expNumMoves int
		expMask     uint64
	}{
		{0, 2, 132096},
		{1, 3, 329728},
		{2, 4, 659712},
		{36, 8, 11333767002587136},
		{47, 4, 4620693356194824192},
		{48, 3, 288234782788157440},
		{54, 4, 1152939783987658752},
		{55, 3, 2305878468463689728},
		{56, 2, 1128098930098176},
		{63, 2, 9077567998918656},
	}

	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			t.Parallel()
			require := require.New(t)
			numMoves, mask := Horse(tt.pos)

			require.Equal(tt.expMask, mask)
			require.Equal(tt.expNumMoves, numMoves)
		})
	}
}
