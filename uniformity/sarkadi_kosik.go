package uniformity

import (
	"log"
	"sort"
)

// Sarkadi-Kosik test for the hypothesis of uniformity
// GPL 3.0 License
// https://github.com/cran/uniftest/blob/master/R/sarkadi.unif.test.r
func SarkadiKosik(sortedData *[]float64, nreplArg ...int) (testStatistics float64, P_value float64) {
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

	sum_d := 0.0
	sum_d2 := 0.0
	for idx := range *sortedData {
		temp := ((*sortedData)[idx] - (float64(idx)+1.0)/(n_float64+1.0)) / ((float64(idx) + 1.0) * (n_float64 - (float64(idx) + 1.0) + 1))
		sum_d += temp
		sum_d2 += temp * temp
	}
	j := n_float64*n_float64*sum_d2 - n_float64*sum_d*sum_d
	testStatistics = j

	for i := 0; i < nrepl; i++ {
		z := runif(n)
		sort.Float64s(z)
		sum_D := 0.0
		sum_D2 := 0.0
		for idx := range z {
			temp := (z[idx] - (float64(idx)+1.0)/(n_float64+1.0)) / ((float64(idx) + 1.0) * (n_float64 - (float64(idx) + 1.0) + 1))
			sum_D += temp
			sum_D2 += temp * temp
		}
		J := n_float64*n_float64*sum_D2 - n_float64*sum_D*sum_D
		if J > j {
			l = l + 1.0
		}
	}
	P_value = l / float64(nrepl)
	return
}
