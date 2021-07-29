package uniformity

import "log"

// Kuiper test for the hypothesis of uniformity
// GPL 3.0 License
// https://github.com/cran/uniftest/blob/master/R/kuiper.unif.test.r
func Kuiper(data *[]float64, nreplArg ...int) (testStatistics float64, P_value float64) {
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

	var max1, max2 float64
	max1 = (*data)[0]
	max2 = 1.0/n_float64 - (*data)[0]
	for idx := 1; idx < n; idx++ {
		temp1 := (*data)[idx] - (float64(idx) / n_float64)
		temp2 := float64(idx+1)/n_float64 - (*data)[idx]
		if max1 < temp1 {
			max1 = temp1
		}
		if max2 < temp2 {
			max2 = temp2
		}
	}
	testStatistics = max1 + max2

	for i := 0; i < nrepl; i++ {
		z := runif(n)
		max1 = z[0]
		max2 = 1.0/n_float64 - z[0]
		for idx := 1; idx < n; idx++ {
			temp1 := z[idx] - (float64(idx) / n_float64)
			temp2 := float64(idx+1)/n_float64 - z[idx]
			if max1 < temp1 {
				max1 = temp1
			}
			if max2 < temp2 {
				max2 = temp2
			}
		}
		V := max1 + max2
		if V > testStatistics {
			l = l + 1.0
		}
	}
	P_value = l / float64(nrepl)
	return
}
