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

func UniformityTest(data *[]float64) {
	sort.Float64s(*data)

	t = table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Test Name", "Test Statistics", "P-value", "Pass or Fail"})

	// 1. Dudewicz-van der Meulen test
	testStatistics, P_value = DudewiczVanDerMeulen(data)
	doTest("Dudewicz-van der Meulen test", testStatistics, P_value, 0.05)
	t.AppendSeparator()

	// 2. Frosini test for the hypothesis of uniformity
	testStatistics, P_value = Frosini(data)
	doTest("Frosini test", testStatistics, P_value, 0.05)
	t.AppendSeparator()

	// 3. Hegazy-Green test for the hypothesis of uniformity
	testStatistics, P_value = HegazyGreen(data)
	doTest("Hegazy-Green test", testStatistics, P_value, 0.05)
	t.AppendSeparator()

	// 4. Kolmogorov-Smirnov test for the hypothesis of uniformity

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
