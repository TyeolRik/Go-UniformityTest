package uniformity

import (
	"log"
	"math"
	"sort"
)

// Hegazy-Green test for the hypothesis of uniformity
// anotherArg should be 2 value. (nrepl, p)
// GPL 3.0 License
// https://github.com/cran/uniftest/blob/master/R/hegazy.unif.test.r
func HegazyGreen(data *[]float64, anotherArg ...float64) (testStatistics float64, P_value float64) {
	var nrepl int
	var p float64
	if len(anotherArg) == 0 {
		nrepl = 2000
		p = 1.0
	} else if len(anotherArg) == 2 {
		nrepl = int(anotherArg[0])
		p = anotherArg[1]
	} else {
		log.Println("Input nrepl is wrong. Length of anotherArg is", len(anotherArg), anotherArg)
		return
	}

	l := 0.0
	n := len(*data)
	n_float64 := float64(n)

	tempSum := 0.0
	for i := 0; i < n; i++ {
		tempSum += math.Pow(math.Abs((*data)[i]-float64(i+1)/(n_float64+1.0)), p)
	}
	testStatistics = tempSum / n_float64

	// Monte Carlo simulation for P-value
	for i := 0; i < nrepl; i++ {
		z := runif(n)
		sort.Float64s(z)
		tempSum := 0.0
		for i := 0; i < n; i++ {
			tempSum += math.Pow(math.Abs(z[i]-float64(i+1)/(n_float64+1.0)), p)
		}
		T := tempSum / n_float64
		if T > testStatistics {
			l = l + 1
		}
	}

	P_value = l / float64(nrepl)
	return
}
