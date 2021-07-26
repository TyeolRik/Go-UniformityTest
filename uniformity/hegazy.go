package uniformity

import "log"

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

	n := len(*data)
	n_float64 := float64(n)

}
