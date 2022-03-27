package chess

func King(pos uint64) (int, uint64) {
	const leftSide uint64 = 0x7f7f7f7f7f7f7f7f
	const rightSide uint64 = 0xfefefefefefefefe
	pos = 1 << pos

	var posLeftSide uint64 = pos & leftSide
	var posRightSide uint64 = pos & rightSide

	mask := posRightSide<<7 | pos<<8 | posLeftSide<<9 |
		posRightSide>>1 | posLeftSide<<1 |
		posRightSide>>9 | pos>>8 | posLeftSide>>7

	return CountMoves(mask), mask
}

func CountMoves(mask uint64) int {
	moves := 0

	for mask > 0 {
		mask &= mask - 1
		moves++
	}

	return moves
}
