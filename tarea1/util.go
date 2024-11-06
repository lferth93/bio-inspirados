package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type fixedList[T comparable] struct {
	data        []T
	cap, size   int
	first, last int
}

func newFixedList[T comparable](cap int) *fixedList[T] {
	return &fixedList[T]{
		data: make([]T, cap),
		cap:  cap,
	}
}

func (list *fixedList[T]) append(elem T) {
	list.data[list.last] = elem

	if list.last < list.cap-1 {
		list.last++
	} else {
		list.last = 0
	}

	if list.size == list.cap {
		list.first = list.last
	}
	list.size = min(list.size+1, list.cap)
}

func (list *fixedList[T]) has(elem T) bool {
	for i, c := list.first, 0; c < list.size; c++ {
		if list.data[i] == elem {
			return true
		}
		if i < list.cap-1 {
			i++
		} else {
			i = 0
		}
	}
	return false
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