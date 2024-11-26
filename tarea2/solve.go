package main

import (
	"math"
	"math/rand"
)

const (
	EXP      = 0
	BINOMIAL = 1
)

func solve(data [][]float64, p int, iter int, f, cr float64, mType int) (float64, []int) {
	pop := populate(p)
	vecs := make([][]int, p)
	costs := make([]float64, p)

	for i := range pop {
		c, v := getCost(pop[i], data)
		costs[i] = c
		vecs[i] = v
	}

	for i := 0; i < iter; i++ {
		for j := range pop {
			mut := mutate(pop, j, f)
			son := merge(pop[j], mut, cr, mType)
			cost, vec := getCost(son, data)
			if cost < costs[j] {
				copy(pop[j], son)
				copy(vecs[j], vec)
				costs[j] = cost
			}
		}
	}

	bestCost := costs[0]
	bestSol := vecs[0]

	for i := range costs {
		if costs[i] < bestCost {
			bestCost = costs[i]
			bestSol = vecs[i]
		}
	}

	return bestCost, bestSol
}

func merge(f [][]float64, mut [][]float64, cr float64, mType int) [][]float64 {
	panic("unimplemented")
}

func mutate(pop [][][]float64, i int, f float64) [][]float64 {
	panic("unimplemented")
}

func populate(ps int) [][][]float64 {
	p := make([][][]float64, ps)
	for i := range p {
		p[i] = make([][]float64, 3)
		for j := range p[i] {
			p[i][j] = make([]float64, 4)
			for k := range p[i][j] {
				p[i][j][k] = rand.Float64()
			}
		}
	}
	return p
}

func getCost(sol [][]float64, data [][]float64) (float64, []int) {
	cost := 0.0
	d := float64(0)
	cl := 0
	vec := make([]int, len(data))
	for i := range vec {
		d = math.MaxFloat64
		cl = 0
		for j := range sol {
			tmp := manhattan(sol[j], data[i])
			if tmp < d {
				d = tmp
				cl = j
			}
		}
		cost += d
		vec[i] = cl
	}
	return cost, vec
}
