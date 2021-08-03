package rng

import (
	mathRand "math/rand"
)

func MathRand(seeds *[]int64) (ret []float64) {
	ret = make([]float64, len(*seeds))
	for i := range *seeds {
		mathRand.Seed((*seeds)[i])
		ret[i] = mathRand.Float64()
	}
	return
}

// CSPRNG
func CryptoRand(howMany int) (ret []float64) {
	ret = make([]float64, howMany)
	for i := range ret {
		ret[i] = GetCryptoRandFloat(0, 1)
	}
	return ret
}
