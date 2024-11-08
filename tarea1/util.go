package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

type fixedList[T comparable] struct {
	data        []T
	cap, size   int
	first, last int
}

func newFixedList[T comparable](cap int) *fixedList[T] {
	return &fixedList[T]{
		data: make([]T, cap),
		cap:  cap,
	}
}

func (list *fixedList[T]) append(elem T) {
	list.data[list.last] = elem

	if list.last < list.cap-1 {
		list.last++
	} else {
		list.last = 0
	}

	if list.size == list.cap {
		list.first = list.last
	}
	list.size = min(list.size+1, list.cap)
}

func (list *fixedList[T]) has(elem T) bool {
	for i, c := list.first, 0; c < list.size; c++ {
		if list.data[i] == elem {
			return true
		}
		if i < list.cap-1 {
			i++
		} else {
			i = 0
		}
	}
	return false
}

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
