package uniformity

import (
	"log"
	"math"
	"sort"
)

// Dudewicz-van der Meulen test for the hypothesis of uniformity.
// GPL 3.0 License
// https://github.com/cran/uniftest/blob/master/R/dudewicz.unif.test.r
// x is a numeric vector of data values.
// nrepl is the number of replications in Monte Carlo simulation. Default 2000
func DudewiczVanDerMeulen(data *[]float64, nreplArg ...int) (testStatistics float64, P_value float64) {
	var nrepl int
	if len(nreplArg) == 0 {
		nrepl = 2000
	} else if len(nreplArg) == 1 {
		nrepl = nreplArg[0]
	} else {
		log.Println("Input nrepl is wrong. Length of nreplArg is", len(nreplArg), nreplArg)
		return
	}

	m := len(*data) / 2
	l := 0.0
	n := len(*data)
	a := (*data)[0]
	b := (*data)[n-1]

	front := make([]float64, m)
	end := make([]float64, m)
	for i := range front {
		front[i] = a
		end[i] = b
	}
	x := append(front, (*data)...)
	x = append(x, end...)

	tempSum := 0.0
	for i := 0; i < n; i++ {
		tempSum += math.Log2(float64(n) / 2.0 / float64(m) * (x[i+2*m] - x[i]))
	}
	h := (-1.0) / float64(n) * tempSum
	testStatistics = h

	// Monte Carlo simulation for P-value
	for i := 0; i < nrepl; i++ {
		z := runif(n)
		sort.Float64s(z)
		a = z[0]
		b = z[n-1]

		front = make([]float64, m)
		end = make([]float64, m)
		for idx := range front {
			front[idx] = a
			end[idx] = b
		}
		z = append(front, z...)
		z = append(z, end...)

		tempSum = 0.0
		for i := 0; i < n; i++ {
			tempSum += math.Log2(float64(n) / 2.0 / float64(m) * (z[i+2*m] - z[i]))
		}
		H := (-1.0) / float64(n) * tempSum
		if H > h {
			l = l + 1.0
		}

		z = nil
	}
	P_value = l / float64(nrepl)

	return
}
