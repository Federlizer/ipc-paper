package main

import "math/big"

func multiplyArray(arr []*big.Int) *big.Int {
	total := big.NewInt(1)

	for _, val := range arr {
		total.Mul(total, val)
	}

	return total
}
