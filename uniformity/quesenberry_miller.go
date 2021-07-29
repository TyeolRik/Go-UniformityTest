package uniformity

import (
	"log"
	"sort"
)

// Greenwood-Quesenberry-Miller test for uniformity
// GPL 3.0 License
// https://github.com/cran/uniftest/blob/master/R/quesenberry.unif.test.r
func QuesenberryMiller(sortedData *[]float64, nreplArg ...int) (testStatistics float64, P_value float64) {
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

	x := make([]float64, 1, len(*sortedData)+2)
	x[0] = 0
	x = append(x, *sortedData...)
	x = append(x, 1)
	tempSum1 := 0.0
	for idx := 0; idx < n+1; idx++ {
		tempSum1 += (x[idx+1] - x[idx]) * (x[idx+1] - x[idx])
	}
	tempSum2 := 0.0
	for idx := 0; idx < n; idx++ {
		tempSum2 += (x[idx+1] - x[idx]) * (x[idx+2] - x[idx+1])
	}
	q := tempSum1 + tempSum2
	testStatistics = q

	for i := 0; i < nrepl; i++ {
		tempZ := runif(n)
		sort.Float64s(tempZ)
		z := make([]float64, 1, len(tempZ)+2)
		z[0] = 0
		z = append(z, tempZ...)
		z = append(z, 1)
		tempSum1 = 0.0
		for idx := 0; idx < n+1; idx++ {
			tempSum1 += (z[idx+1] - z[idx]) * (z[idx+1] - z[idx])
		}
		tempSum2 = 0.0
		for idx := 0; idx < n; idx++ {
			tempSum2 += (z[idx+1] - z[idx]) * (z[idx+2] - z[idx+1])
		}
		Q := tempSum1 + tempSum2
		if Q > q {
			l = l + 1.0
		}
	}

	P_value = l / float64(nrepl)
	return
}
