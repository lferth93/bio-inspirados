package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"

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
			sol, cost := solve(data, iter, 20, 7)
			values[i][j] = cost
			if cost < bestCost {
				copy(bestSol, sol)
				bestCost = cost
			}
		}
	}

	out, err := os.Create("TS.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	w := csv.NewWriter(out)
	for i := 0; i < 30; i++ {
		_, cost := solve(data, 200, 20, 7)
		w.Write([]string{fmt.Sprint(cost)})
	}
	w.Flush()
	fmt.Println("Best cost: ", bestCost)
	fmt.Println(bestSol[:50])
	fmt.Println(bestSol[50:100])
	fmt.Println(bestSol[100:])
	makePlot(values, names)

}
