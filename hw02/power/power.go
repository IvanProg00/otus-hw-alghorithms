package power

func PowerIter(x float64, y int) float64 {
	var res float64 = 1

	for i := 0; i < y; i++ {
		res *= x
	}

	return res
}
