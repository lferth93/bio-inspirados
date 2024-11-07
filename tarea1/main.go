package main

import "gonum.org/v1/plot/plotter"

func main() {
	data := readData("IRIS.csv")
	normalize(data)
	values := make([]plotter.Values, 3)
	names := []string{"50", "200", "500"}
	for i, iter := range []int{50, 200, 500} {
		values[i] = make(plotter.Values, 30)
		for j := range values[i] {
			_, cost := solve(data, iter, 50, 15)
			values[i][j] = cost
		}
	}
	makePlot(values, names)
}
