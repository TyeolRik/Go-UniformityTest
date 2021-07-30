package uniformity

import (
	"log"
	"path/filepath"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func Histogram(data *[]float64, locationToSave string) (savedFileLocation string) {
	p := plot.New()
	p.Title.Text = "Histogram plot"

	var value plotter.Values = *data

	hist, err := plotter.NewHist(value, 30)
	if err != nil {
		panic(err)
	}
	p.Add(hist)

	if locationToSave[len(locationToSave)-1] != '/' {
		locationToSave = locationToSave + "/"
	}

	if err := p.Save(3*vg.Inch, 3*vg.Inch, locationToSave+"histogram.png"); err != nil {
		panic(err)
	}

	path, err := filepath.Abs(locationToSave)
	if err != nil {
		log.Fatal(err)
	}
	savedFileLocation = path + "/histogram.png"
	return
}
