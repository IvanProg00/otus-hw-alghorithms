package fibonachi

func Recursive(n int) uint64 {
	if n <= 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	return Recursive(n-1) + Recursive(n-2)
}

func Iter(n int) uint64 {
	if n <= 0 {
		return 0
	}

	var x, y uint64 = 0, 1

	for i := 1; i < n; i++ {
		x, y = y, x+y
		// Is equal to:
		// t := y
		// y = x + y
		// x = t
	}

	return y
}
