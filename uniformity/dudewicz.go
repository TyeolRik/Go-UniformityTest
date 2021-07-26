package uniformity

import "log"

// Dudewicz-van der Meulen test for the hypothesis of uniformity.
//
// GPL 3.0 License
// https://github.com/cran/uniftest/blob/master/R/dudewicz.unif.test.r
// x is a numeric vector of data values.
// nrepl is the number of replications in Monte Carlo simulation. Default 2000
func DudewiczVanDerMeulen(data *[]float64, nreplArg ...float64) (testStatistics float64, P_value float64) {
	var nrepl float64
	if len(nreplArg) == 0 {
		nrepl = 2000.0
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

	for i := 0; i < n; i++ {

	}

	return
}
