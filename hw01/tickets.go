package tickets

func CountNumSum(decimals int) int {
	return counterNumSum(decimals, 0, 0, 0)
}

func counterNumSum(decimals, sumA, sumB, count int) int {
	if decimals == 0 {
		if sumA == sumB {
			count++
		}
		return count
	}

	for a := 0; a < 10; a++ {
		for b := 0; b < 10; b++ {
			count = counterNumSum(decimals-1, sumA+a, sumB+b, count)
		}
	}

	return count
}
