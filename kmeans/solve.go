package main

import (
	"math"
	"math/rand"
)

func solve(data [][]float64, k, iter int) (float64, []int) {
	centroids := make([][]float64, k)
	sol := make([]int, len(data))

	for i := range centroids {
		centroids[i] = make([]float64, len(data[0]))
		for j := range centroids[i] {
			centroids[i][j] = rand.Float64()
		}
	}
	cost := assign(sol, centroids, data)
	for i := 0; i < iter; i++ {
		update(centroids, data, sol)
		cost = assign(sol, centroids, data)
	}
	return cost, sol
}

func assign(sol []int, centroids, data [][]float64) float64 {
	cost := 0.0
	for i := range sol {
		cen := 0
		lcost := math.MaxFloat64
		for j := range centroids {
			tmp := manhattan(centroids[j], data[i])
			if tmp < lcost {
				lcost = tmp
				cen = j
			}
		}
		sol[i] = cen
		cost += lcost
	}
	return cost
}

func update(centroids, data [][]float64, sol []int) {
	for i := range centroids {
		for j := range centroids[i] {
			centroids[i][j] = 0.0
		}
	}
	count := make([]int, len(centroids))
	for i, c := range sol {
		count[c]++
		for j := range centroids[c] {
			centroids[c][j] += data[i][j]
		}
	}
	for i := range centroids {
		for j := range centroids[i] {
			centroids[i][j] /= float64(count[i])
		}
	}
}
