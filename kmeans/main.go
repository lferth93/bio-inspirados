package main

import (
	"fmt"
	"math"

	"gonum.org/v1/plot/plotter"
)

func main() {
	data := readData("IRIS.csv")
	normalize(data)
	values := make([]plotter.Values, 4)
	names := []string{"50", "100", "200", "500"}
	bestCost := math.MaxFloat64
	bestSol := make([]int, 150)
	for i, iter := range []int{50, 100, 200, 500} {
		values[i] = make(plotter.Values, 50)
		for j := range values[i] {
			cost, sol := solve(data, 3, iter)
			values[i][j] = cost
			if cost < bestCost {
				copy(bestSol, sol)
				bestCost = cost
			}
		}
	}
	fmt.Println("Best cost: ", bestCost)
	fmt.Println(bestSol[:50])
	fmt.Println(bestSol[50:100])
	fmt.Println(bestSol[100:])
	makePlot(values, names)
}
