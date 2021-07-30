package rng

import "math/rand"

func MathRand(seeds *[]int64) (ret []float64) {
	ret = make([]float64, len(*seeds))
	for i := range *seeds {
		rand.Seed((*seeds)[i])
		ret[i] = rand.Float64()
	}
	return
}
