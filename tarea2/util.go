package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func normalize(data [][]float64) {
	maxs := make([]float64, len(data[0]))
	mins := make([]float64, len(data[0]))

	copy(mins, data[0])

	for i := range data {
		for j := range data[i] {
			if data[i][j] > maxs[j] {
				maxs[j] = data[i][j]
			}
			if data[i][j] < mins[j] {
				mins[j] = data[i][j]
			}
		}
	}
	for i := range data {
		for j := range data[i] {
			data[i][j] = (data[i][j] - mins[j]) / (maxs[j] - mins[j])
		}
	}
}

func manhattan(a, b []float64) float64 {
	d := float64(0)
	for i := range a {
		d += math.Abs(a[i] - b[i])
	}
	return d
}

func readData(file string) [][]float64 {
	data := make([][]float64, 0, 150)
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewScanner(f)
	for reader.Scan() {
		l := strings.Split(reader.Text(), ",")
		ld := make([]float64, 0, 4)
		for i := 0; i < 4 && i < len(l); i++ {
			df := 0.0
			fmt.Sscanf(l[i], "%f", &df)
			ld = append(ld, df)
		}
		data = append(data, ld)
	}

	return data
}

func makePlot(values []plotter.Values, names []string) error {

	p := plot.New()

	p.Title.Text = "Gráfica de cajas"
	p.Y.Label.Text = "Costos"
	p.X.Label.Text = "Iteraciones"

	plotData := make([]interface{}, 0, 2*len(values))

	for i := range values {
		plotData = append(plotData, names[i], values[i])
	}

	err := plotutil.AddBoxPlots(p, vg.Points(25), plotData...)
	if err != nil {
		return err
	}

	if err := p.Save(4*vg.Inch, 4*vg.Inch, "boxplot.png"); err != nil {
		return err
	}
	return nil
}
