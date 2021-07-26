package uniformity

import (
	"log"
	"math"
	"sort"
)

// Frosini test for the hypothesis of uniformity
// GPL 3.0 License
// https://github.com/cran/uniftest/blob/master/R/frosini.unif.test.r
func Frosini(data *[]float64, nreplArg ...int) (testStatistics float64, P_value float64) {
	var nrepl int
	if len(nreplArg) == 0 {
		nrepl = 2000
	} else if len(nreplArg) == 1 {
		nrepl = nreplArg[0]
	} else {
		log.Println("Input nrepl is wrong. Length of nreplArg is", len(nreplArg), nreplArg)
		return
	}

	l := 0.0
	n := len(*data)
	n_float64 := float64(n)

	tempSum := 0.0
	for i := 0; i < n; i++ {
		tempSum += math.Abs((*data)[i] - (float64(i)+0.5)/n_float64)
	}
	testStatistics = 1.0 / math.Sqrt(n_float64) * tempSum

	// Monte Carlo simulation for P-value
	for i := 0; i < nrepl; i++ {
		z := runif(n)
		sort.Float64s(z)

		tempSum = 0.0
		for idx := 0; idx < n; idx++ {
			tempSum += math.Abs(z[idx] - (float64(idx)+0.5)/n_float64)
		}
		B := 1.0 / math.Sqrt(n_float64) * tempSum
		if B > testStatistics {
			l = l + 1.0
		}
	}

	P_value = l / float64(nrepl)
	return
}
