package main

import (
	"fmt"
	"math/rand"
)

func solve(data [][]float64, iter int) []int {

	centroid := [][]float64{
		[]float64{0, 0, 0, 0},
		[]float64{0, 0, 0, 0},
		[]float64{0, 0, 0, 0},
	}

	sol := initSolution(data, centroid)
	return sol
}

func initSolution(data, centroid [][]float64) []int {
	sol := make([]int, len(data))
	cont := []int{0, 0, 0}

	for i := range sol {
		sol[i] = rand.Intn(3)
		cont[sol[i]]++
	}

	fmt.Println(cont)

	return sol
}
