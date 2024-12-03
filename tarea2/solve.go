package main

import (
	"math"
	"math/rand"
)

func solve(data [][]float64, p int, iter int, f, cr float64) (float64, []int) {
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
			son := merge(pop[j], mut, cr)
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

	return realCost(bestSol, data), bestSol
}

func merge(f [][]float64, mut [][]float64, cr float64) [][]float64 {
	res := make([][]float64, len(f))
	for i := range res {
		res[i] = make([]float64, len(f[0]))
		for j := range res[0] {
			r := rand.Float64()
			if r < cr {
				res[i][j] = mut[i][j]
			} else {
				res[i][j] = f[i][j]
			}
		}
	}
	i, j := rand.Intn(len(res)), rand.Intn(len(res[0]))
	res[i][j] = mut[i][j]
	return res
}

func mutate(pop [][][]float64, t int, f float64) [][]float64 {
	elected := make([]bool, len(pop))
	elected[t] = true
	v := [3]int{}
	for i := range v {
		tmp := rand.Intn(len(pop))
		for elected[tmp] {
			tmp = rand.Intn(len(pop))
		}
		v[i] = tmp
	}

	mut := make([][]float64, len(pop[0]))
	for i := range mut {
		mut[i] = make([]float64, len(pop[0][0]))
		for j := range mut[i] {
			mut[i][j] = pop[v[0]][i][j] + f*(pop[v[1]][i][j]-pop[v[2]][i][j])
		}
	}
	return mut
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

func realCost(sol []int, data [][]float64) float64 {
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
