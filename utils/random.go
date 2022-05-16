package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

const MaxNumber = 1_000_000

func RandomSlice(size int) ([]int, error) {
	arr := make([]int, size)

	for i := range arr {
		n, err := rand.Int(rand.Reader, big.NewInt(MaxNumber))
		if err != nil {
			return []int{}, fmt.Errorf("error get random int: %w", err)
		}

		arr[i] = int(n.Int64())
	}

	return arr, nil
}
