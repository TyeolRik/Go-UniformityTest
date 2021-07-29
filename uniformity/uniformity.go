package uniformity

import (
	"fmt"
	"os"
	"sort"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

var testStatistics, P_value float64
var testNumber int = 1
var totalPassFail int = 0
var PassFail string = "PASS"

var t table.Writer

func UniformityTest(unSortedData *[]float64) {
	sortedData := make([]float64, len(*unSortedData))
	copy(sortedData, *unSortedData)
	sort.Float64s(sortedData)

	t = table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Test Name", "Test Statistics", "P-value", "Pass or Fail"})

	// 1. Dudewicz-van der Meulen test
	testStatistics, P_value = DudewiczVanDerMeulen(&sortedData)
	doTest("Dudewicz-van der Meulen test", testStatistics, P_value, 0.05)
	t.AppendSeparator()

	// 2. Frosini test for the hypothesis of uniformity
	testStatistics, P_value = Frosini(&sortedData)
	doTest("Frosini test", testStatistics, P_value, 0.05)
	t.AppendSeparator()

	// 3. Hegazy-Green test for the hypothesis of uniformity
	testStatistics, P_value = HegazyGreen(&sortedData)
	doTest("Hegazy-Green test", testStatistics, P_value, 0.05)
	t.AppendSeparator()

	// 4. Kolmogorov-Smirnov test for the hypothesis of uniformity
	//testStatistics, P_value = KolmogorovSmirnov(unSortedData, 2000, -1)
	//doTest("Kolmogorov-Smirnov test (k = -1)", testStatistics, P_value, 0.05)
	testStatistics, P_value = KolmogorovSmirnov(unSortedData, 2000, 0)
	doTest("Kolmogorov-Smirnov test (k =  0)", testStatistics, P_value, 0.05)
	//testStatistics, P_value = KolmogorovSmirnov(unSortedData, 2000, 1)
	//doTest("Kolmogorov-Smirnov test (k = +1)", testStatistics, P_value, 0.05)
	t.AppendSeparator()

	// 5. Kuiper test for the hypothesis of uniformity
	testStatistics, P_value = Kuiper(unSortedData)
	doTest("Kuiper test", testStatistics, P_value, 0.05)
	t.AppendSeparator()

	// 6. Neyman-Barton test for the hypothesis of uniformity
	testStatistics, P_value = NeymanBarton(&sortedData)
	doTest("Neyman-Barton test", testStatistics, P_value, 0.05)
	t.AppendSeparator()

	// 7. Greenwood-Quesenberry-Miller test for uniformity
	testStatistics, P_value = QuesenberryMiller(&sortedData)
	doTest("Quesenberry–Miller test", testStatistics, P_value, 0.05)
	t.AppendSeparator()

	// 8. Sarkadi-Kosik test for the hypothesis of uniformity
	testStatistics, P_value = SarkadiKosik(&sortedData)
	doTest("Sarkadi-Kosik test", testStatistics, P_value, 0.05)
	t.AppendSeparator()

	var resultColor text.Colors = text.Colors{text.FgHiRed}
	resultPassFail := "FAIL"
	if totalPassFail > (testNumber / 2) {
		resultPassFail = "PASS"
		resultColor = text.Colors{text.FgHiGreen}
	}

	t.SetColumnConfigs([]table.ColumnConfig{
		{Number: 3, Align: text.AlignRight, AlignHeader: text.AlignCenter, AlignFooter: text.AlignCenter},
		{Number: 4, Align: text.AlignRight, AlignHeader: text.AlignCenter, AlignFooter: text.AlignCenter, ColorsFooter: resultColor},
		{Number: 5, Align: text.AlignCenter, AlignHeader: text.AlignCenter, AlignFooter: text.AlignCenter, ColorsFooter: resultColor},
	})

	t.AppendFooter(table.Row{"", "", "", "Result", resultPassFail})
	t.Render()

	t = nil
}

func doTest(testName string, testStatistics float64, P_value float64, significanceLevel float64) {
	if P_value > 0.05 {
		totalPassFail++
		PassFail = "PASS"
	} else {
		PassFail = "Fail"
	}
	t.AppendRow([]interface{}{testNumber, testName, fmt.Sprintf("%.08f", testStatistics), fmt.Sprintf("%.07f", P_value), PassFail})
	testNumber++
}
