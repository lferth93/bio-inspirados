package main

import (
	"fmt"
	"math"
	"math/rand"
)

func solve(data [][]float64, iter int) []int {

	centroid := [][]float64{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	cont := []int{0, 0, 0}
	tabu := newFixedList[int](7)

	sol, cost := initSolution(data, centroid, cont)
	bestSol, bestCost := make([]int, 150), cost
	fmt.Println(bestCost)
	copy(bestSol, sol)
	for i := 0; i < iter; i++ {
		n := neighborhood(10)
		bestLocal := math.MaxFloat64
		bestLocalN := 0
		for j := range n {
			src := sol[n[j][0]]
			dst := (src + n[j][1]) % 3
			nCost := goal(data, centroid, sol, cont[src], cont[dst], n[j][0], dst)
			if nCost < bestLocal {
				bestLocal = nCost
				bestLocalN = j
			}
		}
		if bestLocal < bestCost {
			bestCost = bestLocal
			fmt.Println(bestCost)
		} else {
			if tabu.has(n[bestLocalN][0]) {
				continue
			}
		}
		tabu.append(n[bestLocalN][0])
		src := sol[n[bestLocalN][0]]
		dst := (src + n[bestLocalN][1]) % 3
		updateSol(data, centroid, sol, cont, n[bestLocalN][0], dst)
		if bestLocal == bestCost {
			copy(bestSol, sol)
		}
	}
	return bestSol
}

func initSolution(data, centroid [][]float64, cont []int) ([]int, float64) {
	sol := make([]int, len(data))

	for i := range sol {
		sol[i] = rand.Intn(3)
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

	cost := 0.0

	for i := range sol {
		for j := range centroid[sol[i]] {
			cost += math.Abs(data[i][j] - centroid[sol[i]][j])
		}
	}

	return sol, cost
}

func neighborhood(len int) [][2]int {
	visited := make([]bool, 150)
	n := make([][2]int, len)
	for i := range n {
		d := 0
		for d = rand.Intn(150); visited[d]; d = rand.Intn(150) {
		}
		visited[d] = true
		c := rand.Intn(2) + 1
		n[i][0] = d
		n[i][1] = c
	}
	return n
}

func goal(data, centroid [][]float64, sol []int, srcC, dstC, pos, dst int) float64 {
	cost := 0.0
	point := data[pos]
	srcCent := make([]float64, len(centroid[0]))
	copy(srcCent, centroid[sol[pos]])
	dstCent := make([]float64, len(centroid[0]))
	copy(dstCent, centroid[dst])
	for i := range srcCent {
		srcCent[i] = (float64(srcC)*srcCent[i] - point[i]) / float64(max(srcC-1, 1))
		dstCent[i] = (float64(dstC)*dstCent[i] + point[i]) / float64(dstC+1)
	}
	prevClust := sol[pos]
	sol[pos] = dst
	for i := range sol {
		cent := centroid[sol[i]]
		if sol[i] == dst {
			cent = dstCent
		}
		if sol[i] == prevClust {
			cent = srcCent
		}

		for j := range cent {
			cost += math.Abs(data[i][j] - cent[j])
		}
	}
	sol[pos] = prevClust
	return cost
}

func updateSol(data [][]float64, centroid [][]float64, sol []int, cont []int, pos, dst int) {
	point := data[pos]
	src := sol[pos]
	for i := range centroid[0] {
		centroid[src][i] = (float64(cont[src])*centroid[src][i] - point[i]) / float64(max(cont[src]-1, 1))
		centroid[dst][i] = (float64(cont[dst])*centroid[dst][i] + point[i]) / float64(cont[dst]+1)
	}
	cont[src]--
	cont[dst]--
	sol[pos] = dst
}
