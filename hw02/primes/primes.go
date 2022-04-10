package primes

func Primes1(n int) int {
	var res int = 0

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
