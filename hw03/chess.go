package chess

func King(pos uint64) (int, uint64) {
	const (
		leftSide  uint64 = 0x7f7f7f7f7f7f7f7f
		rightSide uint64 = 0xfefefefefefefefe
	)

	pos = 1 << pos

	posLeftSide := pos & leftSide
	posRightSide := pos & rightSide

	mask := posRightSide<<7 | pos<<8 | posLeftSide<<9 |
		posRightSide>>1 | posLeftSide<<1 |
		posRightSide>>9 | pos>>8 | posLeftSide>>7

	return CountMoves1(mask), mask
}

func Horse(pos uint64) (int, uint64) {
	const (
		leftSideOne  uint64 = 0x7f7f7f7f7f7f7f7f
		leftSideTwo  uint64 = 0x3f3f3f3f3f3f3f3f
		rightSideOne uint64 = 0xfefefefefefefefe
		rightSideTwo uint64 = 0xfcfcfcfcfcfcfcfc
	)

	pos = 1 << pos
	posLeftSideOne := pos & leftSideOne
	posLeftSideTwo := pos & leftSideTwo
	posRightSideOne := pos & rightSideOne
	posRightSideTwo := pos & rightSideTwo

	mask := posRightSideOne<<15 | posLeftSideOne<<17 |
		posRightSideTwo<<6 | posLeftSideTwo<<10 |
		posRightSideTwo>>10 | posLeftSideTwo>>6 |
		posRightSideOne>>17 | posLeftSideOne>>15

	return CountMoves2(mask), mask
}

func CountMoves1(mask uint64) int {
	moves := 0

	for mask > 0 {
		mask &= mask - 1
		moves++
	}

	return moves
}

func CountMoves2(mask uint64) int {
	moves := 0

	for mask > 0 {
		if mask&1 == 1 {
			moves++
		}

		mask >>= 1
	}

	return moves
}
