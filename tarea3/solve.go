package main

import (
	"math"
	"math/rand"
)

func solve(data [][]float64, iter, population int, p, m, ps float64) ([]int, float64) {
	pop := initPop(population, len(data))
	costs := make([]float64, population)
	updateCosts(data, pop, costs)
	bestSol := make([]int, len(data))
	bi := getBest(costs)
	bestCost := costs[bi]
	copy(bestSol, pop[bi])

	for i := 0; i < iter; i++ {
		p1 := selectP(costs, ps)
		p2 := selectP(costs, ps)
		sons := make([][]int, len(p1)*2)
		for j := range p1 {
			sons[2*j] = make([]int, len(data))
			sons[2*j+1] = make([]int, len(data))
			copy(sons[2*j], pop[p1[j]])
			copy(sons[2*j+1], pop[p2[j]])
			for k := range len(data) {
				if rand.Float64() < p {
					tmp := sons[2*j][k]
					sons[2*j][k] = sons[2*j+1][k]
					sons[2*j+1][k] = tmp
				}
				if rand.Float64() < m {
					sons[2*j][k] += rand.Intn(2) + 1
					sons[2*j][k] %= 3
				}
				if rand.Float64() < m {
					sons[2*j+1][k] += rand.Intn(2) + 1
					sons[2*j+1][k] %= 3
				}
			}
		}
		for j := range pop {
			copy(pop[j], sons[j])
		}

		updateCosts(data, pop, costs)
		bi = getBest(costs)
		bestCost = costs[bi]
		copy(bestSol, pop[bi])
	}

	return bestSol, bestCost
}

func selectP(costs []float64, ps float64) []int {
	s := rand.Perm(len(costs))
	p := make([]int, len(costs)/2)
	for i := range p {
		bi := 2 * i
		if (costs[s[bi]] < costs[s[bi+1]]) == (rand.Float64() < ps) {
			p[i] = s[bi]
		} else {
			p[i] = s[bi+1]
		}
	}
	return p
}

func initPop(p, l int) [][]int {
	pop := make([][]int, p)
	for i := range pop {
		pop[i] = make([]int, l)
		for j := range pop[i] {
			pop[i][j] = rand.Intn(3)
		}
	}
	return pop
}

func updateCosts(data [][]float64, pop [][]int, costs []float64) {
	for i := range pop {
		costs[i] = getCost(pop[i], data)
	}
}

func getCost(sol []int, data [][]float64) float64 {
	centroid := make([][]float64, 3)

	for i := range centroid {
		centroid[i] = make([]float64, len(data[0]))
	}

	count := make([]int, len(centroid))

	for i, c := range sol {
		count[c]++
		for j := range centroid[c] {
			centroid[c][j] += data[i][j]
		}
	}
	for i := range centroid {
		for j := range centroid[i] {
			centroid[i][j] /= float64(count[i])
		}
	}
	cost := 0.0
	for i, c := range sol {
		cost += manhattan(data[i], centroid[c])
	}
	return cost
}

func getBest(costs []float64) int {
	b := 0
	cost := math.MaxFloat64
	for i := range costs {
		if costs[i] < cost {
			cost = costs[i]
			b = i
		}
	}
	return b
}
