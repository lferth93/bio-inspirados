package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func main() {
	in1, err := os.Open("tarea1/TS.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer in1.Close()
	in2, err := os.Open("tarea2/DE.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer in2.Close()
	r1 := csv.NewReader(in1)
	r2 := csv.NewReader(in2)
	values := make([]plotter.Values, 2)
	values[0] = make(plotter.Values, 30)
	values[1] = make(plotter.Values, 30)
	ts, err := r1.ReadAll()

	if err != nil {
		log.Fatal(err)
	}
	for i := range ts {
		values[0][i], err = strconv.ParseFloat(ts[i][0], 64)
		if err != nil {
			log.Fatal(err)
		}
	}
	de, err := r2.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	for i := range de {
		values[1][i], err = strconv.ParseFloat(de[i][0], 64)
		if err != nil {
			log.Fatal(err)
		}
	}

	p := plot.New()

	p.Title.Text = "Gráfica DE vs TS"
	p.Y.Label.Text = "Costos"
	p.X.Label.Text = "Heurística"

	err = plotutil.AddBoxPlots(p, vg.Points(25), "TS", values[0], "DE", values[1])
	if err != nil {
		log.Fatal(err)
	}

	if err := p.Save(2.5*vg.Inch, 4*vg.Inch, "tarea2/TSvsDE.png"); err != nil {
		log.Fatal(err)
	}
}
