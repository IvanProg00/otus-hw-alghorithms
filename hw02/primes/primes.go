package primes

func Primes(n int) int {
	res := 0

	for i := 2; i <= n; i++ {
		count := 0

		for k := 2; k < i; k++ {
			if i%k == 0 {
				count++
			}
		}

		if count <= 0 {
			res++
		}
	}

	return res
}
