package main

import (
	"math"
	"math/rand"
)

func solve(data [][]float64, iter int, neighbors, t int) ([]int, float64) {
	//fmt.Println("iterations: ",iter)

	tabu := newFixedList[int](t)

	sol, cost := initSolution(data)
	bestSol, bestCost := make([]int, 150), cost
	copy(bestSol, sol)
	for i := 0; i < iter; i++ {
		n := neighborhood(sol, neighbors)
		bestLocal := math.MaxFloat64
		bestLocalN := 0
		for j := range n {
			pos := n[j][0]
			dst := n[j][1]
			nCost := goal(data, sol, pos, dst)
			if nCost < bestLocal {
				bestLocal = nCost
				bestLocalN = j
			}
		}

		pos := n[bestLocalN][0]
		dst := n[bestLocalN][1]
		if tabu.has(pos) && bestLocal > bestCost {
			continue
		}
		tabu.append(pos)
		sol[pos] = dst
		if bestLocal < bestCost {
			bestCost = bestLocal
			copy(bestSol, sol)
		}
	}
	return bestSol, bestCost
}

func initSolution(data [][]float64) ([]int, float64) {
	sol := make([]int, len(data))

	for i := range sol {
		sol[i] = rand.Intn(3)
	}

	cost := getCost(sol, data)

	return sol, cost
}

func neighborhood(sol []int, len int) [][2]int {
	visited := make([]bool, 150)
	n := make([][2]int, len)
	for i := range n {
		d := 0
		for d = rand.Intn(150); visited[d]; d = rand.Intn(150) {
		}
		visited[d] = true
		c := (sol[d] + rand.Intn(2) + 1) % 3
		n[i][0] = d
		n[i][1] = c
	}
	return n
}

func goal(data [][]float64, sol []int, pos, dst int) float64 {
	prevClust := sol[pos]
	sol[pos] = dst
	cost := getCost(sol, data)
	sol[pos] = prevClust
	return cost
}

func getCost(sol []int, data [][]float64) float64 {
	cost := 0.0
	centroid := [][]float64{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}
	cont := []int{0, 0, 0}
	for i := range sol {
		cont[sol[i]]++
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

	for i := range sol {
		for j := range centroid[sol[i]] {
			cost += math.Abs(data[i][j] - centroid[sol[i]][j])
		}
	}
	return cost
}
