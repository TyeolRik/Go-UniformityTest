package uniformity

import (
	"log"
	"math"
	"sort"

	"github.com/tyeolrik/Go-Polynomial/polynomial"
)

// Neyman-Barton test for the hypothesis of uniformity
// Need Sorted Data.
// GPL 3.0 License
// https://github.com/cran/uniftest/blob/master/R/neyman.unif.test.r
func NeymanBarton(data *[]float64, anotherArg ...int) (testStatistics float64, P_value float64) {
	var nrepl int
	var k int
	if len(anotherArg) == 0 {
		nrepl = 2000
		k = 5
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

	pos := polynomial.LegendrePolynomials(k, true)
	ne := 0.0
	for index := 0; index < k; index++ {
		tempSum := 0.0
		for _, eachX := range *data {
			tempSum += pos[index+1].Evaluate(2*eachX-1) * math.Sqrt2
		}
		tempSum = tempSum / math.Sqrt(n_float64)
		tempSum = tempSum * tempSum
		ne += tempSum
	}
	testStatistics = ne

	for i := 0; i < nrepl; i++ {
		z := runif(n)
		sort.Float64s(z)
		Ne := 0.0

		for index := 0; index < k; index++ {
			tempSum := 0.0
			for _, eachX := range z {
				tempSum += pos[index+1].Evaluate(2*eachX-1) * math.Sqrt2
			}
			tempSum = tempSum / math.Sqrt(n_float64)
			tempSum = tempSum * tempSum
			Ne += tempSum
		}

		if Ne > ne {
			l = l + 1.0
		}
	}

	P_value = l / float64(nrepl)

	return
}
