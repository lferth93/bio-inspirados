package main

import (
	"fmt"
	"math/rand"
)

func solve(data [][]float64, iter int) []int {

	centroid := [][]float64{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	sol := initSolution(data, centroid)
	fmt.Println(centroid)
	fmt.Println(data)
	return sol
}

func initSolution(data, centroid [][]float64) []int {
	sol := make([]int, len(data))
	cont := []int{0, 0, 0}

	for i := range sol {
		sol[i] = rand.Intn(3)
		cont[sol[i]]++
	}

	for i := range sol {
		for j := range centroid[sol[i]] {
			centroid[sol[i]][j] += data[i][j]
		}
	}

	for i := range centroid {
		den := float64(max(1, cont[i]))
		for j := range centroid[i] {
			centroid[i][j] /= den
		}
	}

	return sol
}
