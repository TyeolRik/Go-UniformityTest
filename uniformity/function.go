package uniformity

import (
	"math/rand"
	"time"
)

// runif implements runif in R.
func runif(n int) (ret []float64) {
	// in R, default form of runif is Mersenne Twister
	// https://github.com/SurajGupta/r-source/blob/master/src/main/RNG.c#L132-L133
	// **This function could be flaw**, that math/rand of golang is not mt19937
	// (And also I check that there is difference between runif in R and math/rand in go with same seed.)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	ret = make([]float64, n)
	for i := range ret {
		ret[i] = r.Float64()
	}
	return
}
