package uniformity

import (
	"log"
	"math"
)

// Kolmogorov-Smirnov test
// GPL 3.0 License
// https://github.com/cran/uniftest/blob/master/R/kolmogorov.unif.test.r
func KolmogorovSmirnov(data *[]float64, anotherArg ...int) (testStatistics float64, P_value float64) {
	var nrepl int
	var k int
	if len(anotherArg) == 0 {
		nrepl = 2000
		k = 0
	} else if len(anotherArg) == 2 {
		nrepl = anotherArg[0]
		k = anotherArg[1]
	} else {
		log.Println("Input nrepl is wrong. Length of anotherArg is", len(anotherArg), anotherArg)
		return
	}

	l := 0.0
	n := len(*data)
	n_float64 := float64(n)
	switch k {
	case -1:
		// Finding Max
		d := (*data)[0] - 1.0/(n_float64+1.0)
		for idx := 1; idx < n; idx++ {
			temp := (*data)[idx] - (float64(idx)+1.0)/(n_float64+1.0)
			if d < temp {
				d = temp
			}
		}
		testStatistics = d
		for i := 0; i < nrepl; i++ {
			z := runif(n)
			D := z[0] - 1.0/(n_float64+1.0)
			for idx := 1; idx < n; idx++ {
				temp := z[idx] - (float64(idx)+1.0)/(n_float64+1.0)
				if D < temp {
					D = temp
				}
			}
			if D > d {
				l = l + 1.0
			}
		}
	case 0:
		// Finding Max
		d := math.Abs(1.0/(n_float64+1.0) - (*data)[0])
		for idx := 1; idx < n; idx++ {
			temp := math.Abs((float64(idx)+1.0)/(n_float64+1.0) - (*data)[idx])
			if d < temp {
				d = temp
			}
		}
		testStatistics = d
		for i := 0; i < nrepl; i++ {
			z := runif(n)
			D := math.Abs(1.0/(n_float64+1.0) - z[0])
			for idx := 1; idx < n; idx++ {
				temp := math.Abs((float64(idx)+1.0)/(n_float64+1.0) - z[idx])
				if D < temp {
					D = temp
				}
			}
			if D > d {
				l = l + 1.0
			}
		}
	case 1:
		// Finding Max
		d := 1.0/(n_float64+1.0) - (*data)[0]
		for idx := 1; idx < n; idx++ {
			temp := (float64(idx)+1.0)/(n_float64+1.0) - (*data)[idx]
			if d < temp {
				d = temp
			}
		}
		testStatistics = d
		for i := 0; i < nrepl; i++ {
			z := runif(n)
			D := 1.0/(n_float64+1.0) - z[0]
			for idx := 1; idx < n; idx++ {
				temp := (float64(idx)+1.0)/(n_float64+1.0) - z[idx]
				if D < temp {
					D = temp
				}
			}
			if D > d {
				l = l + 1.0
			}
		}
	default:
		log.Println("Input K is wrong. Input K should be one of {-1, 0, 1}. But we got K =", k)
	}

	P_value = l / float64(nrepl)
	return
}
