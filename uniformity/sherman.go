package uniformity

import (
	"log"
	"math"
	"sort"
)

// Sherman test for the hypothesis of uniformity
// GPL 3.0 License
// https://github.com/cran/uniftest/blob/master/R/sherman.unif.test.R
func Sherman(sortedData *[]float64, nreplArg ...int) (testStatistics float64, P_value float64) {
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
	n := len(*sortedData)
	n_float64 := float64(n)
	x := make([]float64, 1, n+2)
	x[0] = 0
	x = append(x, *sortedData...)
	x = append(x, 1)

	w := 0.0
	for idx := 0; idx < n+1; idx++ {
		w += math.Abs(x[idx+1] - x[idx] - 1.0/(n_float64+1.0))
	}
	w = w / 2.0
	testStatistics = w

	for i := 0; i < nrepl; i++ {
		temp := runif(n)
		sort.Float64s(temp)
		z := make([]float64, 1, n+2)
		z[0] = 0.0
		z = append(z, temp...)
		z = append(z, 1)

		W := 0.0
		for idx := 0; idx < n+1; idx++ {
			W += math.Abs(z[idx+1] - z[idx] - 1.0/(n_float64+1.0))
		}
		W = W / 2.0
		if W > w {
			l = l + 1.0
		}
	}
	P_value = l / float64(nrepl)
	return
}
