package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	data := readData("IRIS.csv")
	normalize(data)
	sol := solve(data, 50)

	fmt.Println(sol)
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
