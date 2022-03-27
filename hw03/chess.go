package chess

import "fmt"

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

func Horse(pos uint64) (int, uint64) {
	const leftSideOne uint64 = 0x7f7f7f7f7f7f7f7f
	const leftSideTwo uint64 = 0x3f3f3f3f3f3f3f3f
	const rightSideOne uint64 = 0xfefefefefefefefe
	const rightSideTwo uint64 = 0xfcfcfcfcfcfcfcfc
	pos = 1 << pos

	var posLeftSideOne uint64 = pos & leftSideOne
	var posLeftSideTwo uint64 = pos & leftSideTwo
	var posRightSideOne uint64 = pos & rightSideOne
	var posRightSideTwo uint64 = pos & rightSideTwo
	fmt.Println(posLeftSideOne, posLeftSideTwo)

	mask := posRightSideOne<<15 | posLeftSideOne<<17 |
		posRightSideTwo<<6 | posLeftSideTwo<<10 |
		posRightSideTwo>>10 | posLeftSideTwo>>6 |
		posRightSideOne>>17 | posLeftSideOne>>15

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
